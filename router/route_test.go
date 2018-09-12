package router

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

var (
	r          *mux.Router
	mockServer *httptest.Server
)

func init() {
	r = Init()
	mockServer = httptest.NewServer(r)
}

func TestRouterNonExistRoute(t *testing.T) {
	resp, err := http.Get(mockServer.URL + "/")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("status code should be 404, got %d", resp.StatusCode)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	respString := string(b)
	expected := "404 page not found\n"

	if respString != expected {
		t.Errorf("unexpected response, should be %s, got %s", expected, respString)
	}
}
