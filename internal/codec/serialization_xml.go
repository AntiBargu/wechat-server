package codec

import "encoding/xml"

type XMLSerialization struct {
}

func (*XMLSerialization) Unmarshal(data []byte, body interface{}) error {
	return xml.Unmarshal(data, &body)
}

func (*XMLSerialization) Marshal(body interface{}) ([]byte, error) {
	return xml.Marshal(body)
}
