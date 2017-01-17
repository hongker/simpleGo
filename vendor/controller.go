package vendor

import "net/http"

// ControllerInterface interface
type ControllerInterface interface {
	Init(ct *Context, controllerName string)
	Get()
}

// Context struct
type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Params         map[int]string
}

// Controller struct
type Controller struct {
	ct   *Context
	name string
}

// Init func
func (c *Controller) Init(ct *Context, controllerName string) {
	c.ct = ct
	c.name = controllerName
}

// GetControllerName func
func (c *Controller) GetControllerName() string {
	return c.name
}

// GetContext func
func (c *Controller) GetContext() *Context {
	return c.ct
}
