package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/vicpoo/apiShop/src/core"
	"github.com/vicpoo/apiShop/src/clothes/domain"
	"github.com/vicpoo/apiShop/src/clothes/domain/entities"
)

type MySQLClothRepository struct {
	conn *sql.DB
}

func NewMySQLClothRepository() domain.IClothRepository {
	conn := core.GetBD()
	return &MySQLClothRepository{conn: conn}
}

func (mysql *MySQLClothRepository) Save(cloth *entities.Cloth) error {
	query := `
		INSERT INTO clothes (
			name, description, size, price, stock, imagen_url
		)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	
	result, err := mysql.conn.Exec(query,
		cloth.GetName(),
		mysql.nullString(cloth.GetDescription()),
		mysql.nullString(cloth.GetSize()),
		mysql.nullFloat64(cloth.GetPrice()),
		mysql.nullInt32(cloth.GetStock()),
		mysql.nullString(cloth.GetImageURL()),
	)
	
	if err != nil {
		log.Println("Error al guardar la prenda:", err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener ID generado:", err)
		return err
	}
	cloth.SetIDCloth(int32(id))

	return nil
}

func (mysql *MySQLClothRepository) Update(cloth *entities.Cloth) error {
	query := `
		UPDATE clothes
		SET name = ?, description = ?, size = ?, 
		    price = ?, stock = ?, imagen_url = ?
		WHERE id_clothes = ?
	`
	
	result, err := mysql.conn.Exec(query,
		cloth.GetName(),
		mysql.nullString(cloth.GetDescription()),
		mysql.nullString(cloth.GetSize()),
		mysql.nullFloat64(cloth.GetPrice()),
		mysql.nullInt32(cloth.GetStock()),
		mysql.nullString(cloth.GetImageURL()),
		cloth.GetIDCloth(),
	)
	
	if err != nil {
		log.Println("Error al actualizar la prenda:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("prenda con ID %d no encontrada", cloth.GetIDCloth())
	}

	return nil
}

func (mysql *MySQLClothRepository) Delete(id int32) error {
	query := "DELETE FROM clothes WHERE id_clothes = ?"
	result, err := mysql.conn.Exec(query, id)
	if err != nil {
		log.Println("Error al eliminar la prenda:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("prenda con ID %d no encontrada", id)
	}

	return nil
}

func (mysql *MySQLClothRepository) GetByID(id int32) (*entities.Cloth, error) {
	query := `
		SELECT 
			id_clothes, name, description, size, 
			price, stock, imagen_url
		FROM clothes
		WHERE id_clothes = ?
	`
	row := mysql.conn.QueryRow(query, id)

	var cloth entities.Cloth
	var (
		description sql.NullString
		size        sql.NullString
		price       sql.NullFloat64
		stock       sql.NullInt32
		imageURL    sql.NullString
	)
	
	err := row.Scan(
		&cloth.IDCloth,
		&cloth.Name,
		&description,
		&size,
		&price,
		&stock,
		&imageURL,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("prenda con ID %d no encontrada", id)
		}
		log.Println("Error al buscar la prenda por ID:", err)
		return nil, err
	}

	// Asignar valores nulos si existen
	if description.Valid {
		cloth.SetDescription(description.String)
	}
	if size.Valid {
		cloth.SetSize(size.String)
	}
	if price.Valid {
		cloth.SetPrice(price.Float64)
	}
	if stock.Valid {
		cloth.SetStock(stock.Int32)
	}
	if imageURL.Valid {
		cloth.SetImageURL(imageURL.String)
	}

	return &cloth, nil
}

func (mysql *MySQLClothRepository) GetAll() ([]entities.Cloth, error) {
	query := `
		SELECT 
			id_clothes, name, description, size, 
			price, stock, imagen_url
		FROM clothes
		ORDER BY name
	`
	rows, err := mysql.conn.Query(query)
	if err != nil {
		log.Println("Error al obtener todas las prendas:", err)
		return nil, err
	}
	defer rows.Close()

	var clothes []entities.Cloth
	for rows.Next() {
		var cloth entities.Cloth
		var (
			description sql.NullString
			size        sql.NullString
			price       sql.NullFloat64
			stock       sql.NullInt32
			imageURL    sql.NullString
		)
		
		err := rows.Scan(
			&cloth.IDCloth,
			&cloth.Name,
			&description,
			&size,
			&price,
			&stock,
			&imageURL,
		)
		
		if err != nil {
			log.Println("Error al escanear la prenda:", err)
			return nil, err
		}

		// Asignar valores nulos si existen
		if description.Valid {
			cloth.SetDescription(description.String)
		}
		if size.Valid {
			cloth.SetSize(size.String)
		}
		if price.Valid {
			cloth.SetPrice(price.Float64)
		}
		if stock.Valid {
			cloth.SetStock(stock.Int32)
		}
		if imageURL.Valid {
			cloth.SetImageURL(imageURL.String)
		}

		clothes = append(clothes, cloth)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return clothes, nil
}

func (mysql *MySQLClothRepository) FindByName(name string) ([]entities.Cloth, error) {
	query := `
		SELECT 
			id_clothes, name, description, size, 
			price, stock, imagen_url
		FROM clothes
		WHERE name LIKE ?
		ORDER BY name
	`
	
	rows, err := mysql.conn.Query(query, "%"+name+"%")
	if err != nil {
		log.Println("Error al buscar prendas por nombre:", err)
		return nil, err
	}
	defer rows.Close()

	return mysql.scanClothes(rows)
}

func (mysql *MySQLClothRepository) FindBySize(size string) ([]entities.Cloth, error) {
	query := `
		SELECT 
			id_clothes, name, description, size, 
			price, stock, imagen_url
		FROM clothes
		WHERE size = ?
		ORDER BY name
	`
	
	rows, err := mysql.conn.Query(query, size)
	if err != nil {
		log.Println("Error al buscar prendas por talla:", err)
		return nil, err
	}
	defer rows.Close()

	return mysql.scanClothes(rows)
}

func (mysql *MySQLClothRepository) FindByPriceRange(minPrice, maxPrice float64) ([]entities.Cloth, error) {
	query := `
		SELECT 
			id_clothes, name, description, size, 
			price, stock, imagen_url
		FROM clothes
		WHERE price BETWEEN ? AND ?
		ORDER BY price, name
	`
	
	rows, err := mysql.conn.Query(query, minPrice, maxPrice)
	if err != nil {
		log.Println("Error al buscar prendas por rango de precio:", err)
		return nil, err
	}
	defer rows.Close()

	return mysql.scanClothes(rows)
}

func (mysql *MySQLClothRepository) scanClothes(rows *sql.Rows) ([]entities.Cloth, error) {
	var clothes []entities.Cloth
	
	for rows.Next() {
		var cloth entities.Cloth
		var (
			description sql.NullString
			size        sql.NullString
			price       sql.NullFloat64
			stock       sql.NullInt32
			imageURL    sql.NullString
		)
		
		err := rows.Scan(
			&cloth.IDCloth,
			&cloth.Name,
			&description,
			&size,
			&price,
			&stock,
			&imageURL,
		)
		
		if err != nil {
			log.Println("Error al escanear la prenda:", err)
			return nil, err
		}

		if description.Valid {
			cloth.SetDescription(description.String)
		}
		if size.Valid {
			cloth.SetSize(size.String)
		}
		if price.Valid {
			cloth.SetPrice(price.Float64)
		}
		if stock.Valid {
			cloth.SetStock(stock.Int32)
		}
		if imageURL.Valid {
			cloth.SetImageURL(imageURL.String)
		}

		clothes = append(clothes, cloth)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return clothes, nil
}

func (mysql *MySQLClothRepository) nullString(s string) interface{} {
	if s == "" {
		return sql.NullString{}
	}
	return s
}

func (mysql *MySQLClothRepository) nullInt32(i int32) interface{} {
	if i == 0 {
		return sql.NullInt32{}
	}
	return i
}

func (mysql *MySQLClothRepository) nullFloat64(f float64) interface{} {
	if f == 0.0 {
		return sql.NullFloat64{}
	}
	return f
}