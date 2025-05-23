package internal

const ServicesContent = `package services

import (
	"github.com/DKhorkov/libs/logging"

	"github.com/DKhorkov/hmtm-<service-name>/internal/interfaces"
)

func New<service-name-title>Service(
	<service-name>Repository  interfaces.<service-name-title>Repository,
	logger            		  logging.Logger,
) *<service-name-title>Service {
	return &<service-name-title>Service{
		<service-name>Repository:  <service-name>Repository,
		logger:            		   logger,
	}
}

type <service-name-title>Service struct {
	<service-name>Repository  interfaces.<service-name-title>Repository
	logger            		  logging.Logger
}
`
