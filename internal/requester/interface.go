package requester

import "context"

type IRequester interface {
	DO() (*IResponseBody, error)
	DOWithContext(ctx context.Context) (*IResponseBody, error)
}
