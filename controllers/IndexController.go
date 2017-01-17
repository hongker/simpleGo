package controllers

import (
	"fmt"
	"framework/vendor"
)

// IndexController struct
type IndexController struct {
	vendor.Controller
}

// Index func
func (c *IndexController) Index() {
	fmt.Fprintln(c.GetResponseWriter(), "Hello,welcome to index page")
}
