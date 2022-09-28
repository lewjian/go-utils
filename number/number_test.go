package number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoundInt(t *testing.T) {
	type args struct {
		number float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1.0",
			args: args{number: 1.0},
			want: 1,
		},
		{
			name: "1",
			args: args{number: 1},
			want: 1,
		},
		{
			name: "1.4",
			args: args{number: 1.4},
			want: 1,
		},
		{
			name: "1.421414",
			args: args{number: 1.421414},
			want: 1,
		},
		{
			name: "1.50",
			args: args{number: 1.50},
			want: 2,
		},
		{
			name: "1.499999",
			args: args{number: 1.499999},
			want: 1,
		},
		{
			name: "1.5000000001",
			args: args{number: 1.5000000001},
			want: 2,
		},
		{
			name: "1.999999999",
			args: args{number: 1.999999999},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, RoundInt(tt.args.number), "RoundInt(%v)", tt.args.number)
		})
	}
}

func TestRound(t *testing.T) {
	type args struct {
		num       float64
		precision uint
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "1.22222",
			args: args{
				num:       1.22222,
				precision: 2,
			},
			want: 1.22,
		},
		{
			name: "1.22422",
			args: args{
				num:       1.22422,
				precision: 2,
			},
			want: 1.22,
		},
		{
			name: "1.22522",
			args: args{
				num:       1.22522,
				precision: 2,
			},
			want: 1.23,
		},
		{
			name: "1.22922",
			args: args{
				num:       1.22922,
				precision: 2,
			},
			want: 1.23,
		},
		{
			name: "1.52922",
			args: args{
				num:       1.52922,
				precision: 0,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Round(tt.args.num, tt.args.precision), "Round(%v, %v)", tt.args.num, tt.args.precision)
		})
	}
}
