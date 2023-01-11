package errors

import "errors"

var (
	Unknown       = errors.New("Unknown error")
	NotFound      = errors.New("Not found")
	InvalidParams = errors.New("Invalid params")
	DbNoAffected  = errors.New("0 row affected")
)

var (
	suStr                    = " Service unavailable"
	BonusServiceUnavailable  = errors.New("Bonus" + suStr)
	FlightServiceUnavailable = errors.New("Flight" + suStr)
	TicketServiceUnavailable = errors.New("Ticket" + suStr)
)

type ErrorResponse struct {
	Msg string `json:"message,omitempty"`
}

func ServiceUnavailable(service string) error {
	return errors.New(service + suStr)
}

func ToErrResponse(err error) ErrorResponse {
	return ErrorResponse{Msg: err.Error()}
}
