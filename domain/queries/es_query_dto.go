package queries

type EsQuery struct {
	Equals []FieldValue `json:"equals"`
}

type FieldValue struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}

/*
to make common interface interactiong with different elasticsearch providers
and for use in search constuctions like this

{
	"equals": [
		{
			"field": "status",
			"value": "pending"
		},
		{
			"field": "seller",
			"value": "1"
		},
	],
	"any_equals": [
		{
			"field": "status",
			"value": "pending"
		},
		{
			"field": "seller",
			"value": "1"
		},
	]
}
*/
