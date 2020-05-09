package controllers

type Context interface {
	Param(string) string
	FormValue(string) string
	Bind(interface{}) error
	JSON(int, interface{}) error
}
