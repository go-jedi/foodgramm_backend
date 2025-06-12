package fileserver

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-jedi/foodgramm_backend/config"
	fileserver "github.com/go-jedi/foodgramm_backend/internal/domain/file_server"
	"github.com/go-jedi/foodgramm_backend/pkg/apperrors"
	"github.com/google/uuid"
	"github.com/h2non/bimg"
)

// constants for file server configuration.
const (
	defaultMaxSize  = 10 << 20          // default maximum file size (10MB)
	defaultQuality  = 30                // default image quality for conversion
	defaultDirPerm  = os.FileMode(0755) // default directory permissions
	defaultFilePerm = os.FileMode(0644) // default file permissions
	webpExt         = ".webp"           // extension for WebP images
)

// error definitions.
var (
	ErrFileTooLarge     = errors.New("file size exceeds maximum allowed limit")
	ErrInvalidPath      = errors.New("invalid file path")
	ErrDirectoryMissing = errors.New("directory does not exist")
)

// IFileServer defines the interface for the file server.
//
//go:generate mockery --name=IFileServer --output=mocks --case=underscore
type IFileServer interface {
	UploadAndConvertToWebP(ctx context.Context, fileHeader *multipart.FileHeader) (string, error)
}

// FileServer handles file uploads and processing.
type FileServer struct {
	url          string      // Url for connect to image from client
	dir          string      // Base directory for file storage
	maxFileSize  int64       // Maximum allowed file size
	imageQuality int         // Quality for image conversion
	dirPerm      os.FileMode // Permission mode for directories
	filePerm     os.FileMode // Permission mode for files
}

// New creates a new FileServer instance with the given configuration.
func New(cfg config.FileServerConfig) *FileServer {
	fs := &FileServer{
		url:          cfg.URL,
		dir:          cfg.Dir,
		maxFileSize:  cfg.MaxFileSize,
		imageQuality: cfg.ImageQuality,
		dirPerm:      os.FileMode(cfg.DirPerm),
		filePerm:     os.FileMode(cfg.FilePerm),
	}

	fs.init()

	return fs
}

// init sets default values for any unconfigured FileServer properties.
func (fs *FileServer) init() {
	if fs.maxFileSize == 0 {
		fs.maxFileSize = defaultMaxSize
	}

	if fs.imageQuality == 0 {
		fs.imageQuality = defaultQuality
	}

	if fs.dirPerm == 0 {
		fs.dirPerm = defaultDirPerm
	}

	if fs.filePerm == 0 {
		fs.filePerm = defaultFilePerm
	}
}

// getFileExt get file extension.
func (fs *FileServer) getFileExt(filename string) string {
	return strings.TrimSuffix(filepath.Ext(filename), filepath.Base(filename))
}

// sanitizePath checks if the given path is safe to use.
// returns cleaned path or error if path is invalid.
func (fs *FileServer) sanitizePath(path string) (string, error) {
	if path == "" {
		return "", ErrInvalidPath
	}

	const (
		parentDir       = ".."
		parentDirPrefix = ".." + string(filepath.Separator) // "../" or "..\"
		parentDirInPath = string(filepath.Separator) + ".." // "/.." or "\.."
	)

	// check for path traversal attempts.
	switch {
	case strings.HasPrefix(path, parentDirPrefix),
		strings.Contains(path, parentDirInPath),
		path == parentDir:
		return "", ErrInvalidPath
	}

	return filepath.Clean(path), nil
}

// validateFile checks if the file meets size and format requirements.
func (fs *FileServer) validateFile(fileHeader *multipart.FileHeader) error {
	if fileHeader.Size > fs.maxFileSize {
		return ErrFileTooLarge
	}

	contentType := fileHeader.Header.Get("Content-Type")
	if _, ok := fileserver.SupportedImageTypes[contentType]; !ok {
		log.Printf("unsupported file type: %s", contentType)
		return apperrors.ErrUnsupportedFormat
	}

	return nil
}

