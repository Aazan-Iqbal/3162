package main

import "github.com/Aazan-Iqbal/3161/quiz-2/recsystem/internal/models"

type templateData struct {
	Equipment    *models.Equipment
	Flash        string
	ScheduleByte []*models.Equipment //used to hold byte data I guess
}
