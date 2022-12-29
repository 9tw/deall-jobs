package domain

type UserCore struct {
	ID       uint
	Fullname string
	Email    string
	Password string
	Role     uint
	Token    string
}

type Repository interface {
	GetMyUser(userID uint) (UserCore, error)
	Update(updatedUser UserCore, userID uint) (UserCore, error)
	Delete(userID uint) (UserCore, error)
	AddUser(newUser UserCore) (UserCore, error)
	GetUser(existUser UserCore) (UserCore, error)
	Show() ([]UserCore, error)
}

type Service interface {
	MyProfile(userID uint) (UserCore, error)
	UpdateProfile(updatedUser UserCore, userID uint) (UserCore, error)
	Deactivate(userID uint) (UserCore, error)
	Register(newUser UserCore) (UserCore, error)
	Login(existUser UserCore) (UserCore, error)
	ShowAll() ([]UserCore, error)
}
