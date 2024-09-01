package entities

import (
	"github.com/NViktorovich/testing_platform_service/internal/entities/strategy"
	"github.com/pkg/errors"
)

type Topic string

type Question struct {
	id             int
	qType          strategy.QuestionType
	topic          Topic
	question       string
	variants       []string
	userAnswers    []string
	correctAnswers []string
	resolver       strategy.ResolveStrategy
	result         struct {
		status  bool
		correct bool
	}
}

func NewQuestion(
	id int,
	questionType strategy.QuestionType,
	resolver strategy.ResolveStrategy,
	topic Topic,
	question string,
	variants []string,
	correctAnswers []string,
) (*Question, error) {
	if id == 0 {
		return nil, errors.Wrap(ErrInvalidParam, "id must not be 0")
	}
	if _, err := questionType.Type(); err != nil {
		return nil, errors.Wrap(ErrInvalidParam, err.Error())
	}
	if resolver == nil {
		return nil, errors.Wrap(ErrInvalidParam, "resolve strategy is empty")
	}
	if topic == "" {
		return nil, errors.Wrap(ErrInvalidParam, "topic is empty")
	}
	if question == "" {
		return nil, errors.Wrap(ErrInvalidParam, "question is empty")
	}
	if len(variants) == 0 {
		return nil, errors.Wrap(ErrInvalidParam, "variants is empty")
	}
	if len(correctAnswers) == 0 {
		return nil, errors.Wrap(ErrInvalidParam, "correctAnswers is empty")
	}

	return &Question{
		id:             id,
		qType:          questionType,
		topic:          topic,
		question:       question,
		variants:       variants,
		correctAnswers: correctAnswers,
		resolver:       resolver,
	}, nil
}

func (q *Question) SetUserAnswers(answers []string) {
	q.userAnswers = answers
}

func (q *Question) GetResult() (bool, error) {
	if !q.result.status {
		res, err := q.resolver.Resolve(q.correctAnswers, q.userAnswers)
		if err != nil {
			return false, errors.Wrap(err, "failed to resolve question")
		}
		q.result.status = true
		q.result.correct = res
	}

	return q.result.correct, nil
}
