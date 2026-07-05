package domain

type Pagination struct {
	TotalRecords int64	`json:"total_records"`
	TotalPages int	`json:"total_pages"`
	CurrentPage int	`json:"current_page"`
	Limit int	`json:"limit"`
	Data interface{}	`json:"data"`
}
