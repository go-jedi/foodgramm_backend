package serializer

// JSONSerializer defines the interface for the json or jsoniter.
//
//go:generate mockery --name=JSONSerializer --output=mocks --case=underscore
type JSONSerializer interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

// if need tests with mocks jsoniter.Marshal and Unmarshal for json, jsoniter
//
/*
type jsoniterSerializer struct{}

func (j jsoniterSerializer) Marshal(v interface{}) ([]byte, error) {
	return jsoniter.Marshal(v)
}

func (j jsoniterSerializer) Unmarshal(data []byte, v interface{}) error {
	return jsoniter.Unmarshal(data, v)
}

var defaultJSONMarshaler JSONSerializer = jsoniterSerializer{}
*/
