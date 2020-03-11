package main

import (
	"LunarAssignment/engine/dao"
	"LunarAssignment/engine/routes"
	"LunarAssignment/engine/service"
	"fmt"
)

func main() {

	dao, err := dao.NewDao()
	if err != nil {
		fmt.Println(err)
	}

	s := service.NewLogicService(dao)
	routes.StartRestServer(s, dao)

}
