package routes

import (
	"LunarAssignment/LunarEngine/model"
	"encoding/json"
	"net/http"
)

func (rs *RestServer) Update(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Authorization")

	if token != "Bearer SecretInternalKey" {
		w.WriteHeader(401)
		w.Write([]byte("Unauthorized"))
		return
	} else {

		account := &model.Account{}
		json.NewDecoder(r.Body).Decode(account)

		increasedVal, err := rs.service.Update(account)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Internal Error"))
			return
		}

		body, err := json.Marshal(increasedVal)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Internal Error"))
			return
		}

		w.WriteHeader(201)
		w.Write(body)
	}

}
