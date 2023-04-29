package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/Aazan-Iqbal/3161/quiz-2/recsystem/internal/models"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	//w.Write([]byte("Welcome to Polly!"))
	ts, err := template.ParseFiles(tmpl)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)

	}
}

func equipmentData(equipmentData *models.Equipment) *models.Equipment {
	return equipmentData
}
