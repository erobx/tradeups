package user

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Hash  string `json:"hash"`
}
