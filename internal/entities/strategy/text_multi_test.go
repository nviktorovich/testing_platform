package strategy_test

import (
	"testing"

	"github.com/NViktorovich/testing_platform_service/internal/entities/strategy"
	"github.com/stretchr/testify/require"
)

//nolint:funlen //ok
func TestMultiTextAnswer_Resolve(t *testing.T) {
	t.Parallel()

	type args struct {
		correct interface{}
		current interface{}
		err     error
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				correct: []string{"4", "5"},
				current: []string{"4", "5"},
				err:     nil,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "assertion fail",
			args: args{
				correct: "df",
				current: "...",
				err:     strategy.ErrAssertion,
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "different length of answers slices",
			args: args{
				correct: []string{"4", "5", "6"},
				current: []string{"4", "5"},
				err:     nil,
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "only one correct answer",
			args: args{
				correct: []string{"4"},
				current: []string{"4"},
				err:     strategy.ErrInvalidParam,
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "different answers slices",
			args: args{
				correct: []string{"4", "5"},
				current: []string{"4", "6"},
				err:     nil,
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			m := strategy.NewMultiTextAnswer()
			require.NotNil(t, m)

			got, err := m.Resolve(tt.args.correct, tt.args.current)
			require.Equal(t, tt.want, got)

			if tt.wantErr {
				require.ErrorIs(t, err, tt.args.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
