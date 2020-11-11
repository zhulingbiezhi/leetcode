package interview

import "testing"

func TestMinCoin(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 6",
			args: args{
				n: 6,
			},
			want: 2,
		},
		{
			name: "test 5",
			args: args{
				n: 5,
			},
			want: 1,
		},
		{
			name: "test 10",
			args: args{
				n: 10,
			},
			want: 2,
		},
		{
			name: "test 9",
			args: args{
				n: 9,
			},
			want: 2,
		},
		{
			name: "test 11",
			args: args{
				n: 11,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinCoin(tt.args.n); got != tt.want {
				t.Errorf("MinCoin() = %v, want %v", got, tt.want)
			}
		})
	}
}
