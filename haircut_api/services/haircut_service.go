package services

import "haircut_api/models"

var haircuts = []models.Haircut{
	{ID: 1, Name: "Buzz Cut", Description: "tes Description"},
}

func GetAllHaircut() []models.Haircut {
	return haircuts
}

func GetHaircutById(id int) *models.Haircut {
	for _, haircut := range haircuts {
		if haircut.ID == id {
			return &haircut
		}
	}
	return nil
}

func AddHaircut(newHaircut models.Haircut) {
	haircuts = append(haircuts, newHaircut)
}
