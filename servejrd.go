package webfinger

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/reiver/go-etag"
)

func ServeJRDBytes(responseWriter http.ResponseWriter, request *http.Request, jrd []byte) {

	var digest [sha256.Size]byte
	{
		digest = sha256.Sum256(jrd)
	}

	var cacheDigest string
	{
		cacheDigest = fmt.Sprintf("sha-256=:%s:", base64.StdEncoding.EncodeToString(digest[:]))
	}

	{
		const contentType string = "application/jrd+json"

		var header http.Header = responseWriter.Header()
		if nil == header {
			httpError(responseWriter, http.StatusInternalServerError)
			return
		}

		header.Add("Access-Control-Allow-Origin", "*")
		header.Add("Cache-Control", "max-age=907")
		header.Add("Content-Digest", cacheDigest)
		header.Add("Content-Type", contentType)
	}

	var eTag string
	{
		var format string = fmt.Sprintf("sha256=0x%%0%dX", sha256.Size*2)
		eTag = fmt.Sprintf(format, digest[:])
	}

	{
		var handled bool = etag.Handle(responseWriter, request, eTag)
		if handled {
			return
		}
	}

	{
		responseWriter.WriteHeader(http.StatusOK)
		responseWriter.Write(jrd)
	}
}
