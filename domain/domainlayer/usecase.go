package domainlayer

type DomainLayer struct {
	storage Storage
}

func (dl DomainLayer) Create() error {
	return dl.storage.Create()
}
