// +build go1.8

package prometheus

import (
	"io"
	"net/http"
)

func newResponseWriterDelegator(w http.ResponseWriter, delegate *responseWriterDelegator) http.ResponseWriter {
	_, cn := w.(http.CloseNotifier)
	_, fl := w.(http.Flusher)
	_, hj := w.(http.Hijacker)
	_, rf := w.(io.ReaderFrom)
	_, ps := w.(http.Pusher)
	if cn && fl && hj && rf && !ps {
		return &httpResponseWriterDelegator{delegate}
	}
	if cn && fl && !hj && !rf && ps {
		return &http2ResponseWriterDelegator{delegate}
	}
	return delegate
}

type http2ResponseWriterDelegator struct {
	*responseWriterDelegator
}

func (h *http2ResponseWriterDelegator) CloseNotify() <-chan bool {
	return h.ResponseWriter.(http.CloseNotifier).CloseNotify()
}

func (h *http2ResponseWriterDelegator) Flush() {
	h.ResponseWriter.(http.Flusher).Flush()
}

func (h *http2ResponseWriterDelegator) Push(path string, opts *http.PushOptions) error {
	return h.ResponseWriter.(http.Pusher).Push(path, opts)
}
