// user_repository_impl.go
package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/vicpoo/apiShop/src/core"
	"github.com/vicpoo/apiShop/src/users/domain"
	"github.com/vicpoo/apiShop/src/users/domain/entities"
)

type MySQLUserRepository struct {
	conn *sql.DB
}

func NewMySQLUserRepository() domain.IUserRepository {
	conn := core.GetBD()
	return &MySQLUserRepository{conn: conn}
}

func (mysql *MySQLUserRepository) Save(user *entities.User) error {
	query := `
		INSERT INTO users (
			email, password, name, lastname
		)
		VALUES (?, ?, ?, ?)
	`
	
	result, err := mysql.conn.Exec(query,
		user.GetEmail(),
		user.GetPassword(),
		mysql.nullString(user.GetName()),
		mysql.nullString(user.GetLastname()),
	)
	
	if err != nil {
		log.Println("Error al guardar el usuario:", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener ID generado:", err)
		return err
	}
	user.SetIDUsuario(int32(id))

	return nil
}

func (mysql *MySQLUserRepository) Update(user *entities.User) error {
	query := `
		UPDATE users
		SET email = ?, password = ?, name = ?, lastname = ?
		WHERE id_usuario = ?
	`
	
	result, err := mysql.conn.Exec(query,
		user.GetEmail(),
		user.GetPassword(),
		mysql.nullString(user.GetName()),
		mysql.nullString(user.GetLastname()),
		user.GetIDUsuario(),
	)
	
	if err != nil {
		log.Println("Error al actualizar el usuario:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("usuario con ID %d no encontrado", user.GetIDUsuario())
	}

	return nil
}

func (mysql *MySQLUserRepository) Delete(id int32) error {
	query := "DELETE FROM users WHERE id_usuario = ?"
	result, err := mysql.conn.Exec(query, id)
	if err != nil {
		log.Println("Error al eliminar el usuario:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("usuario con ID %d no encontrado", id)
	}

	return nil
}

func (mysql *MySQLUserRepository) GetByID(id int32) (*entities.User, error) {
	query := `
		SELECT 
			id_usuario, email, password, name, lastname
		FROM users
		WHERE id_usuario = ?
	`
	row := mysql.conn.QueryRow(query, id)

	var user entities.User
	var (
		name     sql.NullString
		lastname sql.NullString
	)
	
	err := row.Scan(
		&user.IDUsuario,
		&user.Email,
		&user.Password,
		&name,
		&lastname,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("usuario con ID %d no encontrado", id)
		}
		log.Println("Error al buscar el usuario por ID:", err)
		return nil, err
	}

	// Asignar valores nulos si existen
	if name.Valid {
		user.SetName(name.String)
	}
	if lastname.Valid {
		user.SetLastname(lastname.String)
	}

	return &user, nil
}

func (mysql *MySQLUserRepository) GetAll() ([]entities.User, error) {
	query := `
		SELECT 
			id_usuario, email, password, name, lastname
		FROM users
		ORDER BY id_usuario
	`
	rows, err := mysql.conn.Query(query)
	if err != nil {
		log.Println("Error al obtener todos los usuarios:", err)
		return nil, err
	}
	defer rows.Close()

	var users []entities.User
	for rows.Next() {
		var user entities.User
		var (
			name     sql.NullString
			lastname sql.NullString
		)
		
		err := rows.Scan(
			&user.IDUsuario,
			&user.Email,
			&user.Password,
			&name,
			&lastname,
		)
		
		if err != nil {
			log.Println("Error al escanear el usuario:", err)
			return nil, err
		}

		// Asignar valores nulos si existen
		if name.Valid {
			user.SetName(name.String)
		}
		if lastname.Valid {
			user.SetLastname(lastname.String)
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return users, nil
}

func (mysql *MySQLUserRepository) Register(user *entities.User) error {
	// Primero verificamos si el email ya existe
	exists, err := mysql.ExistsByEmail(user.GetEmail())
	if err != nil {
		return err
	}
	
	if exists {
		return fmt.Errorf("el email %s ya está registrado", user.GetEmail())
	}

	// Si no existe, lo guardamos
	return mysql.Save(user)
}

func (mysql *MySQLUserRepository) Login(email, password string) (*entities.User, error) {
	query := `
		SELECT 
			id_usuario, email, password, name, lastname
		FROM users
		WHERE email = ? AND password = ?
	`
	row := mysql.conn.QueryRow(query, email, password)

	var user entities.User
	var (
		name     sql.NullString
		lastname sql.NullString
	)
	
	err := row.Scan(
		&user.IDUsuario,
		&user.Email,
		&user.Password,
		&name,
		&lastname,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("credenciales inválidas")
		}
		log.Println("Error al buscar el usuario para login:", err)
		return nil, err
	}

	// Asignar valores nulos si existen
	if name.Valid {
		user.SetName(name.String)
	}
	if lastname.Valid {
		user.SetLastname(lastname.String)
	}

	return &user, nil
}

// Métodos auxiliares
func (mysql *MySQLUserRepository) nullString(s string) interface{} {
	if s == "" {
		return sql.NullString{}
	}
	return s
}

func (mysql *MySQLUserRepository) ExistsByEmail(email string) (bool, error) {
	query := "SELECT COUNT(*) FROM users WHERE email = ?"
	row := mysql.conn.QueryRow(query, email)
	
	var count int
	err := row.Scan(&count)
	if err != nil {
		log.Println("Error al verificar existencia de email:", err)
		return false, err
	}
	
	return count > 0, nil
}