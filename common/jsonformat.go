package common

type JsonFormat struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data map[string]interface{} `json:"data"`
}
