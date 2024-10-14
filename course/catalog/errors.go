package catalog

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrNotEnoughSeats           = errors.New("no seat available")
	ErrClassSoldOut             = ErrInvalidStateChange{Message: "class is sold out"}
	ErrClassNotAvailableForSale = errors.New("class is not available for sale")
)

type ErrInvalidStateChange struct {
	Message string
}

func (e ErrInvalidStateChange) Error() string {
	return e.Message
}

func (e ErrInvalidStateChange) GRPCStatus() *status.Status {
	return status.New(codes.FailedPrecondition, e.Error())
}
