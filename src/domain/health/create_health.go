package health

//CreateHealthResponse return http status and description
type CreateHealthResponse struct {
	Status      int    `json:"status"`
	Description string `json:"description"`
}
