package webfinger

import (
	"encoding/json"
)

func jsonMarshalNameStringsValue(name string, values ...string) ([]byte, error) {
	var buffer [256]byte
	var bytes []byte = buffer[0:0]

	{
		encoded, err := json.Marshal(name)
		if nil != err {
			return nil, err
		}

		bytes = append(bytes, encoded...)
	}

	bytes = append(bytes, ":["...)

	for index, value := range values {
		if 1 <= index {
			bytes = append(bytes, ',')
		}

		encoded, err := json.Marshal(value)
		if nil != err {
			return nil, err
		}

		bytes = append(bytes, encoded...)
	}

	bytes = append(bytes, ']')

	return bytes, nil
}
