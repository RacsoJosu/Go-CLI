package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)



func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "El metodo no se puede llamar")
		return
	}
	fmt.Fprintf(w, "metodo: %s", r.Method)
	fmt.Fprintf(w, "Hola %s", "visitante")
}

func getCountries(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(countries)

}
func addCountry(w http.ResponseWriter, r *http.Request){

}