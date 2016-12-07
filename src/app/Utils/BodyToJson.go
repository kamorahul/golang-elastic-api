package Utils

import (
	"net/http"
	"fmt"
	"reflect"
	"encoding/json"
)

func BodyToJson(r *http.Request)  map[string]interface{}{
	decoder := json.NewDecoder(r.Body);
	fmt.Println(reflect.TypeOf(r.Body).Kind())
	var dat map[string]interface{}
	err := decoder.Decode(&dat)
	if err!= nil{
		panic(err);
	}
	return dat
}
