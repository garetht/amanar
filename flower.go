package amanar

type Flower interface {
	Name() string
	UpdateWithCredentials(credentials *Credentials) error
	PersistChanges() error
}
