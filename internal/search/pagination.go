package search

// Paginate splits search results into smaller pages
// This simulates LinkedIn-style pagination for safer processing
func Paginate(results []SearchResult, pageSize int) [][]SearchResult {

	// Holds all paginated result sets
	var pages [][]SearchResult

	// -------------------------------------------------
	// SPLIT RESULTS INTO FIXED-SIZE CHUNKS
	// -------------------------------------------------
	// Continue slicing until remaining results are less than pageSize
	for pageSize < len(results) {

		// Append one page and move forward
		results, pages = results[pageSize:], append(
			pages,
			results[0:pageSize:pageSize],
		)
	}

	// Append remaining results as the last page
	pages = append(pages, results)

	return pages
}
