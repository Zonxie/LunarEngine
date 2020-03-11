package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (rs *RestServer) GetAccount(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Authorization")

	if token != "Bearer SecretInternalKey" {
		w.WriteHeader(401)
		w.Write([]byte("Unauthorized"))
	} else {

		vars := mux.Vars(r)
		id, found := vars["id"]
		if !found {
			w.WriteHeader(500)
			w.Write([]byte("Bad Request"))
		}

		idInt, err := strconv.Atoi(id)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(500)
			w.Write([]byte("Internal Error"))
			return
		}

		result, err := rs.service.GetAccount(idInt)
		if err != nil {
			if result == nil {
				w.WriteHeader(400)
				w.Write([]byte("Account does not exist"))
				return
			}
			log.Fatal(err)
			w.WriteHeader(400)
			w.Write([]byte("Internal Error"))
			return
		} else {

			body, err := json.Marshal(result)
			if err != nil {
				log.Fatal(err)
				w.WriteHeader(500)
				w.Write([]byte("Internal Error"))
				return
			}
			w.WriteHeader(200)
			w.Write(body)
		}
	}

}
