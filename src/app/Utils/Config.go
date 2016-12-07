package Utils

func ElasticUrl() string{

	str := GetRecordSrv("elastic.service.consul")
	if str == ""{
		return "http://192.168.86.103:9200"
	}
	return str;


}



func DefaultIndex() string{
	return "search"
}
