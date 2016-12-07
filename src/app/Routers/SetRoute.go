package Routers

import (
	//"fmt"
	"net/http"
	"app/Utils"
	"app/Models"
	"golang.org/x/net/context"
	//"gopkg.in/olivere/elastic.v5"
)

func SetHandler(w http.ResponseWriter, r *http.Request)  {
	dat :=  Utils.BodyToJson(r)
	eType := dat["type"].(string)
	bodyData := dat["source"]
	id := dat["id"].(string)
	parent_id := dat["parent_id"].(string)
	operation := dat["operation"].(string)
	client :=  Models.GetElasticCon(Utils.ElasticUrl())
	indexService := client.Index().Index(Index)
	updateSevice := client.Update().Index(Index)
	deleteService := client.Delete().Index(Index)
	var err *error
	if operation == "add"{
		if parent_id != "" {
			indexService =  indexService.Parent(parent_id)
		}
		_,_ = indexService.Id(id).Type(eType).BodyJson(bodyData).Do(context.Background())
	}else if operation == "update"{

		if parent_id != "" {
			updateSevice = updateSevice.Parent(parent_id)

		}

		_,_ = updateSevice.Type(eType).Id(id).Doc(bodyData).DetectNoop(true).Do(context.TODO())
	}else if(operation=="delete"){
		if parent_id != ""{
			deleteService = deleteService.Id(id)
		}
		_,_ = deleteService.Type(eType).Do(context.TODO())
	}
	if err != nil {
		panic(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status": "ok"}`))
}
