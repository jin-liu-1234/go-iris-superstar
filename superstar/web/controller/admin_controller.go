package controller

import (
	"github.com/kataras/iris"
	"_iris/superstar/services"
	"github.com/kataras/iris/mvc"
	"_iris/superstar/models"
	"log"
	"time"
)

type AdminController struct {
	Ctx     iris.Context
	Service services.SuperstarService
}

func (c *AdminController) Get() mvc.Result {
	list, _ := c.Service.GetAll(1, 10)
	return mvc.View{
		Name: "index.html",
		Data: iris.Map{
			"Title": "球星数据库-管理后台",
			"List":  list,
		},
		Layout: "layout.html",
	}
}

func (c *AdminController) GetEdit() mvc.Result {
	id, err := c.Ctx.URLParamInt("id")
	var data *models.StarInfo
	if err == nil {
		data = c.Service.Get(id)
	}

	return mvc.View{
		Name: "edit.html",
		Data: iris.Map{
			"Title": "球星数据库-管理后台",
			"Info":  data,
		},
		Layout: "layout.html",
	}
}

func (c *AdminController) PostSave() mvc.Result {
	info := models.StarInfo{}
	err := c.Ctx.ReadForm(&info)

	if err != nil {
		log.Println(err)
	}

	if info.Id > 0 {
		info.Updated = int(time.Now().Unix())
		c.Service.Update(&info, []string{
			"name_zh", "name_eh", "avatar", "birthday", "height", "weight", "club",
			"jersy", "country", "birth_address", "feature", "more_info",
		})
	} else {
		// 创建
		info.Created = int(time.Now().Unix())
		info.Updated = int(time.Now().Unix())
		c.Service.Create(&info)
	}

	return mvc.Response{
		Path: "/admin/",
	}
}

func (c *AdminController) Delete() mvc.Result {
	id, err := c.Ctx.URLParamInt("id")
	if err == nil {
		c.Service.Delete(id)
	}

	return mvc.Response{
		Path: "/admin/",
	}
}
