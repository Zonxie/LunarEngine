package service

import (
	"LunarAssignment/LunarEngine/model"
	"log"
)

func (s *LogicService) Update(cr *model.Account) (string, error) {

	result, err := s.dao.UpdateAccount(cr)
	if err != nil {
		log.Fatal(err)
	}

	return result, nil
}
