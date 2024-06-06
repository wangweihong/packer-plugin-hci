package iexample

import "fmt"

type ResponseResult struct {
	ErrorCode   int32    `json:"scode"`
	Message     string   `json:"message_cn"`
	Stack       []string `json:"stack"`
	Description string   `json:"desc"`
}

func (rr ResponseResult) Error() error {
	if rr.ErrorCode == 0 {
		return nil
	}

	return fmt.Errorf("scode:%d, message:%v,stack:%v,desc:%v", rr.ErrorCode, rr.Message, rr.Stack, rr.Description)
}
