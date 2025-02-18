package internal

const UsecasesContent = `package usecases

import (
	"github.com/DKhorkov/hmtm-<service-name>/internal/interfaces"
)

func New(
	service interfaces.<service-name-title>Service,
) *UseCases {
	return &UseCases{
		<service-name>Service:  service,
	}
}

type UseCases struct {
	<service-name>Service  interfaces.<service-name-title>Service
}
`
