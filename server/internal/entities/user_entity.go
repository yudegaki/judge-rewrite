package entities

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (u *User) GetID() uint {
	return u.ID
}

func (u *User) SetID(id uint) {
	u.ID = id
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) SetName(name string) {
	u.Name = name
}
