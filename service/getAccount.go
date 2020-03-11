package service

import (
	"LunarAssignment/LunarEngine/model"
)

func (s *LogicService) GetAccount(accountID int) (*model.Account, error) {

	result, err := s.dao.GetAccount(accountID)
	if err != nil {
		return nil, err
	}

	return result, nil
}
