package search

func Paginate(results []SearchResult, pageSize int) [][]SearchResult {
	var pages [][]SearchResult

	for pageSize < len(results) {
		results, pages = results[pageSize:], append(pages, results[0:pageSize:pageSize])
	}
	pages = append(pages, results)

	return pages
}
