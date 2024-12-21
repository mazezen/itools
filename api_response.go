package itools

import "time"

var Response *result

type result struct {
	TraceId string      `json:"trace_id"`
	Lasting string      `json:"lasting"`
	Status  bool        `json:"status"`
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func (r *result) ResponseSuccess(traceId, message string, data ...interface{}) *result {
	var d interface{}
	if data != nil {
		d = data[0]
	} else {
		d = data
	}

	if message == "" {
		message = "success"
	}
	return &result{
		TraceId: traceId,
		Lasting: time.Now().Format(time.DateTime),
		Status:  true,
		Code:    2000,
		Msg:     message,
		Data:    d,
	}
}

func (r *result) ResponseError(traceId string, code int, message string, data ...interface{}) *result {
	var d interface{}
	if data != nil {
		d = data[0]
	} else {
		d = data
	}

	return &result{
		TraceId: traceId,
		Lasting: time.Now().Format(time.DateTime),
		Status:  false,
		Code:    code,
		Msg:     message,
		Data:    d,
	}
}

var ResponsePage *pageList

type pageList struct {
	Total int         `json:"total"`
	List  interface{} `json:"list"`
}

func (p *pageList) Pagination(count int, list interface{}) *pageList {
	return &pageList{
		Total: count,
		List:  list,
	}
}
