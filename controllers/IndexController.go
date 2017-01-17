package controllers

import (
	"fmt"
	"framework/vendor"
)

// IndexController struct
type IndexController struct {
	vendor.Controller
}

// Get func
func (c *IndexController) Get() {
	fmt.Println("hello")
}
