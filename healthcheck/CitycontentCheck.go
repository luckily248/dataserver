package healthcheck

import (
	"../models"
)

type CitycontentCheck struct {
}

func (cc *CitycontentCheck) Check() error {
	return models.CheckCitycontentInstance()
}
