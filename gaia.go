package gaia

// User is the user object
type User struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	DisplayName string `json:"display_name"`
	Tokenstring string `json:"tokenstring"`
}