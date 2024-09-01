package strategy

//go:generate mockgen -source=./resolve_strategy.go -destination=./testdata/resolve_strategy.go --package=testdata

type ResolveStrategy interface {
	Resolve(correct, current interface{}) (bool, error)
}
