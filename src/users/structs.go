package users

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

var users = []User{
	{Id: "1", Name: "Alexey", Email: "alexey@gmail.com", Password: "Alexey"},
	{Id: "2", Name: "John", Email: "john@gmail.com", Password: "John"},
	{Id: "3", Name: "Nikita", Email: "nikita@gmail.com", Password: "Nikita"},
	{Id: "4", Name: "Pavel", Email: "pavel@gmail.com", Password: "Pavel"},
}
