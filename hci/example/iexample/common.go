package iexample

type PagingParam struct {
	PageNum  int `form:"page_num" json:"page_num"`
	PageSize int `form:"page_size" json:"page_size"`
}
