package routes

import (
	"LunarAssignment/LunarEngine/dao"
	"LunarAssignment/LunarEngine/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type RestServer struct {
	service *service.LogicService
	dao     *dao.Dao
}

func StartRestServer(service *service.LogicService, dao *dao.Dao) {
	rs := &RestServer{service: service, dao: dao}

	router := mux.NewRouter()

	router.HandleFunc("/update", rs.Update).Methods("POST")
	router.HandleFunc("/account/{id}", rs.GetAccount).Methods("GET")
	router.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	})

	fmt.Println("Logic: Listning on port 6000")
	log.Fatal(http.ListenAndServe(":6000", router))

}
