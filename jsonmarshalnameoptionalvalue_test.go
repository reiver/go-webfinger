package webfinger

import (
	"testing"

	"bytes"

	"github.com/reiver/go-opt"
)

func TestJSONMarshalNameOptionalValue(t *testing.T) {
	tests := []struct{
		Name string
		Value opt.Optional[string]
		Expected []byte
	}{
		{
		},



		{
			Name: "rel",
		},
		{
			Name: "rel",
			Value: opt.Something(""),
			Expected: []byte(`"rel":""`),
		},
		{
			Name: "rel",
			Value: opt.Something("author"),
			Expected: []byte(`"rel":"author"`),
		},
		{
			Name: "rel",
			Value: opt.Something("http://example.com/rel/something"),
			Expected: []byte(`"rel":"http://example.com/rel/something"`),
		},



		{
			Name: "type",
		},
		{
			Name: "type",
			Value: opt.Something(""),
			Expected: []byte(`"type":""`),
		},
		{
			Name: "type",
			Value: opt.Something("text/html"),
			Expected: []byte(`"type":"text/html"`),
		},
		{
			Name: "type",
			Value: opt.Something("application/activity+json"),
			Expected: []byte(`"type":"application/activity+json"`),
		},
	}

	for testNumber, test := range tests {

		actual, err := jsonMarshalNameOptionalValue(test.Name, test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			continue
		}

		{
			expected := test.Expected

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual json-marshaled name-value is not what was expected.", testNumber)
				t.Logf("EXPECTED: (%d) %q", len(expected), expected)
				t.Logf("ACTUAL:   (%d) %q", len(actual), actual)
				t.Logf("NAME: %q", test.Name)
				t.Logf("VALUE: %#v", test.Value)
				continue
			}
		}
	}
}
