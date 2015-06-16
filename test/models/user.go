package models

type User struct {
	Id       int    `json:"id"`
	UserName string `json:"user_name"`
	Location string `json:"location"`
}

func (u *User) Load(id int) *User {
	if id != 0 {
		u.Id = id
		u.UserName = "BrandonRomano"
		u.Location = "Brooklyn, NY"
	}
	return u
}
