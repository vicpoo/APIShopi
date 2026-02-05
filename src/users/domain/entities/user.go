//user.go
package entities

type User struct {
	IDUsuario int32   `json:"id_usuario" gorm:"column:id_usuario;primaryKey;autoIncrement"`
	Email     string  `json:"email" gorm:"column:email;unique;not null"`
	Password  string  `json:"password" gorm:"column:password;not null"`
	Name      *string `json:"name,omitempty" gorm:"column:name"`
	Lastname  *string `json:"lastname,omitempty" gorm:"column:lastname"`
}

// Setters
func (u *User) SetIDUsuario(id int32) {
	u.IDUsuario = id
}

func (u *User) SetEmail(email string) {
	u.Email = email
}

func (u *User) SetPassword(password string) {
	u.Password = password
}

func (u *User) SetName(name string) {
	u.Name = &name
}

func (u *User) SetLastname(lastname string) {
	u.Lastname = &lastname
}

// Getters
func (u *User) GetIDUsuario() int32 {
	return u.IDUsuario
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) GetName() string {
	if u.Name == nil {
		return ""
	}
	return *u.Name
}

func (u *User) GetLastname() string {
	if u.Lastname == nil {
		return ""
	}
	return *u.Lastname
}

// Constructor básico con campos requeridos
func NewUser(email, password string) *User {
	return &User{
		Email:    email,
		Password: password,
	}
}

// Constructor completo
func NewUserFull(
	email string,
	password string,
	name *string,
	lastname *string,
) *User {
	return &User{
		Email:     email,
		Password:  password,
		Name:      name,
		Lastname:  lastname,
	}
}