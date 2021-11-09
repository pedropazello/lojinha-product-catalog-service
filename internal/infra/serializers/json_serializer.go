package serializers

import "encoding/json"

type JSONSerializer struct{}

func (s *JSONSerializer) Serialize(obj interface{}) ([]byte, error) {
	jsonObj, err := json.Marshal(obj)
	return jsonObj, err
}
