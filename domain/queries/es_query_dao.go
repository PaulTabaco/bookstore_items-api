package queries

import "fmt"

func (q EsQuery) ToFinalSearchQuery() map[string]interface{} {
	// Parsing EsQuery equls to es match array
	matches := make([]map[string]interface{}, 0)
	for _, eq := range q.Equals {
		match := map[string]interface{}{
			"match": map[string]interface{}{
				eq.Field: eq.Value,
			},
		}
		matches = append(matches, match)
	}
	// set array of matches to final es search query
	return map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": matches,
			},
		},
	}
}

func (q EsQuery) ToFinalUpdateQuery() map[string]interface{} {

	field1 := q.Equals[0].Field
	value1 := q.Equals[0].Value

	// source1 := fmt.Sprintf("ctx._source.%s = params.%s", field1, field1)
	source1 := fmt.Sprintf("ctx._source.%s = params.%s", field1, field1)

	// { "scripted_upsert":true, "script" :{ "source": "if(ctx._version == null) { ctx._source = params; } else { def param = params; def src = ctx._source; for(s in src.documents) { boolean found = false; for(p in param.documents) { if (p.docid == s.docid) { found = true; if(s.added_vals != null) { p.added_vals = s.added_vals; } } } if(!found) param.documents.add(s); } ctx._source = param; }", "lang": "painless", "params" : { "documents": [ { "docid": "ID001", "status" : "cancelled" } ], "id": "1" } }, "upsert" : {  } }

	query := map[string]interface{}{
		"script": map[string]interface{}{
			"source": source1,
			// "source": "if (ctx._source.tags.contains(params.tag)) { ctx.op = 'delete' } else { ctx.op = 'noop' }",
			// "source": "if (ctx._source.tags.contains(params.tag)) { ctx._source.tags.remove(ctx._source.tags.indexOf(params.tag)) }"
			/*
					"source": """
					if ( ctx.op == 'create' ) {
					  ctx._source.counter = params.count
					} else {
					  ctx._source.counter += params.count
					}
				  	""",
			*/
			"lang": "painless",
			"params": map[string]interface{}{
				field1: value1,
			},
		},
	}

	return query
}
