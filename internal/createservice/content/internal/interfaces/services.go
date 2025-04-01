package interfaces

const ServicesContent = `package interfaces

// TODO add your methods inside service
//go:generate mockgen -source=services.go -destination=../../mocks/services/<service-name>_service.go -package=mockservices
type <service-name-title>Service interface {
	<service-name-title>Repository
}
`
