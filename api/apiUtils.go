package api

const NotFound string = "NOT_FOUND"
const Success string = "SUCCESS"

type ApiErrorResponse struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}
