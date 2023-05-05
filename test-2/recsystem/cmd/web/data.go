package main

import "github.com/Aazan-Iqbal/3161/test-2/recsystem/internal/models"

type templateData struct {
	Equipment     *models.Equipment
	Flash         string
	EquipmentByte []*models.Equipment
	CSRFTOKEN     string // Added for authentication
}

// Anything CSRF related is in order to make sure you must be logged in to the system before accessing a page
// In practice it also prevent CSRF attack. ALSO FOR THE CSF TO WORK IT MUST HAVE A HIDDEN INPUT IN THE HTML
// poll.create.page.tmpl has an example at line 8
