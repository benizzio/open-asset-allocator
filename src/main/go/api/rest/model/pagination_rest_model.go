package model

// PaginatedResponse wraps one page of REST items and its optional opaque continuation cursor.
// The nextCursor property is always serialized; nil means that the current page is terminal.
// For example, `PaginatedResponse[string]{Items: []string{}, NextCursor: nil}` serializes nextCursor as null.
// Authored by: OpenCode
type PaginatedResponse[T any] struct {
	Items      []T     `json:"items"`
	NextCursor *string `json:"nextCursor"`
}
