package requester

type IRequesterParams struct {
	Url    string
	Method string
	Body   []byte
	Header []map[string]string
}

type IResponseBody struct {
	HttpCode     int
	ResponseBody []byte
}
