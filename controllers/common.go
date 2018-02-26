package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

type Response struct {
	Context interface{} `json:"context"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Status  int         `json:"status"`
}

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
	c.Response().WriteHeader(http.StatusBadRequest)
	resp := Response{
		Status:  http.StatusBadRequest,
		Code:    http.StatusBadRequest,
		Message: message,
		Data:    data,
		Context: &ctx,
	}
	return responseJSON(c, resp)
}

// 返回 200 成功
func Success(c echo.Context, message string, data interface{}) error {
	setHeader(c)
	c.Response().WriteHeader(http.StatusOK)
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
	c.Response().WriteHeader(http.StatusNotFound)
	resp := Response{
		Status:  http.StatusNotFound,
		Code:    http.StatusNotFound,
		Message: "方法未找到",
	}
	return responseJSON(c, resp)
}
