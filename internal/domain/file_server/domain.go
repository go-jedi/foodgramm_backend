package fileserver

// SupportedImageTypes supported image MIME types.
var SupportedImageTypes = map[string]struct{}{
	"image/jpeg": {},
	"image/png":  {},
}

type UploadAndConvertToWebpResponse struct {
	NameFile       string `json:"name_file"`
	ServerPathFile string `json:"server_path_file"`
	ClientPathFile string `json:"client_path_file"`
	Extension      string `json:"extension"`
	Quality        int    `json:"quality"`
	OldNameFile    string `json:"old_name_file"`
	OldExtension   string `json:"old_extension"`
}
