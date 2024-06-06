package iexample

type ServerIPListRequest struct{}

type ServerIPListResponse struct {
	ResponseResult
	Data struct {
		Leader     string
		Candidates []string
	}
}
