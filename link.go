package webfinger

import (
	"encoding/json"

	"github.com/reiver/go-opt"
)

// Link represents a JSON Resource Descriptor (JRD) 'link'.
//
// Some examples of a JRD lin include:
//
//	{
//		"rel" : "http://webfinger.example/rel/profile-page",
//		"href" : "https://www.example.com/~dariush/"
//	}
//
// And:
//
//	{
//		"rel": "self",
//		"type": "application/activity+json",
//		"href": "https://mastodon.social/users/reiver"
//	},
//
// And:
//
//	{
//		"rel": "http://ostatus.org/schema/1.0/subscribe",
//		"template": "https://social.example/authorize_interaction?uri={uri}"
//	},
//
// And:
//
//	{
//		"rel" : "author",
//		"href" : "http://blog.example.com/author/malekeh",
//		"titles" :
//		{
//			"en-us":"Hello world!",
//			"fa":"جهان درود",
//		},
//		"properties" :
//		{
//			"http://example.com/role" : "editor"
//		}
//	}
//
// It is used with [Response].
type Link struct {
	Rel  opt.Optional[string]    `json:"rel"`
	Type opt.Optional[string]    `json:"type"`
	HRef opt.Optional[string]    `json:"href"`
	Titles map[string]string     `json:"titles"`
	Properties map[string]string `json:"properties"`
}

var _ json.Marshaler = Link{}

func (receiver Link) MarshalJSON() ([]byte, error) {
	var buffer [256]byte
	var bytes []byte = buffer[0:0]

	bytes = append(bytes, '{')

	var comma bool

	{
		result, err := jsonMarshalNameOptionalValue("rel", receiver.Rel)
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


	{
		result, err := jsonMarshalNameOptionalValue("type", receiver.Type)
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

	{
		result, err := jsonMarshalNameOptionalValue("href", receiver.HRef)
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

	if 0 < len(receiver.Titles) {
		const prefix string = `"titles":`

		if comma {
			bytes = append(bytes, ',')
		}
		comma = true

		bytes = append(bytes, prefix...)

		{
			result, err := jsonMarshalMap(receiver.Titles)
			if nil != err {
				return nil, err
			}

			bytes = append(bytes, result...)
		}
	}

	if 0 < len(receiver.Properties) {
		const prefix string = `"properties":`

		if comma {
			bytes = append(bytes, ',')
		}
		comma = true

		bytes = append(bytes, prefix...)

		{
			result, err := jsonMarshalMap(receiver.Properties)
			if nil != err {
				return nil, err
			}

			bytes = append(bytes, result...)
		}
	}

	bytes = append(bytes, '}')

	return bytes, nil
}
