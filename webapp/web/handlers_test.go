package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_application_handler(t *testing.T) {
	var handleTest = []struct {
		name               string
		url                string
		expectedStatusCode int
	}{
		{"home", "/", http.StatusOK},
		{"nothome", "/oyo", http.StatusNotFound},
	}

	var app application
	routes := app.routes()

	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, hand := range handleTest {
		resp, err := ts.Client().Get(ts.URL + hand.url)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}

		if resp.StatusCode != hand.expectedStatusCode {
			t.Errorf("for %s expected status is %d bit got %d", hand.name, hand.expectedStatusCode, resp.StatusCode)
		}
	}
}
