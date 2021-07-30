package fw

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	writer  http.ResponseWriter
	request *http.Request
	params  map[string]string
}

/*
	Request
*/
func (fwc *Context) Query(key string) string {
	return fwc.request.URL.Query().Get(key)
}

func (fwc *Context) Param(key string) string {
	return fwc.params[key]
}

func (fwc *Context) Body(v interface{}) error {
	return json.NewDecoder(fwc.request.Body).Decode(v)
}

/*
	ResponseWriter
*/
func (fwc *Context) Header(key string) string {
	return fwc.writer.Header().Get(key)
}

func (fwc *Context) Write(res string) (int, error) {
	fwc.writer.Header().Set("Content-Type", "text/plain")
	return fwc.writer.Write([]byte(res))
}

func (fwc *Context) JSON(v interface{}) error {
	fwc.writer.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(fwc.writer).Encode(v)
}
