package nrg

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(*Context)

type Nrg struct {
	handles map[string]HandlerFunc
}

func (n *Nrg) GET(path string, handle HandlerFunc) *Nrg {
	return n.addRoute(http.MethodGet, path, handle)
}

func (n *Nrg) POST(path string, handle HandlerFunc) *Nrg {
	return n.addRoute(http.MethodPost, path, handle)
}

func (n *Nrg) DELETE(path string, handle HandlerFunc) *Nrg {
	return n.addRoute(http.MethodDelete, path, handle)
}

func (n *Nrg) PATCH(path string, handle HandlerFunc) *Nrg {
	return n.addRoute(http.MethodPatch, path, handle)
}

func (n *Nrg) PUT(path string, handle HandlerFunc) *Nrg {
	return n.addRoute(http.MethodPut, path, handle)
}

func (n *Nrg) OPTIONS(path string, handle HandlerFunc) *Nrg {
	return n.addRoute(http.MethodOptions, path, handle)
}

func (n *Nrg) HEAD(path string, handle HandlerFunc) *Nrg {
	return n.addRoute(http.MethodHead, path, handle)
}

func (n *Nrg) addRoute(method, path string, handle HandlerFunc) *Nrg {
	if n.handles == nil {
		n.handles = make(map[string]HandlerFunc)
	}
	n.handles[method+"-"+path] = handle
	return n
}

func (n *Nrg) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	method := req.Method
	key := method + "-" + path

	context := &Context{
		W: w,
		R: req,
	}
	if handle, ok := n.handles[key]; ok {
		handle(context)
	} else {
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprintf(w, "404 NOT FOUND: %s \n", req.URL)
	}
}

func (n *Nrg) Run(addr ...string) error {
	address := getServerAddress(addr)
	fmt.Printf("Listening and serving HTTP on %s\n", address)
	err := http.ListenAndServe(address, n)
	return err
}

func getServerAddress(addr []string) string {
	if len(addr) > 0 {
		return addr[0]
	}
	return ":8080"
}

func NewServer() *Nrg {
	server := &Nrg{}
	return server
}
