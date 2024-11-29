package webfinger

import (
	"encoding/json"
)

func jsonMarshalNameStringValue(name string, value string) ([]byte, error) {
	var buffer [256]byte
	var bytes []byte = buffer[0:0]

	{
		encoded, err := json.Marshal(name)
		if nil != err {
			return nil, err
		}

		bytes = append(bytes, encoded...)
	}

	bytes = append(bytes, ':')

	{
		encoded, err := json.Marshal(value)
		if nil != err {
			return nil, err
		}

		bytes = append(bytes, encoded...)
	}

	return bytes, nil
}
