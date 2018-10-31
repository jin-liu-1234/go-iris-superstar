package services

import (
	"_iris/superstar/models"
	"_iris/superstar/dao"
	"_iris/superstar/datasource"
)

type SuperstarService interface {
	GetAll(page, per_page int) ([]models.StarInfo, int64)
	Get(id int) *models.StarInfo
	Delete(id int) error
	Update(user *models.StarInfo, columns []string) error
	Create(user *models.StarInfo) error
	Search(country string, page, per_page int) ([]models.StarInfo, int64)
}

type superstarService struct {
	dao *dao.SuperstarDao
}

func NewSuperstarService() SuperstarService {
	return &superstarService{
		dao: dao.NewSuperstarDao(datasource.InstanceMaster()),
	}
}

func (s *superstarService) GetAll(page, per_page int) ([]models.StarInfo, int64) {
	return s.dao.GetAll(page, per_page)
}

func (s *superstarService) Get(id int) *models.StarInfo {
	return s.dao.Get(id)
}

func (s *superstarService) Delete(id int) error {
	return s.dao.Delete(id)
}

func (s *superstarService) Update(user *models.StarInfo, columns []string) error {
	return s.dao.Update(user, columns)
}

func (s *superstarService) Create(user *models.StarInfo) error {
	return s.dao.Create(user)
}

func (s *superstarService) Search(country string, page, per_page int) ([]models.StarInfo, int64) {
	return s.dao.Search(country, page, per_page)
}
