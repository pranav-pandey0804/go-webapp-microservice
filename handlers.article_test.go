// handlers.article_test.go

package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Test that the application returns a JSON list of articles
// when the Accept header is set to application/json
func TestArticleListJSON(t *testing.T) {
	r := getRouter(true)

	r.GET("/", showIndexPage)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Accept", "application/json")

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK

		// Test that the response is JSON
		p, err := ioutil.ReadAll(w.Body)
		jsonOK := err == nil && strings.HasPrefix(string(p), "[{\"id\":1,\"title\":\"Article 1\",\"content\":\"Article 1 body\"}]")

		return statusOK && jsonOK
	})
}

// Test the application returns an article in XML format
// when the Accept header is set to application/xml
func TestArticleXML(t *testing.T) {
	r := getRouter(true)

	r.GET("/", showIndexPage)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Accept", "application/xml")

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK

		// Test that the response is XML
		p, err := ioutil.ReadAll(w.Body)
		xmlOK := err == nil && strings.HasPrefix(string(p), "<articles><article><ID>1</ID><Title>Article 1</Title><Content>Article 1 body</Content></article>")

		return statusOK && xmlOK
	})
}

// Test that a GET request to the home page returns the home page with
// the HTTP code 200 for an unauthenticated user
func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := getRouter(true)

	r.GET("/", showIndexPage)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK

		// Test that the page title is "Home Page"
		// we can carry out a lot more detailed tests using libraries that can
		// parse and process HTML pages
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOK && pageOK
	})
}
