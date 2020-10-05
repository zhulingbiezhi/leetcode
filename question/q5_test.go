package question

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test 1",
			args: args{
				s: "abdcd",
			},
			want: "dcd",
		},
		{
			name: "test 2",
			args: args{
				s: "a",
			},
			want: "a",
		},
		{
			name: "test 3",
			args: args{
				s: "aba",
			},
			want: "aba",
		},
		{
			name: "test 4",
			args: args{
				s: "abaabaa",
			},
			want: "abaaba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
