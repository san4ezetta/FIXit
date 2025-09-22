package repository

type User struct {
	Id         int
	Name       string
	Surname    string
	Patronymic string
	Email      string
	Password   string
}

type UserStore struct {
	users  map[int]*User
	nextId int
}

func NewUserStore() *UserStore {
	return &UserStore{
		users:  make(map[int]*User),
		nextId: 1,
	}
}

func (s *UserStore) Create(user User) error {
	user.Id = s.nextId
	s.users[s.nextId] = &user
	s.nextId++
	return nil
}

func (s *UserStore) FindById(id int) (*User, error) {
	return s.users[id], nil
}
