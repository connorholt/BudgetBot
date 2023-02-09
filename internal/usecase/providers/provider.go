package providers

type Provider interface {
	GetValue() int
}

type Composite struct {
}

func (c Composite) GetValue() int {
	return 0
}

func NewProvider() Provider {
	return &Composite{}
}
