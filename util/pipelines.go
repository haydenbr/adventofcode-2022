package util

type Pipeline[TStartingValue any] struct{}

func NewPipeline[TStartingValue any]() *Pipeline[TStartingValue] {
	return &Pipeline[TStartingValue]{}
}

func (p *Pipeline[T]) Aggregate(aggregateName string, zero T, f func()) {

}
