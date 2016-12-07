package Routers

import (
	"fmt"
	"app/Models"
	"gopkg.in/olivere/elastic.v5"
	"app/Utils"
	"net/http"
	"golang.org/x/net/context"
	"app/Entity"
	"encoding/json"
)
var Index = "search"

func GetHandler(w http.ResponseWriter, r *http.Request)  {
	client := Models.GetElasticCon(Utils.ElasticUrl());
	dat := Utils.BodyToJson(r)
	eType := dat["type"].(string)
	query_type := dat["query_type"].(string)
	child_type := dat["child_type"].(string)
	start_index := int(dat["start_index"].(float64))
	array_of_json := dat["query_json"].([]interface{});
	size := int(dat["size"].(float64))
	sorting,err1 := dat["sort"].(map[string]interface{});
	if err1 != true {
		panic(err1)
	}
	var  fieldName string
	var  sortType bool
	for i:= range sorting{
		if(i=="field"){
			fieldName =  sorting[i].(string)
		}else if(i=="asc"){
			sortType = true;
		}
	}
	bq := elastic.NewBoolQuery()
	if query_type=="parent" {
		datRecord := array_of_json[0]
		res := datRecord.(map[string]interface{})
		key := res["key"].(string)
		value := res["value"].(string)

		matchChildQuery := elastic.NewHasChildQuery(child_type, elastic.NewMatchQuery(key , value)).
			InnerHit(elastic.NewInnerHit().Name("messages"))
		bq = bq.Must(elastic.NewMatchAllQuery())
		bq = bq.Filter(matchChildQuery)

	}else {
		//newQ := elastic.NewBoolQuery()
		for i := 0; i < len(array_of_json); i++ {
			datRecord := array_of_json[i]
			res := datRecord.(map[string]interface{})
			qType := res["query_type"].(string)
			matchQueryType := res["match"].(string)
			key := res["key"].(string)
			value := res["value"].(interface{})
			//switch vv := value.(type) {
			//case string:
			//
			//case int:
			//
			//case []interface{}:
			//	for i, u := range vv {
			//		fmt.Println(i, u)
			//	}
			//default:
			//	fmt.Println(k, "is of a type I don't know how to handle")
			//}
			var matchType  *elastic.MatchQuery
			var termQuery *elastic.TermQuery
			var rangeQuery *elastic.RangeQuery
			match := 0
			switch matchQueryType {
			case "text" :
				value := res["value"].(string)
				fmt.Println(value)
				matchType = elastic.NewMatchQuery(key,value)
				break;
			case "keyword" :
				match = 1
				termQuery = elastic.NewTermQuery(key,value)
				break;
			case "range" :
				match = 2
				rangeQuery = elastic.NewRangeQuery(key)
				valueRange := value.(map[string]interface{})

				for i := range valueRange{
					switch i {
					case "gte" :
						rangeQuery =  rangeQuery.Gte(valueRange[i])
						break
					case "gt" :
						rangeQuery =  rangeQuery.Gt(valueRange[i])
						break
					case "lte" :
						rangeQuery =  rangeQuery.Lte(valueRange[i])
						break
					case "lt" :
						rangeQuery =  rangeQuery.Lt(valueRange[i])
						break
					}

				}
				break;
			}
			switch qType {
			case "must" :
				if match == 0 {
					bq = bq.Must(matchType)
				}else {
					bq = bq.Must(termQuery)
				}
				break;
			case "filter":
				if match ==0 {

					bq = bq.Filter(matchType)
				}else {
					bq = bq.Filter(termQuery)
				}
				break;
			case "must_not":
				if match ==0 {

					bq = bq.MustNot(matchType)
				}else {
					bq = bq.MustNot(termQuery)

				}
				break;
			case "should":
				if match==0 {

					bq =  bq.Should(matchType)
				}else {
					bq =  bq.Should(termQuery)

				}
				break;
			}
			//newQ = newQ.Should(matchType)
		}
		//bq.Filter(newQ)
	}

	fmt.Println(start_index,size)
	var searchResult *elastic.SearchResult
	eQuery := client.Search().
		Index(Index).
		Type(eType).
		Query(bq).From(start_index).
		Size(size)
	if(fieldName != ""){
		eQuery = eQuery.Sort(fieldName,sortType)
	}
	searchResult,err := eQuery.Pretty(true).Do(context.Background())
	if err!= nil {
		panic(err)
	}
	hits := searchResult.Hits.Hits

	datArray := make([]map[string]interface{},len(hits))
	var dat1 map[string]interface{}

	for i := 0;i < len(hits) ; i++ {
		hit := searchResult.Hits.Hits[i]
		if err := json.Unmarshal(*hit.Source,&dat1); err != nil {
			panic(err)
		}
		fmt.Println(dat1)
		datArray[i] = dat1;
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")


	response := Entity.JsonResponse{"data_source":datArray,"status" :true,"length" :len(hits)}
	b,err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w,string(b))

}
