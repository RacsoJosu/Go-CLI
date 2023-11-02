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
	// especificar el contenido
	w.Header().Set("Content-Type", "application/json")
	// 
	json.NewEncoder(w).Encode(countries)

}
func addCountry(w http.ResponseWriter, r *http.Request){
	country := &Country{}
	err := json.NewDecoder(r.Body).Decode(country)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}
	fmt.Println(country)
	countries = append(countries, country)
	fmt.Fprintf(w, "country ha sido agregado")

}

func updateCountry(w http.ResponseWriter, r *http.Request){
	country := &Country{}
	err := json.NewDecoder(r.Body).Decode(country)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}

	for _, country_ := range countries {

		if country_.ID == country.ID {
			if country.Name != ""  {
				country_.Name = country.Name
				
			}

			if country.Language != ""  {
				country_.Language = country.Language
				
			}


			
		}

		
	}
	fmt.Fprintf(w, "country ha sido actualizado")

}