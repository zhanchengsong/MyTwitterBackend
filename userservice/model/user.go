package model

type User struct {
	DisplayName string   `json:"displayName"`
	Username    string   `json:"username"`
	Email       string   `json:"email"`
	Password    string   `json:"Password"`
	ID          string   `json:-`
	Friends     []string `json:-`
}
