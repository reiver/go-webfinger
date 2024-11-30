package webfinger

import (
	"encoding/json"

	"github.com/reiver/go-opt"
)

// Response represents a JSON Resource Descriptor (JRD) response.
type Response struct {
	Subject opt.Optional[string]
	Aliases []string
	Properties map[string]string
	Links []Link
}

func (receiver Response) MarshalJSON() ([]byte, error) {
	var buffer [256]byte
	var bytes []byte = buffer[0:0]

	var comma bool

	bytes = append(bytes, '{')

	{
		result, err := jsonMarshalNameOptionalValue("subject", receiver.Subject)
		if nil != err {
			return nil, err
		}
		if 0 < len(result) {
			if comma {
				bytes = append(bytes, ',')
			}
			comma = true
			bytes = append(bytes, result...)
		}
	}

	if 0 < len(receiver.Aliases) {
		result, err := jsonMarshalNameStringsValue("aliases", receiver.Aliases...)
		if nil != err {
			return nil, err
		}
		if 0 < len(result) {
			if comma {
				bytes = append(bytes, ',')
			}
			comma = true
			bytes = append(bytes, result...)
		}
	}

	if 0 < len(receiver.Properties) {
		const prefix string = `"properties":`

		result, err := jsonMarshalMap(receiver.Properties)
		if nil != err {
			return nil, err
		}
		if 0 < len(result) {
			if comma {
				bytes = append(bytes, ',')
			}
			comma = true
			bytes = append(bytes, prefix...)
			bytes = append(bytes, result...)
		}
	}

	if 0 < len(receiver.Links) {
		const prefix string = `"links":`

		result, err := json.Marshal(receiver.Links)
		if nil != err {
			return nil, err
		}
		if 0 < len(result) {
			if comma {
				bytes = append(bytes, ',')
			}
			comma = true
			bytes = append(bytes, prefix...)
			bytes = append(bytes, result...)
		}
	}

	bytes = append(bytes, '}')

	return bytes, nil
}
