package webfinger

import (
	"github.com/reiver/go-opt"
)

func jsonMarshalNameOptionalValue(name string, value opt.Optional[string]) ([]byte, error) {
	unwrapedValue, found := value.Get()
	if !found {
		return nil, nil
	}

	return jsonMarshalNameStringValue(name, unwrapedValue)
}
