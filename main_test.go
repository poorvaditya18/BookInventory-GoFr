package main

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
	"gofr.dev/pkg/gofr/request"
	"net/url"
)

func Test_Integration(t *testing.T) {
	go main()

	time.Sleep(5 * time.Second)

	createBook := url.Values{}
	createBook.Set("Title", "TestBook")
	createBook.Set("Author", "TestAuthor")
	createBook.Set("Price", "50")
	createBook.Set("QuantityAvailable", "10")

	

	updateBook := url.Values{}
	updateBook.Set("Title", "TestBook")
	updateBook.Set("QuantityAvailable", "100")

	// bookCreateBody := []byte(`{"Title": "TestBook", "Author": "TestAuthor", "Price": "50", "QuantityAvailable": "10"}`)

	// bookUpdateBody := []byte(`{"Title": "TestBook, "QuantityAvailable":30}`)

	bookCreateBody := []byte(createBook.Encode())
	bookUpdateBody := []byte(updateBook.Encode())

	successResp := `{"data":"Successfully Created Book"}`
	successUpdateResp := `{"data":"Successfully Updated Quantity Of Book"}`

	testCases := []struct {
		desc          string
		method        string
		endpoint      string
		body          []byte
		expStatusCode int
		expResp       string
	}{
		{"Create book", http.MethodPost, "/book", bookCreateBody, http.StatusCreated,
			successResp},
		// {"Get book", http.MethodGet, "/book/1", nil, http.StatusOK, successResp},
		{"Update book", http.MethodPatch, "/book", bookUpdateBody, http.StatusOK,
			successUpdateResp},
		// {"Delete book", http.MethodDelete, "/employee/1", nil, http.StatusNoContent, ``},
	}

	for i, tc := range testCases {

		t.Run(tc.desc, func(t *testing.T) {

			req, _ := request.NewMock(tc.method, "http://localhost:3000"+tc.endpoint, bytes.NewBuffer(tc.body))
			client := http.Client{}

			resp, err := client.Do(req)
			if err != nil {
				t.Fatalf("Error occurred in calling api: %v", err)
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("Error while reading response: %v", err)
			}

			respBody := strings.TrimSpace(string(body))

			assert.Equal(t, tc.expStatusCode, resp.StatusCode, "Test [%d] failed", i+1)
			assert.Equal(t, tc.expResp, respBody, "Test [%d] failed", i+1)

			resp.Body.Close()
		})
	}
}