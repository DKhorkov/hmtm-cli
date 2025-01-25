package internal

const UsecasesContent = `package usecases

import (
	"github.com/DKhorkov/hmtm-<service-name>/internal/interfaces"
)

func NewCommonUseCases(
	service interfaces.<service-name-title>Service,
) *CommonUseCases {
	return &CommonUseCases{
		<service-name>Service:  service,
	}
}

type CommonUseCases struct {
	<service-name>Service  interfaces.<service-name-title>Service
}
`
