package middleware

import "net/http"

const (
	HeaderKeyContentType       = "Content-Type"
	HeaderValueContentTypeJSON = "application/json;charset=utf8"
)

// Sets the Content-Type header to application/json;charset=utf8
// It is used to ensure that the response is encoded in UTF-8
func ContentTypeJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(HeaderKeyContentType, HeaderValueContentTypeJSON)
		next.ServeHTTP(w, r)
	})
}
