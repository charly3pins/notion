package notion

// Pagination allows an integration to request a part of the list, receiving an array of results and a next_cursor in the response.
// The integration can use the next_cursor in another request to receive the next part of the list.
type Pagination struct {
	NextCursor string `json:"next_cursor,omitempty"`
	HasMore    bool   `json:"has_more"`
}
