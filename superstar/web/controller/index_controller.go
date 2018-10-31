package controller

import (
	"github.com/kataras/iris"
	"_iris/superstar/services"
	"github.com/kataras/iris/mvc"
	"math"
)

type IndexController struct {
	Ctx     iris.Context
	Service services.SuperstarService
}

func (c *IndexController) Get() mvc.Result {
	page, _ := c.Ctx.URLParamInt("page")
	if page == -1 {
		page = 1
	}
	per_page, _ := c.Ctx.URLParamInt("per_page")
	if per_page == -1 {
		per_page = 12
	}
	list, count := c.Service.GetAll(page, per_page)

	totalPage := math.Ceil(float64(count) / float64(per_page))
	var viewPage []int
	for i := 1; i <= int(totalPage); i++ {
		viewPage = append(viewPage, i)
	}

	return mvc.View{
		Name: "index.html",
		Data: iris.Map{
			"Title":    "球星数据库",
			"List":     list,
			"Count":    count,
			"viewPage": viewPage,
		},
	}
}

func (c *IndexController) GetBy(id int) mvc.Result {
	if id < 1 {
		return mvc.Response{
			Path: "/",
		}
	}
	data := c.Service.Get(id)
	return mvc.View{
		Name: "info.html",
		Data: iris.Map{
			"Title": "球星数据库",
			"info":  data,
		},
	}
}

func (c *IndexController) GetSearch() mvc.Result {
	key_word := c.Ctx.URLParam("key_word")
	page, _ := c.Ctx.URLParamInt("page")
	if page == -1 {
		page = 1
	}
	per_page, _ := c.Ctx.URLParamInt("per_page")
	if per_page == -1 {
		per_page = 12
	}
	list, count := c.Service.Search(key_word, page, per_page)

	totalPage := math.Ceil(float64(count) / float64(per_page))
	var viewPage []int
	for i := 1; i <= int(totalPage); i++ {
		viewPage = append(viewPage, i)
	}

	return mvc.View{
		Name: "index.html",
		Data: iris.Map{
			"Title":    "球星数据库",
			"List":     list,
			"Count":    count,
			"viewPage": viewPage,
			"key_word": key_word,
		},
	}

	return nil
}
