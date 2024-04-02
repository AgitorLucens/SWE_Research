package main

import (
	"log"
	"net/http"
	"encoding/json"
)

type APIServer struct {
	addr string
}

//Response
type Response struct {
	Status int `json:"status"`
	Data *Record	`json:"data"`

}

type ResponseData struct {
	Status int `json:"status"`
	Data any	`json:"data"`

}
func NewAPIServer(addr string) *APIServer{
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	DB := newAPIDB("user=postgres password=dontwell12 dbname=swe-research-db host=localhost port=5432 sslmode=disable")
	err := DB.initDB()

	//end  points
	router.HandleFunc("GET /record", func(w http.
		ResponseWriter, r *http.Request){

			data, _ := extractRecord(DB.DB)

			if err != nil {
				http.Error(w, "Error al obtener registros", http.StatusInternalServerError)
				log.Fatal(err)
			}

			
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			//log.Printf("%s", data.Descr)
			json.NewEncoder(w).Encode(&Response{ http.StatusOK, &Record{
				Name: data.Name,
				Price: data.Price,
				Descr: data.Descr,
			}})

			
			if err != nil {
            	http.Error(w, "Error al codificar respuesta JSON", http.StatusInternalServerError)
            return
        	}
			
		})//end endpoint
	
	server := http.Server{
		Addr: s.addr,
		Handler: router,
	}

	log.Printf("Server has started %s", s.addr)

	return server.ListenAndServe()
}
