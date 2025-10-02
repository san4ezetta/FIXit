package repository

import (
	"FIXit/backend/internal/config"
	"context"
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type User struct {
	ID         int
	Name       string
	Surname    string
	Patronymic string
	Email      string
	Password   string
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(cfg *config.Config) *UserRepository {
	db, err := sql.Open("pgx", cfg.DatabaseURL)
	if err != nil {
		panic(err)
	}

	return &UserRepository{
		db: db,
	}
}

func getUserRepository() *UserRepository {
	return &UserRepository{}
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

func (s *UserRepository) Create(ctx context.Context, user User) error {
	const q = `INSERT INTO users (name, surname, patronymic, email, password)
               VALUES ($1, $2, $3, $4, $5)`

	stmt, err := s.db.PrepareContext(ctx, q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		user.Name, user.Surname, user.Patronymic, user.Email, user.Password)
	return err
}

func (s *UserRepository) FindById(ctx context.Context, id int) (*User, error) {
	q := `SELECT id, name, surname FROM users WHERE id = $1 LIMIT 1`
	var user User
	err := s.db.QueryRowContext(ctx, q, id).Scan(&user.ID, &user.Name, &user.Surname)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
	q := `SELECT id, name, surname, patronymic, email, password FROM users WHERE email = $1 LIMIT 1`
	var user User
	err := s.db.QueryRowContext(ctx, q, email).Scan(&user.ID, &user.Name, &user.Surname, &user.Patronymic, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
