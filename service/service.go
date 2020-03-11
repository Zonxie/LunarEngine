package service

import "LunarAssignment/LunarEngine/dao"

type LogicService struct {
	dao *dao.Dao
}

func NewLogicService(dao *dao.Dao) *LogicService {
	return &LogicService{dao: dao}
}
