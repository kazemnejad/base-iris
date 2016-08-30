package responses

import "strconv"

type AuthResponse struct {
	GeneralResponse
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}

func NewAuth(userId int64, token string) AuthResponse {
	return AuthResponse{
		GeneralResponse: NewGeneral("success"),
		UserId:          strconv.FormatInt(userId, 10),
		Token:           token,
	}
}
