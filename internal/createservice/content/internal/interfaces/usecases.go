package interfaces

const UsecasesContent = `package interfaces

// TODO add your methods inside usecases
//go:generate mockgen -source=usecases.go -destination=../../mocks/usecases/usecases.go -package=mockusecases
type UseCases interface {
	<service-name-title>Service
}
`
