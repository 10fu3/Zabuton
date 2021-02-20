package Entity

type user struct {
	uuid string `gorm:"primaryKey"`
	name string
	thumbnail string
}

type User interface {
	GetUUID() string
	GetName() string
	GetThumbnail() string
}

func (u *user) GetUUID() string {
	return u.uuid
}

func (u *user) GetName() string {
	return u.name
}

func (u *user) GetThumbnail() string {
	return u.thumbnail
}

type UserBuilder struct {
	u user
}

func EditUser(u User) *UserBuilder {
	 a := UserBuilder{}
	 return a.SetUUID(u.GetUUID()).
	          SetThumbnail(u.GetThumbnail()).
	          SetName(u.GetName())
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{}
}

func (u *UserBuilder) SetUUID(uuid string) *UserBuilder {
	u.u.uuid = uuid
	return u
}

func (u *UserBuilder) SetName(name string) *UserBuilder {
	u.u.name = name
	return u
}

func (u *UserBuilder) SetThumbnail(thumbnail string) *UserBuilder {
	u.u.thumbnail = thumbnail
	return u
}

func (u *UserBuilder) Build() User {
 	return &u.u
}


