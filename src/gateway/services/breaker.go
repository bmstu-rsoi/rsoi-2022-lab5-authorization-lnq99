package services

import (
	"context"
	"io"
	"net/http"
	"time"

	errors "github.com/lnq99/rsoi-2022-lab3-fault-tolerance-lnq99/src/pkg/error"
	breaker "github.com/sony/gobreaker"
)

var (
	CbSetting = breaker.Settings{
		Name:        "Ticket Get",
		MaxRequests: 1,
		Interval:    10 * time.Second,
		Timeout:     2 * time.Second,
		ReadyToTrip: func(counts breaker.Counts) bool {
			return counts.ConsecutiveFailures > 4
		},
		OnStateChange: nil,
		IsSuccessful:  nil,
	}

	bonusCb  = breaker.NewCircuitBreaker(CbSetting)
	ticketCb = breaker.NewCircuitBreaker(CbSetting)
	flightCb = breaker.NewCircuitBreaker(CbSetting)
)

type ServiceResponse struct {
	status int
	body   io.ReadCloser
}

func CallServiceWithCircuitBreaker(
	cb *breaker.CircuitBreaker, method, url string,
	header map[string]string, body io.Reader) (ServiceResponse, error) {

	res, err := cb.Execute(func() (interface{}, error) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		req, _ := http.NewRequestWithContext(ctx, method, url, body)

		for k, v := range header {
			req.Header.Set(k, v)
		}
		res, err := Client.Do(req)
		if err != nil {
			return ServiceResponse{status: http.StatusServiceUnavailable},
				errors.BonusServiceUnavailable
		}
		return ServiceResponse{
			status: res.StatusCode,
			body:   res.Body,
		}, nil
	})

	if err == breaker.ErrOpenState {
		return ServiceResponse{}, err
	}

	return res.(ServiceResponse), err
}
