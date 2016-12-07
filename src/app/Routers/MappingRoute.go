package Routers

import (
	"net/http"
	"fmt"
	"app/Models"
	"app/Utils"
	//"gopkg.in/olivere/elastic.v5"
	//"gopkg.in/olivere/elastic.v5"
	//"golang.org/x/crypto/ssh/terminal"
	"context"
	//"github.com/kataras/go-serializer/json"
	"encoding/json"
)

func MappingHandler(w http.ResponseWriter, r *http.Request)  {
	client :=  Models.GetElasticCon(Utils.ElasticUrl())
	data := Utils.BodyToJson(r);
	eType := data["entity"].(string)
	mappingData := data["mapping_json"].(map[string]interface{})
	//mappingIndex := elastic.NewAliasAddAction()
	putMappingResponse,err :=  client.PutMapping().Index(Utils.DefaultIndex()).Type(eType).BodyJson(mappingData).Do(context.TODO())
	if err != nil{
		panic(err);
	}
	fmt.Println(putMappingResponse)
	newMapping,err := client.GetMapping().Index(Utils.DefaultIndex()).Type(eType).Do(context.TODO())
	//var b map[string]interface{}
	b,marshalErr :=  json.Marshal(newMapping)
	if marshalErr != nil{
		panic(marshalErr)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	//w.Write([]byte("Gorilla Map!\n"))
	//fmt.Println(client.ClusterState())
	fmt.Fprint(w,string(b))

}

