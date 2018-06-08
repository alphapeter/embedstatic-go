package main

import (
	"mime"
	"net/http"
	"strings"
)

type EmbeddedHandler struct {
	data            map[string][]byte
	defaultResource string
}

func CreateHandler(data map[string][]byte, defaultResource string) EmbeddedHandler {
	return EmbeddedHandler{
		data:            data,
		defaultResource: defaultResource,
	}
}

func (h *EmbeddedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	resource := strings.TrimLeft(r.RequestURI, "/")
	if len(resource) == 0 {
		resource = h.defaultResource
	}
	d, found := getData(h.data, resource)

	if !found {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found"))
		return
	}

	contentType := getContentType(resource)
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusOK)
	w.Write(d)

	return

}

func getData(data map[string][]byte, resource string) ([]byte, bool) {
	_, found := data[resource]
	if !found {
		resource = strings.TrimRight(resource, "/") + "/index.html"
	}
	d, found := data[resource]
	return d, found
}

func getContentType(resourceName string) string {
	lastDot := strings.LastIndex(resourceName, ".")
	if lastDot > -1 {
		extension := resourceName[lastDot:]
		return mime.TypeByExtension(strings.ToLower(extension))
	}
	return mime.TypeByExtension("application/octet-stream")
}
