package interview

import "testing"

func TestIsInArray(t *testing.T) {
	type args struct {
		arr    [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test 1",
			args: args{
				arr: [][]int{
					{
						1, 2, 3,
					},
					{
						4, 6, 6,
					},
					{
						7, 8, 9,
					},
				},
				target: 5,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsInArray(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("IsInArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
