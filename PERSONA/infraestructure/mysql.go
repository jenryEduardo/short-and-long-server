package infraestructure

import (
	"recu/PERSONA/domain"
	"recu/PERSONA/core"

)

type MySQLRepository struct {
	conn *core.Conn_MySQL
}

func NewMySQLRepository() *MySQLRepository {
	conn := core.GetDBPool()
	return &MySQLRepository{conn: conn}
}

func (r *MySQLRepository) Save(p *domain.Person) error {
	query := "INSERT INTO personas (edad, nombre, sexo) VALUES(?, ?, ?)"
	_, err := r.conn.DB.Exec(query,&p.Edad,&p.Nombre,&p.Sexo)
	return err
}

func (r *MySQLRepository) GetAll() ([]domain.Person, error) {
	query := "SELECT edad, nombre, sexo FROM personas ORDER BY id DESC LIMIT 5"
	rows, err := r.conn.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var people []domain.Person
	for rows.Next() {
		var person domain.Person
		if err := rows.Scan(&person.Edad, &person.Nombre, &person.Sexo); err != nil {
			return nil, err
		}
		people = append(people, person)
	}
	return people, nil
}


func (r *MySQLRepository)GetCountM()(int,error){
    var total int

    query := "SELECT COUNT(*) FROM personas WHERE sexo = ?"
    err := r.conn.DB.QueryRow(query, "M").Scan(&total)
    if err != nil {
        return 0, err
    }

    return total, nil
}

func (r *MySQLRepository) GetCountH() (int, error) {
    var total int

    query := "SELECT COUNT(*) FROM personas WHERE sexo = ?"
    err := r.conn.DB.QueryRow(query, "H").Scan(&total)
    if err != nil {
        return 0, err
    }

    return total, nil
}

