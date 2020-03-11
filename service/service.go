package service

import "LunarAssignment/engine/dao"

type LogicService struct {
	dao *dao.Dao
}

func NewLogicService(dao *dao.Dao) *LogicService {
	return &LogicService{dao: dao}
}
