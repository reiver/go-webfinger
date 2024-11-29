package webfinger

import (
	"slices"
)

func jsonMarshalMap(m map[string]string) ([]byte, error) {
	var buffer [256]byte
	var bytes []byte = buffer[0:0]


	bytes = append(bytes, '{')

	{
		var names []string
		{
			for name, _ := range m {
				names = append(names, name)
			}

			slices.Sort(names)
		}

		for index, name := range names {
			var value string = m[name]

			result, err := jsonMarshalNameStringValue(name, value)
			if nil != err {
				return nil, err
			}
			if 1 <= index {
				bytes = append(bytes, ',')
			}
			bytes = append(bytes, result...)
		}
	}

	bytes = append(bytes, '}')

	return bytes, nil
}
