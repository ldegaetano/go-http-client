// Code generated by mockery v1.1.2. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// HttpClient is an autogenerated mock type for the HttpClient type
type HttpClient struct {
	mock.Mock
}

// Get provides a mock function with given fields: url
func (_m *HttpClient) Get(url string) (*http.Response, []byte, error) {
	ret := _m.Called(url)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(string) *http.Response); ok {
		r0 = rf(url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 []byte
	if rf, ok := ret.Get(1).(func(string) []byte); ok {
		r1 = rf(url)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(url)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetWithHeader provides a mock function with given fields: url, header
func (_m *HttpClient) GetWithHeader(url string, header *http.Header) (*http.Response, []byte, error) {
	ret := _m.Called(url, header)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(string, *http.Header) *http.Response); ok {
		r0 = rf(url, header)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 []byte
	if rf, ok := ret.Get(1).(func(string, *http.Header) []byte); ok {
		r1 = rf(url, header)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, *http.Header) error); ok {
		r2 = rf(url, header)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Post provides a mock function with given fields: url, b
func (_m *HttpClient) Post(url string, b []byte) (*http.Response, []byte, error) {
	ret := _m.Called(url, b)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(string, []byte) *http.Response); ok {
		r0 = rf(url, b)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 []byte
	if rf, ok := ret.Get(1).(func(string, []byte) []byte); ok {
		r1 = rf(url, b)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, []byte) error); ok {
		r2 = rf(url, b)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PostWithHeader provides a mock function with given fields: url, b, header
func (_m *HttpClient) PostWithHeader(url string, b []byte, header *http.Header) (*http.Response, []byte, error) {
	ret := _m.Called(url, b, header)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(string, []byte, *http.Header) *http.Response); ok {
		r0 = rf(url, b, header)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 []byte
	if rf, ok := ret.Get(1).(func(string, []byte, *http.Header) []byte); ok {
		r1 = rf(url, b, header)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, []byte, *http.Header) error); ok {
		r2 = rf(url, b, header)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Put provides a mock function with given fields: url, b
func (_m *HttpClient) Put(url string, b []byte) (*http.Response, []byte, error) {
	ret := _m.Called(url, b)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(string, []byte) *http.Response); ok {
		r0 = rf(url, b)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 []byte
	if rf, ok := ret.Get(1).(func(string, []byte) []byte); ok {
		r1 = rf(url, b)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, []byte) error); ok {
		r2 = rf(url, b)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PutWithHeader provides a mock function with given fields: url, b, header
func (_m *HttpClient) PutWithHeader(url string, b []byte, header *http.Header) (*http.Response, []byte, error) {
	ret := _m.Called(url, b, header)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(string, []byte, *http.Header) *http.Response); ok {
		r0 = rf(url, b, header)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 []byte
	if rf, ok := ret.Get(1).(func(string, []byte, *http.Header) []byte); ok {
		r1 = rf(url, b, header)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, []byte, *http.Header) error); ok {
		r2 = rf(url, b, header)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}