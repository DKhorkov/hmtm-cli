package interfaces

const RepositoriesContent = `package interfaces

// TODO add your methods inside repository
//go:generate mockgen -source=repositories.go -destination=../../mocks/repositories/<service-name>_repository.go -package=mockrepositories
type <service-name-title>Repository interface {}
`
