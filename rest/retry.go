package rest

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/gommon/log"
)

func retry(attempts int, sleep time.Duration, f func() (*http.Response, []byte, error)) (*http.Response, []byte, error) {
	res, body, err := f()
	if err != nil {
		if attempts--; attempts > 0 {
			log.Infof("[retry_attemp:%d]", attempts)
			// Add some randomness to prevent creating a Thundering Herd
			jitter := time.Duration(rand.Int63n(int64(sleep)))
			sleep = sleep + jitter/2

			time.Sleep(sleep)
			return retry(attempts, 2*sleep, f)
		}
		return res, body, err
	}

	return res, body, nil
}
