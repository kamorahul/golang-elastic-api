# golang-elastic-api
A working api with golang and oliver elastic search
	Search	/set POST	"{
		""api_key"" : ""String"",
		""entity"" : ""String"",
		""operation"" : ""add"" || ""update || delete"",
		""id"" : ""String"",
		""parent_id"" : ""String"", <Optional>
		""source"" : <json> 

	}
	Response 	{status : true}



	Search 	/get POST	"{
		""query_type"" : ""filter || parent"",
		""request_data"" : ""JSON"",
		""type"" : ""conversation"",
		""child_type"" : ""message"",
		""start_index"" : ""1"",
		""size"" : ""2"",
		""sort"" : {
			""field"" : ""conersation_id"",
			""asc || dec"" : true,

		}
		""query_json"" :[{
			""key"" : ""message_id"", 
			""match"" : ""text || keyword || range"",
			""query_type"" : ""must || should || must_not || filter""
			""value"" : ""12123jb1n2bn123b12nb3n123""}
			]
		}        
	}
	Response 	"{
	    status: ""boolean"",
	    total_records : <Number>,
	    request_data : <JSON>.
	    data_source : [
	    ]    

	}"
	Search 	/map POST	"{
		""type"" : ""String"",
		""mapping_json"" : ""JSON""
	}"	

	Return {<Mapping Json For that type>}
