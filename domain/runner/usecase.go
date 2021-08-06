package runner

type Runner struct {
	layers UseCaseLayers
}

func (r Runner) Run() {
	for _, layer := range r.layers {
		if err := layer.Create(); err != nil {
			// handle error
		}
	}
}
