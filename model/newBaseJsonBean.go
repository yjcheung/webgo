package model

type BaseJsonBean struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func NewBaseJsonBean() *BaseJsonBean {

	return &BaseJsonBean{}
}
