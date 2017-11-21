package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/JincorTech/backend-cert2addr/services"
	"github.com/gorilla/mux"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]interface{}{"status": code, "error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *Application) convertCertPemToEthereumLikeAddress(w http.ResponseWriter, r *http.Request) {
	certAddrRequest := CertificateAddressRequest{}
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&certAddrRequest); err != nil {
		fmt.Printf("%v\n", err)
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	cid, err := services.GetCIDFromPem([]byte(certAddrRequest.Pem))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	address, err := services.GetEtheriumLikeAddressFromCid([]byte(cid))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, &ServerResponse{
		Status: http.StatusOK,
		Data: &CertificateAddressResponse{
			Address: address,
		},
	})
}

type Application struct {
	Router           *mux.Router
	httpAuthUsername string
	httpAuthPassword string
}

func (a *Application) basicAuth(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(a.httpAuthUsername) > 0 && len(a.httpAuthPassword) > 0 {
			user, pass, _ := r.BasicAuth()
			if user != a.httpAuthUsername || pass != a.httpAuthPassword {
				respondWithError(w, http.StatusUnauthorized, "Unauthorized")
				return
			}
		}
		fn(w, r)
	}
}

func (a *Application) Initialize(user, password string) {
	a.httpAuthUsername = user
	a.httpAuthPassword = password

	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/api/certificates/actions/getaddress", a.basicAuth(a.convertCertPemToEthereumLikeAddress)).Methods("POST")
}

func (a *Application) Run(addr string) {
	log.Printf("Start HTTP service at %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
