package utils

import (
	"backend/utils/rand"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Devuelve los parametros introducidos en la peticion
func GetParamsRequest(request *http.Request) ( []string, error) {
	paths := []string{"name", "descr", "quant", "price", "topic", "created", "published", "page", "limit"}
	list2 := []string{"", "", "", "", "", "", "", "", ""}
	for i := 0; i < len(paths); i++ {
		filter := request.URL.Query().Get(paths[i])
		if filter != "" && filter != " " {
			err := validateFilter(filter, list2, paths, i)
			if err != nil {
				return list2, err
			}
		} else if filter == "" && paths[i] == "page" {
			list2[i] = "1"
		} else if filter == "" && paths[i] == "limit" {
			list2[i] = "10"
		}
	}
	return  list2, nil
}

func RandomData() (string, float64, string, string, int, string, string) {
	randNumber := rand.RandomInt(0, 59)
	name := rand.GetName(randNumber)
	price := rand.RandomFloat(10.0, 100.0)
	description := rand.GetDescription(randNumber)
	img := rand.GetImage(randNumber)
	quant := rand.RandomInt(1, 1000)
	topic := rand.GetTopic(randNumber)
	created := rand.RandomDate()
	return name, price, description, img, quant, topic, created
}

// Agrega elementos a la cabecera de la respuesta segun el estandar CORS
func SetCORSHeaders(writer http.ResponseWriter, request *http.Request) http.ResponseWriter {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	if request.Method == "OPTIONS" {
		writer.WriteHeader(http.StatusOK)
		return writer
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	return writer
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func IsValidInteger(str string) bool {
	number, err := strconv.Atoi(str)
	if number < 0 {
		return false
	}
	return err == nil
}

func IsValidFloat(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}

func IsValidDate(dateStr string, tag string) bool {
	result := false
	formats := []string{
		"2006-01-02",
		"02-01-2006",
		"2006-01-02T15:04:05Z",
		"2006-01-02T15:04:05.999999Z",
	}
	if tag == "created" {
		for i := 0; i < 3; i++ {
			_, err := time.Parse(formats[i], dateStr)
			if err != nil {
				result = true
			}
		}
	} else {
		_, err := time.Parse(formats[3], dateStr)
		if err != nil {
			result = true
		}
	}
	return result
}

func IsEmpty(str string) bool {
	if str != "" && str != " " {
		return false
	}
	return true
}

func validatePage(list []string, filter string, tag string, iter int) error {
	if filter == "0" && tag == "page" {
		list[iter] = "1"
	} else if !IsValidInteger(filter) && tag == "page" {
		list[iter] = filter
		return errors.New("formato invalido de Pagina")
	} else if tag == "page" {
		list[iter] = filter
	}
	return nil
}

func validateLimit(list []string, filter string, tag string, iter int) error {
	if filter == "0" && tag == "limit" {
		list[iter] = "10"
	} else if !IsValidInteger(filter) && tag == "limit" {
		list[iter] = filter
		return errors.New("formato invalido de Limite")
	} else if tag == "limit" {
		list[iter] = filter
	}
	return nil
}

func validateQuant(list []string, filter string, tag string, iter int) error {
	if tag == "quant" && IsValidInteger(filter) {
		list[iter] = filter
	} else if tag == "quant" && !IsValidInteger(filter) {
		list[iter] = filter
		return errors.New("formato invalido de Cantidad")
	}
	return nil
}

func validateCreated(list []string, filter string, tag string, iter int) error {

	if tag == "created" && IsValidDate(list[iter], "created") {
		list[iter] = filter
	} else if tag == "created" && !IsValidDate(list[iter], "created") {
		list[iter] = filter
		return errors.New("formato invalido de Creacion")
	}
	return nil
}

func validatePublished(list []string, filter string, tag string, iter int) error {
	if tag == "published" && IsValidDate(list[iter], "published") {
		list[iter] = filter
	} else if tag == "published" && !IsValidDate(list[iter], "published") {
		list[iter] = filter
		return errors.New("formato invalido de Creacion")
	}
	return nil
}

// Valida nombre, descripcion y tema que no esten vacios
func validateString(str string, list []string, iter int) {
	if str != "" || str != " " {
		list[iter] = str
	}
}

func validateFilter(filter string, list2 []string, paths []string, iter int) error {
	err := validatePage(list2, filter, paths[iter], iter)
	if err != nil {
		return errors.New("formato invalido de Pagina")
	}
	err = validateLimit(list2, filter, paths[iter], iter)
	if err != nil {
		return errors.New("formato invalido de Limite")
	}
	err = validateQuant(list2, filter, paths[iter], iter)
	if err != nil {
		return errors.New("formato invalido de Cantidad")
	}
	err = validateCreated(list2, filter, paths[iter], iter)
	if err != nil {
		return errors.New("formato invalido de Creacion")
	}
	err = validatePublished(list2, filter, paths[iter], iter)
	if err != nil {
		return errors.New("formato invalido de Creacion")
	}
	validateString(filter, list2, iter)
	return err
}

func ResponseMessage(size int) string {
	if size > 0 {
		return "Records found"
	} else {
		return "Records not found"
	}
}
