package nrg

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

func (c *Context) JSON(code int, obj interface{}) {
	c.W.WriteHeader(code)

	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		c.Error(err)
	}

	if _, err = c.W.Write(jsonBytes); err != nil {
		c.Error(err)
	}
}

func (c *Context) Error(err error) {
	c.W.WriteHeader(http.StatusInternalServerError)
	_, _ = c.W.Write([]byte(err.Error()))
}
