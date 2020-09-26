package rest

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/labstack/gommon/log"
)

type client struct {
	http.Client
	ClientCfg
}

type ClientCfg struct {
	RetryAttempts int
	RetryInterval time.Duration
	BasePath      string
	Timeout       time.Duration
}

func NewClient(cfg ClientCfg) HttpClient {
	c := http.Client{
		Timeout: cfg.Timeout,
	}
	return client{c, cfg}
}

func (c client) Put(url string, b []byte) (*http.Response, []byte, error) {
	return retry(c.ClientCfg.RetryAttempts, c.ClientCfg.RetryInterval, c.Do("PUT", c.ClientCfg.BasePath+url, b, nil))
}

func (c client) Post(url string, b []byte) (*http.Response, []byte, error) {
	return retry(c.ClientCfg.RetryAttempts, c.ClientCfg.RetryInterval, c.Do("POST", c.ClientCfg.BasePath+url, b, nil))
}

func (c client) Get(url string) (*http.Response, []byte, error) {
	b := []byte{}
	return retry(c.ClientCfg.RetryAttempts, c.ClientCfg.RetryInterval, c.Do("GET", c.ClientCfg.BasePath+url, b, nil))
}

func (c client) GetWithHeader(url string, header *http.Header) (*http.Response, []byte, error) {
	b := []byte{}
	return retry(c.ClientCfg.RetryAttempts, c.ClientCfg.RetryInterval, c.Do("GET", c.ClientCfg.BasePath+url, b, header))
}

func (c client) PutWithHeader(url string, b []byte, header *http.Header) (*http.Response, []byte, error) {
	return retry(c.ClientCfg.RetryAttempts, c.ClientCfg.RetryInterval, c.Do("PUT", c.ClientCfg.BasePath+url, b, header))
}

func (c client) PostWithHeader(url string, b []byte, header *http.Header) (*http.Response, []byte, error) {
	return retry(c.ClientCfg.RetryAttempts, c.ClientCfg.RetryInterval, c.Do("POST", c.ClientCfg.BasePath+url, b, header))
}

func (c client) Do(method, url string, b []byte, header *http.Header) func() (*http.Response, []byte, error) {
	return func() (*http.Response, []byte, error) {
		request, err := http.NewRequest(method, url, bytes.NewBuffer(b))
		if err != nil {
			log.Errorf("[request_err:%s][process:create_request]", err)
			return nil, nil, err
		}

		if header != nil {
			request.Header = *header
		}

		resp, err := c.Client.Do(request)
		if err != nil {
			log.Errorf("[request_err:%s][process:do_request]", err)
			return nil, nil, err
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Errorf("[request_err:%s][process:read_body]", err)
			return nil, nil, err
		}

		if resp.StatusCode >= http.StatusInternalServerError {
			return resp, body, errors.New("Error in status code")
		}

		return resp, body, err
	}
}
