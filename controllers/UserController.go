package controllers

import (
	"fmt"
	"framework/vendor"
)

// UserController struct
type UserController struct {
	vendor.Controller
}

// Get func
func (c *UserController) Get() {
	fmt.Println("hello")
}

// Hello function
func (c *UserController) Hello() {
	fmt.Println(c.GetControllerName())
	return

}

// Info function
func (c *UserController) Info() {
	ct := c.GetContext()
	uid := ct.Params[1]
	fmt.Fprintln(ct.ResponseWriter, "Hello,"+uid)
	return
}
