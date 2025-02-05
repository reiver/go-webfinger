package webfinger_test

import (
	"testing"

	"encoding/json"
	"reflect"

	"github.com/reiver/go-opt"

	"github.com/reiver/go-webfinger"
)

func TestResponse_UnmarshalJSON(t *testing.T) {

	tests := []struct{
		Data []byte
		Expected webfinger.Response
	}{
		{
			Data: []byte("{}"),
		},



		{
			Expected: webfinger.Response{
				Subject: opt.Something("acct:joeblow@example.com"),
			},
			Data: []byte(`{"subject":"acct:joeblow@example.com"}`),
		},
		{
			Expected: webfinger.Response{
				Subject: opt.Something("http://example.com/~joeblow"),
			},
			Data: []byte(`{"subject":"http://example.com/~joeblow"}`),
		},



		{
			Expected: webfinger.Response{
				Subject: opt.Something("acct:joeblow@example.com"),
				Aliases: []string{
					"mailto:joeblow@example.com",
				},
			},
			Data: []byte(`{"subject":"acct:joeblow@example.com","aliases":["mailto:joeblow@example.com"]}`),
		},
		{
			Expected: webfinger.Response{
				Subject: opt.Something("acct:joeblow@example.com"),
				Aliases: []string{
					"mailto:joeblow@example.com",
					"http://example.com/~joeblow",
				},
			},
			Data: []byte(`{"subject":"acct:joeblow@example.com","aliases":["mailto:joeblow@example.com","http://example.com/~joeblow"]}`),
		},
		{
			Expected: webfinger.Response{
				Subject: opt.Something("acct:joeblow@example.com"),
				Aliases: []string{
					"mailto:joeblow@example.com",
					"http://example.com/~joeblow",
					"https://social.example/@joeblow",
				},
			},
			Data: []byte(`{"subject":"acct:joeblow@example.com","aliases":["mailto:joeblow@example.com","http://example.com/~joeblow","https://social.example/@joeblow"]}`),
		},



		{
			Expected: webfinger.Response{
				Subject: opt.Something("acct:joeblow@example.com"),
				Aliases: []string{
					"mailto:joeblow@example.com",
					"http://example.com/~joeblow",
					"https://social.example/@joeblow",
				},
				Properties: map[string]string{
					"ONE":"apple",
				},
			},
			Data: []byte(`{"subject":"acct:joeblow@example.com","aliases":["mailto:joeblow@example.com","http://example.com/~joeblow","https://social.example/@joeblow"],"properties":{"ONE":"apple"}}`),
		},
		{
			Expected: webfinger.Response{
				Subject: opt.Something("acct:joeblow@example.com"),
				Aliases: []string{
					"mailto:joeblow@example.com",
					"http://example.com/~joeblow",
					"https://social.example/@joeblow",
				},
				Properties: map[string]string{
					"ONE":"apple",
					"TWO":"banana",
				},
			},
			Data: []byte(`{"subject":"acct:joeblow@example.com","aliases":["mailto:joeblow@example.com","http://example.com/~joeblow","https://social.example/@joeblow"],"properties":{"ONE":"apple","TWO":"banana"}}`),
		},
		{
			Expected: webfinger.Response{
				Subject: opt.Something("acct:joeblow@example.com"),
				Aliases: []string{
					"mailto:joeblow@example.com",
					"http://example.com/~joeblow",
					"https://social.example/@joeblow",
				},
				Properties: map[string]string{
					"ONE":"apple",
					"TWO":"banana",
					"THREE":"cherry",
				},
			},
			Data: []byte(`{"subject":"acct:joeblow@example.com","aliases":["mailto:joeblow@example.com","http://example.com/~joeblow","https://social.example/@joeblow"],"properties":{"ONE":"apple","THREE":"cherry","TWO":"banana"}}`),
		},



		{
			Expected: webfinger.Response{
				Subject: opt.Something("acct:joeblow@example.com"),
				Aliases: []string{
					"mailto:joeblow@example.com",
					"http://example.com/~joeblow",
					"https://social.example/@joeblow",
				},
				Properties: map[string]string{
					"ONE":"apple",
					"TWO":"banana",
					"THREE":"cherry",
				},
				Links: []webfinger.Link{
					webfinger.Link{
						HRef:opt.Something("http://example.com/abc.txt"),
					},
				},
			},
			Data: []byte(`{"subject":"acct:joeblow@example.com","aliases":["mailto:joeblow@example.com","http://example.com/~joeblow","https://social.example/@joeblow"],"properties":{"ONE":"apple","THREE":"cherry","TWO":"banana"},"links":[{"href":"http://example.com/abc.txt"}]}`),
		},
		{
			Expected: webfinger.Response{
				Subject: opt.Something("acct:joeblow@example.com"),
				Aliases: []string{
					"mailto:joeblow@example.com",
					"http://example.com/~joeblow",
					"https://social.example/@joeblow",
				},
				Properties: map[string]string{
					"ONE":"apple",
					"TWO":"banana",
					"THREE":"cherry",
				},
				Links: []webfinger.Link{
					webfinger.Link{
						HRef:opt.Something("http://example.com/abc.txt"),
					},
					webfinger.Link{
						Rel:opt.Something("me"),
						HRef:opt.Something("http://social.example/@joeblow"),
					},
				},
			},
			Data: []byte(`{"subject":"acct:joeblow@example.com","aliases":["mailto:joeblow@example.com","http://example.com/~joeblow","https://social.example/@joeblow"],"properties":{"ONE":"apple","THREE":"cherry","TWO":"banana"},"links":[{"href":"http://example.com/abc.txt"},{"rel":"me","href":"http://social.example/@joeblow"}]}`),
		},
		{
			Expected: webfinger.Response{
				Subject: opt.Something("acct:joeblow@example.com"),
				Aliases: []string{
					"mailto:joeblow@example.com",
					"http://example.com/~joeblow",
					"https://social.example/@joeblow",
				},
				Properties: map[string]string{
					"ONE":"apple",
					"TWO":"banana",
					"THREE":"cherry",
				},
				Links: []webfinger.Link{
					webfinger.Link{
						HRef:opt.Something("http://example.com/abc.txt"),
					},
					webfinger.Link{
						Rel:opt.Something("me"),
						HRef:opt.Something("http://social.example/@joeblow"),
					},
					webfinger.Link{
						Rel:opt.Something("self"),
						Type:opt.Something("application/activity+json"),
						HRef:opt.Something("http://social.example/user/joeblow"),
					},
				},
			},
			Data: []byte(`{"subject":"acct:joeblow@example.com","aliases":["mailto:joeblow@example.com","http://example.com/~joeblow","https://social.example/@joeblow"],"properties":{"ONE":"apple","THREE":"cherry","TWO":"banana"},"links":[{"href":"http://example.com/abc.txt"},{"rel":"me","href":"http://social.example/@joeblow"},{"rel":"self","type":"application/activity+json","href":"http://social.example/user/joeblow"}]}`),
		},
	}

	for testNumber, test := range tests {

		var actual webfinger.Response
		err := json.Unmarshal(test.Data, &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("DATA: %s", test.Data)
			continue
		}

		{
			expected := test.Expected

			if !reflect.DeepEqual(expected, actual) {
				t.Errorf("For test #%d, the actual unmarshaled-json is not what was expected.", testNumber)
				t.Logf("EXPECTED: %#v", expected)
				t.Logf("ACTUAL:   %#v", actual)
				t.Logf("DATA: %s", test.Data)
				continue
			}
		}
	}
}
