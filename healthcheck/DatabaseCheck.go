package healthcheck

import (
	"../models"
)

type DatabaseCheck struct {
}

func (dbc *DatabaseCheck) Check() error {
	database := &models.BaseDBmodel{}
	return database.Check()
}
