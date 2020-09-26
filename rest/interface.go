package rest

import "net/http"

type HttpClient interface {
	Put(url string, b []byte) (*http.Response, []byte, error)
	Post(url string, b []byte) (*http.Response, []byte, error)
	Get(url string) (*http.Response, []byte, error)
	GetWithHeader(url string, header *http.Header) (*http.Response, []byte, error)
	PutWithHeader(url string, b []byte, header *http.Header) (*http.Response, []byte, error)
	PostWithHeader(url string, b []byte, header *http.Header) (*http.Response, []byte, error)
}
