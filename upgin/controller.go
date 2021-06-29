/**
 * @Author: yinjinlin
 * @File:  controller.go
 * @Description:
 * @Date: 2021/6/29 下午1:29
 */

package upgin

import "github.com/gin-gonic/gin"

var NotAutoRouter = map[string]bool{
	"Init":    true,
	"Prepare": true,
	"Finish":  true,
}

type Controller struct {
	Ctx    *gin.Context
	Method string
}

func (c *Controller) Init(ctx *gin.Context, method string) {
	c.Ctx = ctx
	c.Method = method
}

func (c *Controller) Prepare(ctx *gin.Context) {

}

func (c *Controller) Finish(ctx *gin.Context) {
}

type ControllerInterface interface {
	Init(ctx *gin.Context, method string)
	Prepare(ctx *gin.Context)
	// Get()
	// Post()
	// Delete()
	// Put()
	// Head()
	// Patch()
	// Options()
	// Trace()
	Finish(ctx *gin.Context)
	// Render() error
	// XSRFToken() string
	// CheckXSRFCookie() bool
	// HandlerFunc(fn string) bool
	// URLMapping()
}
