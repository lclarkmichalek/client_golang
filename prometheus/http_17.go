// +build !go1.8

package prometheus

import (
	"io"
	"net/http"
)

func newResponseWriterDelegator(w http.ResponseWriter, delegate *responseWriteDelegator) http.ResponseWriter {
	_, cn := w.(http.CloseNotifier)
	_, fl := w.(http.Flusher)
	_, hj := w.(http.Hijacker)
	_, rf := w.(io.ReaderFrom)
	if cn && fl && hj && rf {
		return &httpResponseWriterDelegator{delegate}
	}
	return delegate
}
