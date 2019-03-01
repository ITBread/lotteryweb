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

func NewSsqService(engine *xorm.Engine) SsqService {
	return &ssqService{
		engine: engine,
	}
}

type ssqService struct {
	engine *xorm.Engine
}

func (s *ssqService) GetAll() (users []datamodels.User) {
	s.engine.Find(&users)
	return
}
func (s *ssqService) GetAll() (ssq []datamodels.Ssq) {

	return nil
}
