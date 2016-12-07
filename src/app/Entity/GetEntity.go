package Entity

type GetEntity struct {
	eType string
	query_type string
	child_type string
	start_index int
	array_of_json []interface{}
	size int
}

func NewGetEntity() *GetEntity{
	return &GetEntity{}
}
