package swagmodels

type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
