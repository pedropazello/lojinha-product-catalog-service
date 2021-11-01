package serializers

import (
	"encoding/xml"
)

type XMLSerializer struct{}

func (s *XMLSerializer) Serialize(obj interface{}) ([]byte, error) {
	xmlObj, err := xml.Marshal(obj)
	return xmlObj, err
}
