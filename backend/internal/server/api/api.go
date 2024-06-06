package api

import (
	"backend/internal/server/db"
	"backend/utils/types"
	"backend/utils/util"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiServer struct {
	Address string
	DB      *sql.DB
}

// Devuelve objeto que va a representar al servidor con su puerto y DB
func NewApiServer(addr string) *ApiServer {
	return &ApiServer{
		Address: addr,
	}
}

func (apiserver *ApiServer) Run(db *sql.DB) error {

	apiserver.DB = db
	router := http.NewServeMux()
	router.HandleFunc("GET /record", apiserver.HandlerRecord)
	server := http.Server{
		Addr:    apiserver.Address,
		Handler: router,
	}
	fmt.Printf("Server has started %s", apiserver.Address)

	return server.ListenAndServe()
}

// Funcion que devuelve respuesta en la ruta /record
func (apiserver *ApiServer) HandlerRecord(writer http.ResponseWriter, request *http.Request) {
	writer = utils.SetCORSHeaders(writer, request)
	list,err := utils.GetParamsRequest(request)
	if err != nil {
		response := types.Response{
			Status: http.StatusBadRequest,
			Msg:    err.Error(),
			Data:   nil,
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(writer, "Error serializing JSON", http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(jsonResponse))
		return
	}
	data, _ := db.ExtractRecord(apiserver.DB, list)
	msg := utils.ResponseMessage(len(data))

	response := types.Response{
		Status: http.StatusOK,
		Msg:    msg,
		Data:   data,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(writer, "Error serializing JSON", http.StatusInternalServerError)
		return
	}
	writer.Write(jsonResponse)
}
