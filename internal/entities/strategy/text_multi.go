package strategy

import (
	"github.com/pkg/errors"
)

type MultiTextAnswer struct{}

func NewMultiTextAnswer() *MultiTextAnswer {
	return &MultiTextAnswer{}
}

func (m *MultiTextAnswer) Resolve(correct, current interface{}) (bool, error) {
	correctAns, isCorrectSlice := correct.([]string)
	currentAns, isCurrentSlice := current.([]string)
	if !isCorrectSlice || !isCurrentSlice {
		return false, errors.Wrap(ErrAssertion, "assertion failed")
	}

	if len(correctAns) != len(currentAns) {
		return false, nil
	}
	if len(correctAns) < 2 {
		return false, errors.Wrapf(ErrInvalidParam,
			"least 2 correct answers in multi answers question, but got: %d", len(correctAns))
	}
	answersDict := make(map[string]struct{}, len(correctAns))
	for _, v := range currentAns {
		answersDict[v] = struct{}{}
	}

	for _, v := range correctAns {
		if _, ok := answersDict[v]; !ok {
			return false, nil
		}
	}
	return true, nil
}
