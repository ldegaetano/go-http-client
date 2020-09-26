package rest

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type RoundTripFunc func(req *http.Request) (*http.Response, error)

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

//NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) http.Client {
	return http.Client{
		Transport: RoundTripFunc(fn),
	}
}

type mockReadCloser struct {
	mock.Mock
}

func (m *mockReadCloser) Read(p []byte) (n int, err error) {
	args := m.Called(p)
	return args.Int(0), args.Error(1)
}

func (m *mockReadCloser) Close() error {
	args := m.Called()
	return args.Error(0)
}

func TestScuccesPut(t *testing.T) {
	rtID := 543
	expectedURL := fmt.Sprintf("localhost/api/sr-orchestrator/rappitenderos/%v/courses", rtID)

	testCl := NewTestClient(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, expectedURL, req.URL.String())
		assert.Equal(t, "PUT", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`OK`)),
			Header:     make(http.Header),
		}, nil
	})
	cl := NewClient(ClientCfg{}).(client)
	cl.Client = testCl

	resp, body, err := cl.Put(expectedURL, []byte("{}"))

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "OK", string(body))
	assert.Nil(t, err, "Unexpected error in PUT")
}

func TestFailPut(t *testing.T) {
	testCl := NewTestClient(func(req *http.Request) (*http.Response, error) {
		return &http.Response{}, errors.New("Test error transport")
	})
	cl := NewClient(ClientCfg{}).(client)
	cl.Client = testCl

	_, _, err := cl.Put("", []byte("{}"))

	assert.True(t, strings.Contains(err.Error(), "Test error transport"))
}

func TestErrorBody(t *testing.T) {
	mockReadCloser := mockReadCloser{}
	mockReadCloser.On("Read", mock.AnythingOfType("[]uint8")).Return(0, fmt.Errorf("Test error reading"))
	mockReadCloser.On("Close").Return(nil)

	rtID := 543
	expectedURL := fmt.Sprintf("localhost/api/sr-orchestrator/rappitenderos/%v/courses", rtID)

	testCl := NewTestClient(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, expectedURL, req.URL.String())
		return &http.Response{
			StatusCode: 200,
			Body:       &mockReadCloser,
			Header:     make(http.Header),
		}, nil
	})
	cl := NewClient(ClientCfg{}).(client)
	cl.Client = testCl

	resp, body, err := cl.Put(expectedURL, []byte("{}"))

	assert.Equal(t, "Test error reading", err.Error())
	assert.Nil(t, resp, "Resp should be nil")
	assert.Nil(t, body, "Body should be nil")
}

func TestScuccesPost(t *testing.T) {
	expectedURL := "localhost/api/sr-skills/summarized"

	testCl := NewTestClient(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, expectedURL, req.URL.String())
		assert.Equal(t, "POST", req.Method)
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`OK`)),
			Header:     make(http.Header),
		}, nil
	})
	cl := NewClient(ClientCfg{}).(client)
	cl.Client = testCl

	resp, body, err := cl.Post(expectedURL, []byte("{}"))

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "OK", string(body))
	assert.Nil(t, err, "Unexpected error in PUT")
}

func TestErrorStatus500(t *testing.T) {
	rtID := 543
	expectedURL := fmt.Sprintf("localhost/api/sr-orchestrator/rappitenderos/%v/courses", rtID)

	numberOfCalls := 0
	testCl := NewTestClient(func(req *http.Request) (*http.Response, error) {
		numberOfCalls++
		assert.Equal(t, expectedURL, req.URL.String())
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`Internal error`)),
			Header:     make(http.Header),
		}, nil
	})
	cl := NewClient(ClientCfg{
		RetryInterval: time.Millisecond,
		RetryAttempts: 3,
	}).(client)
	cl.Client = testCl

	_, _, err := cl.Put(expectedURL, []byte("{}"))

	assert.Equal(t, "Error in status code", err.Error())
	assert.Equal(t, 3, numberOfCalls)
}
