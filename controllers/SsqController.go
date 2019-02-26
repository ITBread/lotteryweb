package controllers

import (
	"fmt"

	"github.com/ITBread/lotteryweb/datamodels"
	"github.com/ITBread/lotteryweb/services"
	"github.com/kataras/iris"
)

// SsqController struct
type SsqController struct {
	Ctx     iris.Context
	Service services.SsqService
}

// Get Get
func (c *SsqController) Get() (results []datamodels.Ssq) {
	c.Ctx.ViewData("message", "Hello world!")
	// 渲染模板文件： ./views/hello.html
	c.Ctx.View("main.html")
	println("===============", c.Ctx.GetCurrentRoute().Name())
	fmt.Printf("===============%v\n", c.Ctx.GetCurrentRoute())
	return nil
}

// GetHistory Get
func (c *SsqController) GetHistory() (results []datamodels.Ssq) {
	c.Ctx.ViewData("message", "Hello world!")
	// 渲染模板文件： ./views/hello.html
	println("=====****======", c.Ctx.GetCurrentRoute().Name())
	println("=====****======", c.Ctx.GetCurrentRoute().Path())
	println("=====****======", c.Ctx.GetCurrentRoute().ResolvePath())
	fmt.Printf("****===========%v\n", c.Ctx.GetCurrentRoute())
	c.Ctx.View("index.html")
	return nil
}

// GetGen GetGen
func (c *SsqController) GetGen() (results []datamodels.Ssq) {
	c.Ctx.View("gen.html")
	return nil
}

// GetBy GetBy
func (c *SsqController) GetBy(id int64) (user datamodels.Ssq, found bool) {
	u, found := c.Service.GetByID(id)
	if !found {
		c.Ctx.Values().Set("message", "User couldn't be found!")
	}
	return u, found
}

//PutBy PutBy
func (c *SsqController) PutBy(id int64) (datamodels.Ssq, error) {
	u := datamodels.Ssq{}
	if err := c.Ctx.ReadForm(&u); err != nil {
		return u, err
	}
	return c.Service.Update(id, u)
}

//DeleteBy DeleteBy
func (c *SsqController) DeleteBy(id int64) interface{} {
	wasDel := c.Service.DeleteByID(id)
	if wasDel {
		return map[string]interface{}{"deleted": id}
	}
	return iris.StatusBadRequest
}
