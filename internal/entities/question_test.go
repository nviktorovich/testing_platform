package entities_test

import (
	"testing"

	"github.com/NViktorovich/testing_platform_service/internal/entities"
	"github.com/NViktorovich/testing_platform_service/internal/entities/strategy"
	"github.com/NViktorovich/testing_platform_service/internal/entities/strategy/testdata"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

var (
	ErrTest = errors.New("test error")
)

//nolint:funlen //ok
func TestNewQuestion_Failure(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	resolver := testdata.NewMockResolveStrategy(ctrl)
	qid := 1
	qType := strategy.SingleAnswer
	topic := entities.Topic("test_topic")
	question := "2+2=?"
	variants := []string{"1", "2", "3", "4"}
	correctAnswer := []string{"4"}

	type args struct {
		id              int
		questionType    strategy.QuestionType
		resolveStrategy strategy.ResolveStrategy
		topic           entities.Topic
		question        string
		variants        []string
		correctAnswers  []string
	}
	tests := []struct {
		name   string
		args   args
		err    error
		reason string
	}{
		{
			name: "id validation failure",
			args: args{
				id:              0,
				question:        question,
				variants:        variants,
				correctAnswers:  correctAnswer,
				topic:           topic,
				questionType:    qType,
				resolveStrategy: resolver,
			},
			err:    entities.ErrInvalidParam,
			reason: "id must not be 0",
		},
		{
			name: "question type validation failure",
			args: args{
				id:              qid,
				question:        question,
				variants:        variants,
				correctAnswers:  correctAnswer,
				topic:           topic,
				questionType:    strategy.QuestionType(3),
				resolveStrategy: resolver,
			},
			err: entities.ErrInvalidParam,
		},
		{
			name: "topic type validation failure",
			args: args{
				id:              qid,
				question:        question,
				variants:        variants,
				correctAnswers:  correctAnswer,
				topic:           "",
				questionType:    strategy.SingleAnswer,
				resolveStrategy: resolver,
			},
			err:    entities.ErrInvalidParam,
			reason: "topic is empty",
		},
		{
			name: "question validation failure",
			args: args{
				id:              qid,
				question:        "",
				variants:        variants,
				correctAnswers:  correctAnswer,
				topic:           topic,
				questionType:    strategy.SingleAnswer,
				resolveStrategy: resolver,
			},
			err:    entities.ErrInvalidParam,
			reason: "question is empty",
		},
		{
			name: "variants validation failure",
			args: args{
				id:              qid,
				question:        question,
				variants:        nil,
				correctAnswers:  correctAnswer,
				topic:           topic,
				questionType:    strategy.SingleAnswer,
				resolveStrategy: resolver,
			},
			err:    entities.ErrInvalidParam,
			reason: "variants is empty",
		},
		{
			name: "correct answers validation failure",
			args: args{
				id:              qid,
				question:        question,
				variants:        variants,
				correctAnswers:  nil,
				topic:           topic,
				questionType:    strategy.SingleAnswer,
				resolveStrategy: resolver,
			},
			err:    entities.ErrInvalidParam,
			reason: "correctAnswers is empty",
		},
		{
			name: "resolve strategy failure",
			args: args{
				id:              qid,
				question:        question,
				variants:        variants,
				correctAnswers:  correctAnswer,
				topic:           topic,
				questionType:    strategy.SingleAnswer,
				resolveStrategy: nil,
			},
			err:    entities.ErrInvalidParam,
			reason: "resolve strategy is empty",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := entities.NewQuestion(
				tt.args.id,
				tt.args.questionType,
				tt.args.resolveStrategy,
				tt.args.topic,
				tt.args.question,
				tt.args.variants,
				tt.args.correctAnswers,
			)
			require.Nil(t, got)
			require.Contains(t, err.Error(), tt.reason)
		})
	}
}

//nolint:funlen //ok
func TestQuestionResolve(t *testing.T) {
	t.Parallel()

	type args struct {
		setResolverBehaviourFn func(t *testing.T, resolver *testdata.MockResolveStrategy, ca, ua []string)
		isExpectedError        bool
		expectedErr            error
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "successful resolve",
			args: args{
				setResolverBehaviourFn: setSuccessfulCorrectResolve,
				isExpectedError:        false,
				expectedErr:            nil,
			},
		},
		{
			name: "failure resolve",
			args: args{
				setResolverBehaviourFn: setFailureResolve,
				isExpectedError:        true,
				expectedErr:            ErrTest,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			resolver := testdata.NewMockResolveStrategy(ctrl)
			qid := 1
			qType := strategy.SingleAnswer
			topic := entities.Topic("test_topic")
			question := "2+2=?"
			variants := []string{"1", "2", "3", "4"}
			correctAnswer := []string{"4"}
			userAnswer := []string{"4"}

			tt.args.setResolverBehaviourFn(t, resolver, correctAnswer, userAnswer)
			q, err := entities.NewQuestion(
				qid,
				qType,
				resolver,
				topic,
				question,
				variants,
				correctAnswer,
			)
			require.NoError(t, err)
			require.NotNil(t, q)

			q.SetUserAnswers(userAnswer)
			isCorrect, err := q.GetResult()

			if tt.args.isExpectedError {
				require.False(t, isCorrect)
				require.ErrorIs(t, err, tt.args.expectedErr)
				return
			}
			require.True(t, isCorrect)
			require.NoError(t, err)
		})
	}
}

func setSuccessfulCorrectResolve(t *testing.T, resolver *testdata.MockResolveStrategy, ca, ua []string) {
	t.Helper()
	resolver.EXPECT().Resolve(ca, ua).Return(true, nil)
}

func setFailureResolve(t *testing.T, resolver *testdata.MockResolveStrategy, ca, ua []string) {
	t.Helper()
	resolver.EXPECT().Resolve(ca, ua).Return(false, ErrTest)
}
