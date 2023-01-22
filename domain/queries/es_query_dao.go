package queries

func (q EsQuery) ToCorrectEsQuery() map[string]interface{} {
	matches := make([]map[string]interface{}, 0)
	for _, eq := range q.Equals {
		match := map[string]interface{}{
			"match": map[string]interface{}{
				eq.Field: eq.Value,
			},
		}
		matches = append(matches, match)
	}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": matches,
			},
		},
	}
	return query
}

// func (q EsQuery) Build2() types.Query {
// 	equalsQueries := make([]types.Query, 0)

// 	for _, eq := range q.Equals {

// 		value := eq.Value.(string)

// 		matchQueries := map[string]types.MatchQuery{
// 			eq.Field: {Query: value},
// 		}
// 		// matchQueries := map[string]types.MatchQuery{
// 		// 	eq.Field: {Query: value},
// 		// }

// 		equalsQueries = append(equalsQueries, types.Query{
// 			Match: matchQueries,
// 		})

// 		// equalsQueries = append(equalsQueries, types.Query{
// 		// 	Match: map[string]types.MatchQuery{
// 		// 		eq.Field: {Query: eq.Value.(string)},
// 		// 	},
// 		// })
// 	}

// 	// return types.Query{
// 	// 	Bool: &types.BoolQuery{
// 	// 		Must: equalsQueries,
// 	// 	},
// 	// }

// 	query := types.Query{}

// 	boolQuery := types.BoolQuery{}
// 	boolQuery.Must = equalsQueries

// 	query.Bool = &boolQuery
// 	return query
// }
