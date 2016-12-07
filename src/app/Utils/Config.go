package Utils

func ElasticUrl() string{

	str := GetRecordSrv("elastic.service.consul")
	if str == ""{
		return "http://127.0.0.1:9200"
	}
	return str;


}



func DefaultIndex() string{
	return "search"
}
