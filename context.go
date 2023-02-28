package nrg

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type Context struct {
	W http.ResponseWriter
	R *http.Request

	queryCache url.Values
}

func (c *Context) GetQuery(key string) (string, bool) {
	if values, ok := c.getQueryArray(key); ok {
		return values[0], true
	}

	return "", false
}

func (c *Context) getQueryArray(key string) ([]string, bool) {
	if c.queryCache == nil {
		if c.R != nil {
			c.queryCache = c.R.URL.Query()
		} else {
			c.queryCache = url.Values{}
		}
	}

	values, ok := c.queryCache[key]
	return values, ok
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
