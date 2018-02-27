package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

type Response struct {
	Context interface{} `json:"context,omitempty"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Status  int         `json:"status"`
}

// 设置响应头 Content-Type 为 application/json
func setHeader(c echo.Context) {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
}

// 返回 JSON 响应结果
func responseJSON(c echo.Context, resp Response) error {
	respBody := Response{
		Code:    resp.Code,
		Message: resp.Message,
		Context: resp.Context,
		Data:    resp.Data,
		Status:  resp.Status,
	}

	c.Set("respBody", respBody)

	return c.JSON(resp.Status, respBody)
}

// 返回 400 错误
func Error(c echo.Context, message string, data interface{}, ctx ...interface{}) error {
	setHeader(c)
	resp := Response{
		Status:  http.StatusBadRequest,
		Code:    http.StatusBadRequest,
		Message: message,
		Data:    data,
		Context: ctx,
	}
	return responseJSON(c, resp)
}

// 返回 200 成功
func Success(c echo.Context, message string, data interface{}) error {
	setHeader(c)
	resp := Response{
		Status:  http.StatusOK,
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	}
	return responseJSON(c, resp)
}

// 返回 404 错误
func NotFound(c echo.Context) error {
	setHeader(c)
	resp := Response{
		Status:  http.StatusNotFound,
		Code:    http.StatusNotFound,
		Message: "方法未找到",
	}
	return responseJSON(c, resp)
}

// 错误处理
func HttpErrorHandler(err error, c echo.Context) {
	status := http.StatusInternalServerError
	code := http.StatusInternalServerError
	message := http.StatusText(code)

	var ctx interface{}

	if he, ok := err.(*echo.HTTPError); ok {
		status = he.Code
		code = he.Code
		message = he.Message.(string)
	}

	if c.Echo().Debug {
		ctx = err
	}

	resp := Response{
		Status:  status,
		Code:    code,
		Message: message,
		Context: ctx,
	}

	if !c.Response().Committed {
		if err := responseJSON(c, resp); err != nil {
			c.Logger().Error(err)
		}
	}
}
