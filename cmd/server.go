package main

import (
	"LunarAssignment/LunarEngine/dao"
	"LunarAssignment/LunarEngine/routes"
	"LunarAssignment/LunarEngine/service"
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
