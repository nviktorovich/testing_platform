package strategy

import (
	"github.com/pkg/errors"
)

var (
	ErrInvalidParam = errors.New("invalid param")
	ErrAssertion    = errors.New("assertion error")
)

type QuestionType int

const (
	SingleAnswer QuestionType = iota + 1
	MultiAnswer
)

func (qt QuestionType) Type() (string, error) {
	base := [...]string{"single-answer", "multi-answer"}
	if qt == 0 || int(qt) > len(base) {
		return "", errors.Wrap(ErrInvalidParam, "invalid question valuer")
	}
	return base[qt-1], nil
}

type Strategy struct {
	resolver ResolveStrategy
}

func NewStrategy(questionType QuestionType) (*Strategy, error) {
	switch questionType {
	case SingleAnswer:
		return &Strategy{
			resolver: NewOnlyOneTextAnswer(),
		}, nil
	case MultiAnswer:
		return &Strategy{
			resolver: NewMultiTextAnswer(),
		}, nil
	}

	return nil, errors.Wrap(ErrInvalidParam, "invalid question type")
}

func (s *Strategy) Resolve(correct, current interface{}) (bool, error) {
	return s.resolver.Resolve(correct, current)
}
