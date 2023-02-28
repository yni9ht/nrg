package main

import (
	"log"
	"net/http"
)

type Nrg struct {
}

func (n *Nrg) ServeHTTP(w http.ResponseWriter, req *http.Request) {

}

func (n *Nrg) Run(addr ...string) error {
	address := getServerAddress(addr)
	log.Printf("Listening and serving HTTP on %s\n", address)
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

func main() {
	server := NewServer()

	if err := server.Run(); err != nil {
		log.Fatalln(err)
	}
}
