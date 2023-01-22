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
