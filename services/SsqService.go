package services

import (
	"github.com/ITBread/lotteryweb/datamodels"
	"github.com/go-xorm/xorm"
)

type SsqService interface {
	GetAll() []datamodels.Ssq
	GetByID(id int64) (datamodels.Ssq, bool)
	DeleteByID(id int64) bool
	Update(id int64, user datamodels.Ssq) (datamodels.Ssq, error)
	Create(red string, ssq datamodels.Ssq) (datamodels.Ssq, error)
}

// func NewSsqService(engine *xorm.Engine) SsqService {
// 	return &MySsqService{
// 		engine: engine,
// 	}
// }

type MySsqService struct {
	engine *xorm.Engine
}

func (s *MySsqService) GetAll() (users []datamodels.Ssq) {
	s.engine.Find(&users)
	return
}
