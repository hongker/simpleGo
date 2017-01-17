package controllers

import (
	"fmt"
	"framework/vendor"
)

// UserController struct
type UserController struct {
	vendor.Controller
}

// Index function
func (c *UserController) Index() {
	fmt.Fprintln(c.GetResponseWriter(), "Hello,welcome to user module")
	return

}

// Hello function
func (c *UserController) Hello() {
	fmt.Fprintln(c.GetResponseWriter(), "Hello,welcome to user hello function")
	return

}

// Info function
func (c *UserController) Info() {
	params := c.GetParams()
	uid := params[1]
	fmt.Fprintln(c.GetResponseWriter(), "Hello,"+uid)
	return
}
