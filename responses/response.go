package responses

type GeneralResponse struct {
	Status string `json:"status"`
}

func NewGeneral(status string) GeneralResponse {
	return GeneralResponse{status}
}
