package webfinger_test

import (
	"testing"

	"bytes"
	"encoding/json"

	"github.com/reiver/go-opt"

	"github.com/reiver/go-webfinger"
)

func TestResponse_MarshalJSON(t *testing.T) {

	tests := []struct{
		Response webfinger.Response
		Expected []byte
	}{
		{
			Expected: []byte("{}"),
		},



		{
			Response: webfinger.Response{
				Subject: opt.Something("acct:joeblow@example.com"),
			},
			Expected: []byte(`{"subject":"acct:joeblow@example.com"}`),
		},
		{
			Response: webfinger.Response{
				Subject: opt.Something("http://example.com/~joeblow"),
			},
			Expected: []byte(`{"subject":"http://example.com/~joeblow"}`),
		},



		{
			Response: webfinger.Response{
				Subject: opt.Something("acct:joeblow@example.com"),
				Aliases: []string{
					"mailto:joeblow@example.com",
				},
			},
			Expected: []byte(`{"subject":"acct:joeblow@example.com","aliases":["mailto:joeblow@example.com"]}`),
		},
		{
			Response: webfinger.Response{
				Subject: opt.Something("acct:joeblow@example.com"),
				Aliases: []string{
					"mailto:joeblow@example.com",
					"http://example.com/~joeblow",
				},
			},
			Expected: []byte(`{"subject":"acct:joeblow@example.com","aliases":["mailto:joeblow@example.com","http://example.com/~joeblow"]}`),
		},
		{
			Response: webfinger.Response{
				Subject: opt.Something("acct:joeblow@example.com"),
				Aliases: []string{
					"mailto:joeblow@example.com",
					"http://example.com/~joeblow",
					"https://social.example/@joeblow",
				},
			},
			Expected: []byte(`{"subject":"acct:joeblow@example.com","aliases":["mailto:joeblow@example.com","http://example.com/~joeblow","https://social.example/@joeblow"]}`),
		},



		{
			Response: webfinger.Response{
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
			Expected: []byte(`{"subject":"acct:joeblow@example.com","aliases":["mailto:joeblow@example.com","http://example.com/~joeblow","https://social.example/@joeblow"],"properties":{"ONE":"apple"}}`),
		},
		{
			Response: webfinger.Response{
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
			Expected: []byte(`{"subject":"acct:joeblow@example.com","aliases":["mailto:joeblow@example.com","http://example.com/~joeblow","https://social.example/@joeblow"],"properties":{"ONE":"apple","TWO":"banana"}}`),
		},
		{
			Response: webfinger.Response{
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
			Expected: []byte(`{"subject":"acct:joeblow@example.com","aliases":["mailto:joeblow@example.com","http://example.com/~joeblow","https://social.example/@joeblow"],"properties":{"ONE":"apple","THREE":"cherry","TWO":"banana"}}`),
		},



		{
			Response: webfinger.Response{
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
			Expected: []byte(`{"subject":"acct:joeblow@example.com","aliases":["mailto:joeblow@example.com","http://example.com/~joeblow","https://social.example/@joeblow"],"properties":{"ONE":"apple","THREE":"cherry","TWO":"banana"},"links":[{"href":"http://example.com/abc.txt"}]}`),
		},
		{
			Response: webfinger.Response{
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
			Expected: []byte(`{"subject":"acct:joeblow@example.com","aliases":["mailto:joeblow@example.com","http://example.com/~joeblow","https://social.example/@joeblow"],"properties":{"ONE":"apple","THREE":"cherry","TWO":"banana"},"links":[{"href":"http://example.com/abc.txt"},{"rel":"me","href":"http://social.example/@joeblow"}]}`),
		},
		{
			Response: webfinger.Response{
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
			Expected: []byte(`{"subject":"acct:joeblow@example.com","aliases":["mailto:joeblow@example.com","http://example.com/~joeblow","https://social.example/@joeblow"],"properties":{"ONE":"apple","THREE":"cherry","TWO":"banana"},"links":[{"href":"http://example.com/abc.txt"},{"rel":"me","href":"http://social.example/@joeblow"},{"rel":"self","type":"application/activity+json","href":"http://social.example/user/joeblow"}]}`),
		},
	}

	for testNumber, test := range tests {

		actual, err := json.Marshal(test.Response)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("RESPONSE: %#v", test.Response)
			continue
		}

		{
			expected := test.Expected

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual marshaled-jsn is not what was expected.", testNumber)
				t.Logf("EXPECTED: (%d) %q", len(expected), expected)
				t.Logf("ACTUAL:   (%d) %q", len(actual), actual)
				t.Logf("RESPONSE: %#v", test.Response)
				continue
			}
		}
	}
}
