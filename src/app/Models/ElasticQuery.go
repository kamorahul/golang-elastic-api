package Models

import (

	"gopkg.in/olivere/elastic.v5"
	"app/Utils"
	"golang.org/x/net/context"
)
type GetEntity struct {
	eType string
	query_type string
	child_type string
	start_index int
	array_of_json []interface{}
	size int
}



func  SearchParentByChild(getEntity GetEntity)  *elastic.SearchResult{
	client := GetElasticCon(Utils.ElasticUrl())
	bq := elastic.NewBoolQuery()
	datRecord := getEntity.array_of_json[0]
	res := datRecord.(map[string]interface{})
	key := res["key"].(string)
	value := res["value"].(string)

	matchChildQuery := elastic.NewHasChildQuery(getEntity.child_type, elastic.NewMatchQuery(key , value)).
	InnerHit(elastic.NewInnerHit().Name("messages"))
	bq = bq.Must(elastic.NewMatchAllQuery())
	bq = bq.Filter(matchChildQuery)
	searchResult, err := client.Search().
		Index(Utils.DefaultIndex()).
		Type(getEntity.eType).
		Query(bq).From(getEntity.start_index).Size(getEntity.size).
		Pretty(true).
		Do(context.TODO())
	if err != nil {
		panic(err);
	}
	return searchResult
}