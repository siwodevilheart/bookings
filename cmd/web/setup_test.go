package main

import "net/http"
import "os"
import "testing"

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

type myHandler struct {
}

func (mh *myHandler) ServeHTTP(w http.ResponseWriter, h *http.Request) {

}
