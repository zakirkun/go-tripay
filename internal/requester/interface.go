package requester

type IRequester interface {
	DO() (*IResponseBody, error)
}
