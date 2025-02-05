package webfinger_test

import (
	"testing"

	"encoding/json"
	"reflect"

	"github.com/reiver/go-opt"

	"github.com/reiver/go-webfinger"
)

func TestLink_UnmarshalJSON(t *testing.T) {
	tests := []struct{
		Data []byte
		Expected webfinger.Link
	}{
		{
			Data: []byte(`{}`),
		},



		{
			Expected: webfinger.Link{
				Rel: opt.Something(""),
			},
			Data: []byte(`{"rel":""}`),
		},
		{
			Expected: webfinger.Link{
				Type: opt.Something(""),
			},
			Data: []byte(`{"type":""}`),
		},
		{
			Expected: webfinger.Link{
				HRef: opt.Something(""),
			},
			Data: []byte(`{"href":""}`),
		},



		{
			Expected: webfinger.Link{
				Rel: opt.Something("attachment"),
			},
			Data: []byte(`{"rel":"attachment"}`),
		},
		{
			Expected: webfinger.Link{
				Type: opt.Something("application/activity+json"),
			},
			Data: []byte(`{"type":"application/activity+json"}`),
		},
		{
			Expected: webfinger.Link{
				HRef: opt.Something("http://example.com/apple/banana/cherry.txt"),
			},
			Data: []byte(`{"href":"http://example.com/apple/banana/cherry.txt"}`),
		},



		{
			Expected: webfinger.Link{
				Type: opt.Something("application/activity+json"),
				HRef: opt.Something("http://example.com/apple/banana/cherry.txt"),
			},
			Data: []byte(`{"type":"application/activity+json","href":"http://example.com/apple/banana/cherry.txt"}`),
		},
		{
			Expected: webfinger.Link{
				Rel: opt.Something("attachment"),
				HRef: opt.Something("http://example.com/apple/banana/cherry.txt"),
			},
			Data: []byte(`{"rel":"attachment","href":"http://example.com/apple/banana/cherry.txt"}`),
		},
		{
			Expected: webfinger.Link{
				Rel: opt.Something("attachment"),
				Type: opt.Something("application/activity+json"),
			},
			Data: []byte(`{"rel":"attachment","type":"application/activity+json"}`),
		},



		{
			Expected: webfinger.Link{
				Rel: opt.Something("attachment"),
				Type: opt.Something("application/activity+json"),
				HRef: opt.Something("http://example.com/apple/banana/cherry.txt"),
			},
			Data: []byte(`{"rel":"attachment","type":"application/activity+json","href":"http://example.com/apple/banana/cherry.txt"}`),
		},



		{
			Expected: webfinger.Link{
				Titles: map[string]string{
					"en":"Hello world!",
				},
			},
			Data: []byte(`{"titles":{"en":"Hello world!"}}`),
		},
		{
			Expected: webfinger.Link{
				Titles: map[string]string{
					"fa":"جهان درود",
				},
			},
			Data: []byte(`{"titles":{"fa":"جهان درود"}}`),
		},
		{
			Expected: webfinger.Link{
				Titles: map[string]string{
					"en":"Hello world!",
					"fa":"جهان درود",
				},
			},
			Data: []byte(`{"titles":{"en":"Hello world!","fa":"جهان درود"}}`),
		},



		{
			Expected: webfinger.Link{
				Rel: opt.Something("attachment"),
				Type: opt.Something("application/activity+json"),
				HRef: opt.Something("http://example.com/apple/banana/cherry.txt"),
				Titles: map[string]string{
					"en":"Hello world!",
				},
			},
			Data: []byte(`{"rel":"attachment","type":"application/activity+json","href":"http://example.com/apple/banana/cherry.txt","titles":{"en":"Hello world!"}}`),
		},
		{
			Expected: webfinger.Link{
				Rel: opt.Something("attachment"),
				Type: opt.Something("application/activity+json"),
				HRef: opt.Something("http://example.com/apple/banana/cherry.txt"),
				Titles: map[string]string{
					"en":"Hello world!",
					"fa":"جهان درود",
				},
			},
			Data: []byte(`{"rel":"attachment","type":"application/activity+json","href":"http://example.com/apple/banana/cherry.txt","titles":{"en":"Hello world!","fa":"جهان درود"}}`),
		},



		{
			Expected: webfinger.Link{
				Rel: opt.Something("attachment"),
				Type: opt.Something("application/activity+json"),
				HRef: opt.Something("http://example.com/apple/banana/cherry.txt"),
				Titles: map[string]string{
					"en":"Hello world!",
					"fa":"جهان درود",
				},
				Properties: map[string]string{
					"ONE":"apple",
				},
			},
			Data: []byte(`{"rel":"attachment","type":"application/activity+json","href":"http://example.com/apple/banana/cherry.txt","titles":{"en":"Hello world!","fa":"جهان درود"},"properties":{"ONE":"apple"}}`),
		},
		{
			Expected: webfinger.Link{
				Rel: opt.Something("attachment"),
				Type: opt.Something("application/activity+json"),
				HRef: opt.Something("http://example.com/apple/banana/cherry.txt"),
				Titles: map[string]string{
					"en":"Hello world!",
					"fa":"جهان درود",
				},
				Properties: map[string]string{
					"ONE":"apple",
					"TWO":"banana",
				},
			},
			Data: []byte(`{"rel":"attachment","type":"application/activity+json","href":"http://example.com/apple/banana/cherry.txt","titles":{"en":"Hello world!","fa":"جهان درود"},"properties":{"ONE":"apple","TWO":"banana"}}`),
		},
		{
			Expected: webfinger.Link{
				Rel: opt.Something("attachment"),
				Type: opt.Something("application/activity+json"),
				HRef: opt.Something("http://example.com/apple/banana/cherry.txt"),
				Titles: map[string]string{
					"en":"Hello world!",
					"fa":"جهان درود",
				},
				Properties: map[string]string{
					"ONE":"apple",
					"TWO":"banana",
					"THREE":"cherry",
				},
			},
			Data: []byte(`{"rel":"attachment","type":"application/activity+json","href":"http://example.com/apple/banana/cherry.txt","titles":{"en":"Hello world!","fa":"جهان درود"},"properties":{"ONE":"apple","THREE":"cherry","TWO":"banana"}}`),
		},
	}

	for testNumber, test := range tests {

		var actual webfinger.Link
		err := json.Unmarshal(test.Data, &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("DATA:\n%s", test.Data)
			continue
		}

		{
			expected := test.Expected

			if !reflect.DeepEqual(expected, actual) {
				t.Errorf("For test #%d, the actual json-marshaled name-value is not what was expected.", testNumber)
				t.Logf("EXPECTED:\n%#v", expected)
				t.Logf("ACTUAL:\n%#v", actual)
				t.Logf("DATA:\n%s", test.Data)
				continue
			}
		}
	}
}
