package runner

type UseCaseLayer interface {
	Create() error
}

type UseCaseLayers []UseCaseLayer
