package utils

import (
	"encoding/json"
	"go-firebase-auth/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func StructToJSONString(data interface{}) string {
	jsonBytes, err := json.Marshal(data)
	if err!=nil{
		panic(err.Error())
		return ""
	}
	return string(jsonBytes)
}

func WriteJSON(w http.ResponseWriter, data interface{}){
	w.Header().Set("Content-Type", "application/json")
	parseError := json.NewEncoder(w).Encode(&data)
	if parseError != nil {
		panic(parseError.Error())
	}
}

func ReadJSON(r *http.Request) (models.User, error){
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	return user, err
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func GetParamFromRequestUrl(r *http.Request, param string) string{
	return httprouter.ParamsFromContext(r.Context()).ByName(param)
}
