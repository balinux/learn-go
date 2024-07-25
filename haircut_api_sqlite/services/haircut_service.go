package services

import (
	"database/sql"
	"haircut_api/databases"
	"haircut_api/models"
)

var haircuts = []models.Haircut{
	{ID: 1, Name: "Buzz Cut", Description: "tes Description"},
}

func GetAllHaircut() ([]models.Haircut, error) {
	rows, err := databases.DB.Query("SELECT id,name,description,price FROM haircuts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var haircuts []models.Haircut

	for rows.Next() {
		var haircut models.Haircut
		err := rows.Scan(&haircut.ID, &haircut.Name, &haircut.Description, &haircut.Price)
		if err != nil {
			return nil, err
		}
		haircuts = append(haircuts, haircut)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return haircuts, err
}

func GetHaircutByID(id int) (*models.Haircut, error) {
	row := databases.DB.QueryRow("SELECT id, name, description, price FROM haircuts WHERE id = ?", id)

	var haircut models.Haircut
	err := row.Scan(&haircut.ID, &haircut.Name, &haircut.Description, &haircut.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &haircut, nil
}

func AddHaircut(newHaircut models.Haircut) (int64, error) {
	result, err := databases.DB.Exec("INSERT INTO haircuts (name,description,price) VALUES(?,?,?)", newHaircut.Name, newHaircut.Description, newHaircut.Price)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
