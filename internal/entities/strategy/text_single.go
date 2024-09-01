package strategy

import (
	"github.com/pkg/errors"
)

type OnlyOneTextAnswer struct {
}

func NewOnlyOneTextAnswer() *OnlyOneTextAnswer {
	return &OnlyOneTextAnswer{}
}

func (o *OnlyOneTextAnswer) Resolve(correct, current interface{}) (bool, error) {
	correctAns, isCorrectSlice := correct.([]string)
	currentAns, isCurrentSlice := current.([]string)
	if !isCorrectSlice || !isCurrentSlice {
		return false, errors.Wrap(ErrAssertion, "assertion failed")
	}

	if err := o.validate(correctAns, currentAns); err != nil {
		return false, errors.Wrap(err, "validate")
	}
	return correctAns[0] == currentAns[0], nil
}

func (o *OnlyOneTextAnswer) validate(lists ...[]string) error {
	for _, l := range lists {
		if err := o.validateList(l); err != nil {
			return err
		}
	}
	return nil
}

func (o *OnlyOneTextAnswer) validateList(list []string) error {
	if len(list) != 1 {
		return errors.Wrapf(ErrInvalidParam, "list validate failure: %v", list)
	}
	return nil
}
