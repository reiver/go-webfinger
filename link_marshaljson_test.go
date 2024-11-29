package webfinger_test

import (
	"testing"

	"bytes"
	"encoding/json"

	"github.com/reiver/go-opt"

	"github.com/reiver/go-webfinger"
)

func TestLink_MarshalJSON(t *testing.T) {
	tests := []struct{
		Link webfinger.Link
		Expected []byte
	}{
		{
			Expected: []byte(`{}`),
		},



		{
			Link: webfinger.Link{
				Rel: opt.Something(""),
			},
			Expected: []byte(`{"rel":""}`),
		},
		{
			Link: webfinger.Link{
				Type: opt.Something(""),
			},
			Expected: []byte(`{"type":""}`),
		},
		{
			Link: webfinger.Link{
				HRef: opt.Something(""),
			},
			Expected: []byte(`{"href":""}`),
		},



		{
			Link: webfinger.Link{
				Rel: opt.Something("attachment"),
			},
			Expected: []byte(`{"rel":"attachment"}`),
		},
		{
			Link: webfinger.Link{
				Type: opt.Something("application/activity+json"),
			},
			Expected: []byte(`{"type":"application/activity+json"}`),
		},
		{
			Link: webfinger.Link{
				HRef: opt.Something("http://example.com/apple/banana/cherry.txt"),
			},
			Expected: []byte(`{"href":"http://example.com/apple/banana/cherry.txt"}`),
		},



		{
			Link: webfinger.Link{
				Type: opt.Something("application/activity+json"),
				HRef: opt.Something("http://example.com/apple/banana/cherry.txt"),
			},
			Expected: []byte(`{"type":"application/activity+json","href":"http://example.com/apple/banana/cherry.txt"}`),
		},
		{
			Link: webfinger.Link{
				Rel: opt.Something("attachment"),
				HRef: opt.Something("http://example.com/apple/banana/cherry.txt"),
			},
			Expected: []byte(`{"rel":"attachment","href":"http://example.com/apple/banana/cherry.txt"}`),
		},
		{
			Link: webfinger.Link{
				Rel: opt.Something("attachment"),
				Type: opt.Something("application/activity+json"),
			},
			Expected: []byte(`{"rel":"attachment","type":"application/activity+json"}`),
		},



		{
			Link: webfinger.Link{
				Rel: opt.Something("attachment"),
				Type: opt.Something("application/activity+json"),
				HRef: opt.Something("http://example.com/apple/banana/cherry.txt"),
			},
			Expected: []byte(`{"rel":"attachment","type":"application/activity+json","href":"http://example.com/apple/banana/cherry.txt"}`),
		},



		{
			Link: webfinger.Link{
				Titles: map[string]string{
					"en":"Hello world!",
				},
			},
			Expected: []byte(`{"titles":{"en":"Hello world!"}}`),
		},
		{
			Link: webfinger.Link{
				Titles: map[string]string{
					"fa":"جهان درود",
				},
			},
			Expected: []byte(`{"titles":{"fa":"جهان درود"}}`),
		},
		{
			Link: webfinger.Link{
				Titles: map[string]string{
					"en":"Hello world!",
					"fa":"جهان درود",
				},
			},
			Expected: []byte(`{"titles":{"en":"Hello world!","fa":"جهان درود"}}`),
		},



		{
			Link: webfinger.Link{
				Rel: opt.Something("attachment"),
				Type: opt.Something("application/activity+json"),
				HRef: opt.Something("http://example.com/apple/banana/cherry.txt"),
				Titles: map[string]string{
					"en":"Hello world!",
				},
			},
			Expected: []byte(`{"rel":"attachment","type":"application/activity+json","href":"http://example.com/apple/banana/cherry.txt","titles":{"en":"Hello world!"}}`),
		},
		{
			Link: webfinger.Link{
				Rel: opt.Something("attachment"),
				Type: opt.Something("application/activity+json"),
				HRef: opt.Something("http://example.com/apple/banana/cherry.txt"),
				Titles: map[string]string{
					"en":"Hello world!",
					"fa":"جهان درود",
				},
			},
			Expected: []byte(`{"rel":"attachment","type":"application/activity+json","href":"http://example.com/apple/banana/cherry.txt","titles":{"en":"Hello world!","fa":"جهان درود"}}`),
		},



		{
			Link: webfinger.Link{
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
			Expected: []byte(`{"rel":"attachment","type":"application/activity+json","href":"http://example.com/apple/banana/cherry.txt","titles":{"en":"Hello world!","fa":"جهان درود"},"properties":{"ONE":"apple"}}`),
		},
		{
			Link: webfinger.Link{
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
			Expected: []byte(`{"rel":"attachment","type":"application/activity+json","href":"http://example.com/apple/banana/cherry.txt","titles":{"en":"Hello world!","fa":"جهان درود"},"properties":{"ONE":"apple","TWO":"banana"}}`),
		},
		{
			Link: webfinger.Link{
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
			Expected: []byte(`{"rel":"attachment","type":"application/activity+json","href":"http://example.com/apple/banana/cherry.txt","titles":{"en":"Hello world!","fa":"جهان درود"},"properties":{"ONE":"apple","THREE":"cherry","TWO":"banana"}}`),
		},
	}

	for testNumber, test := range tests {

		actual, err := json.Marshal(test.Link)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("LINK: %#v", test.Link)
			continue
		}

		{
			expected := test.Expected

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual json-marshaled name-value is not what was expected.", testNumber)
				t.Logf("EXPECTED: (%d) %q", len(expected), expected)
				t.Logf("ACTUAL:   (%d) %q", len(actual), actual)
				t.Logf("LINK: %#v", test.Link)
				continue
			}
		}
	}
}
