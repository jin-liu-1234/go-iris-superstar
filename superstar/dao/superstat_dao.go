package dao

import (
	"github.com/go-xorm/xorm"
	"_iris/superstar/models"
	"log"
)

type SuperstarDao struct {
	engine *xorm.Engine
}

func NewSuperstarDao(engine *xorm.Engine) (*SuperstarDao) {
	return &SuperstarDao{
		engine: engine,
	}
}

func (d *SuperstarDao) Get(id int) *models.StarInfo {
	data := &models.StarInfo{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	}
	return nil
}

func (d *SuperstarDao) GetAll(page, per_page int) ([]models.StarInfo, int64) {
	list := []models.StarInfo{}

	var count int64
	count, e := d.engine.Count(models.StarInfo{})
	if e != nil {
		return list, 0
	}

	err := d.engine.Desc("id").Limit(per_page, per_page*(page-1)).Find(&list)
	if err != nil {
		return list, 0
	}

	return list, count
}

func (d *SuperstarDao) Search(country string, page, per_page int) ([]models.StarInfo, int64) {
	list := []models.StarInfo{}
	var count int64
	err := d.engine.Where("country like ?", country+"%").Or("name_zh like ?", country+"%").Desc("id").Limit(per_page, per_page*(page-1)).Find(&list)
	count, e := d.engine.Where("country like ?", country+"%").Or("name_zh like ?", country+"%").Count(models.StarInfo{})
	if err != nil || e != nil {
		log.Println(err)
		log.Println(e)
		return list, 0
	}

	return list, count
}

func (d *SuperstarDao) Create(data *models.StarInfo) error {
	_, err := d.engine.Insert(data)
	return err
}

func (d *SuperstarDao) Update(data *models.StarInfo, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)

	return err
}

func (d *SuperstarDao) Delete(id int) error {
	data := &models.StarInfo{Id: id, Status: 1}
	_, err := d.engine.Id(id).Update(data)

	return err
}
