package repositories

import (
	"log"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"koronet_web_app/entities"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) EnsureDatabaseExists() error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(50) UNIQUE NOT NULL,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		password_hash VARCHAR(255) NOT NULL
	);`

	_, err := ur.db.Exec(query)
	if err != nil {
		log.Println("Error creating users table:", err)
		return err
	}

	return nil
}

func (ur *UserRepository) Create(user *entities.User) error {
	query := "INSERT INTO users (username, name, email, password_hash) VALUES (?, ?, ?, ?)"
	stmt, err := ur.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing query:", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.Name, user.Email, user.Password)
	if err != nil {
		log.Println("Error executing query:", err)
		return err
	}

	return nil
}

func (ur *UserRepository) Delete(userID string) error {

	query := "DELETE FROM users WHERE id = ?"
	stmt, err := ur.db.Prepare(query)
	if err != nil {
		log.Println("Error preparing delete query:", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID)
	if err != nil {
		log.Println("Error executing delete query:", err)
		return err
	}

	return nil
}

func (ur *UserRepository) Get(userID string) (*entities.User, error) {
	query := "SELECT id, username, name, email, password_hash FROM users WHERE id = ?"
	row := ur.db.QueryRow(query, userID)

	var user entities.User
	err := row.Scan(&user.ID, &user.Username, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("User not found")
			return nil, nil
		}
		log.Println("Error scanning row:", err)
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) GetAll() ([]*entities.User, error) {
	query := "SELECT * FROM users"
	rows, err := ur.db.Query(query)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var users []*entities.User
	for rows.Next() {
		var user entities.User
		err := rows.Scan(&user.ID, &user.Username, &user.Name, &user.Email, &user.Password)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		log.Println("Error in result set:", err)
		return nil, err
	}

	return users, nil
}