// ensureUploadDirectory checks if the target directory exists.
func (fs *FileServer) ensureUploadDirectory(uploadPath string) error {
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		return fmt.Errorf("%w: %s", ErrDirectoryMissing, uploadPath)
	}
	return nil
}

// readFileData reads file content with context support and size limitation.
func (fs *FileServer) readFileData(ctx context.Context, file multipart.File) ([]byte, error) {
	done := make(chan struct{})
	var data []byte
	var readErr error

	go func() {
		data, readErr = io.ReadAll(io.LimitReader(file, fs.maxFileSize))
		close(done)
	}()

	select {
	case <-done:
		return data, readErr
	case <-ctx.Done():
		return nil, ctx.Err() // return if context is canceled.
	}
}

// convertToWebP converts image data to WebP format with context support.
func (fs *FileServer) convertToWebP(ctx context.Context, rawFile []byte) ([]byte, error) {
	done := make(chan struct{})
	var webp []byte
	var convertErr error

	go func() {
		options := bimg.Options{
			Quality: fs.imageQuality,
			Type:    bimg.WEBP,
		}
		webp, convertErr = bimg.Resize(rawFile, options)
		close(done)
	}()

	select {
	case <-done:
		return webp, convertErr
	case <-ctx.Done():
		return nil, ctx.Err() // return if context is canceled.
	}
}

// UploadAndConvertToWebP handles the file upload process including validation, conversion and storage.
func (fs *FileServer) UploadAndConvertToWebP(ctx context.Context, fileHeader *multipart.FileHeader) (fileserver.UploadAndConvertToWebpResponse, error) {
	// check if context is already canceled.
	if err := ctx.Err(); err != nil {
		return fileserver.UploadAndConvertToWebpResponse{}, err
	}

	// measure upload duration.
	start := time.Now()
	defer func() {
		log.Printf("upload took %v", time.Since(start))
	}()

	// validate the uploaded file.
	if err := fs.validateFile(fileHeader); err != nil {
		return fileserver.UploadAndConvertToWebpResponse{}, err
	}

	// sanitize the target directory path.
	sanitizedDir, err := fs.sanitizePath(fs.dir)
	if err != nil {
		return fileserver.UploadAndConvertToWebpResponse{}, fmt.Errorf("invalid directory: %w", err)
	}

	// open the uploaded file.
	file, err := fileHeader.Open()
	if err != nil {
		return fileserver.UploadAndConvertToWebpResponse{}, fmt.Errorf("failed to open file: %w", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("error closing file: %v", err)
		}
	}()

	// prepare full upload path and verify directory exists.
	if err := fs.ensureUploadDirectory(sanitizedDir); err != nil {
		return fileserver.UploadAndConvertToWebpResponse{}, err
	}

	// read file data with context support.
	rawFile, err := fs.readFileData(ctx, file)
	if err != nil {
		return fileserver.UploadAndConvertToWebpResponse{}, fmt.Errorf("failed to read file: %w", err)
	}

	// convert image to WebP format.
	webp, err := fs.convertToWebP(ctx, rawFile)
	if err != nil {
		return fileserver.UploadAndConvertToWebpResponse{}, fmt.Errorf("failed to convert image: %w", err)
	}

	// generate unique filename and save the converted image.
	newName := uuid.New().String()
	newFilePath := filepath.Join(sanitizedDir, newName+webpExt)

	if err := os.WriteFile(newFilePath, webp, fs.filePerm); err != nil {
		return fileserver.UploadAndConvertToWebpResponse{}, fmt.Errorf("failed to save converted image: %w", err)
	}

	return fileserver.UploadAndConvertToWebpResponse{
		NameFile:       newName + webpExt,
		ServerPathFile: filepath.Join(sanitizedDir, newName+webpExt),
		ClientPathFile: filepath.Join(fs.url, newName+webpExt),
		Extension:      webpExt,
		Quality:        fs.imageQuality,
		OldNameFile:    fileHeader.Filename,
		OldExtension:   fs.getFileExt(fileHeader.Filename),
	}, nil
}
