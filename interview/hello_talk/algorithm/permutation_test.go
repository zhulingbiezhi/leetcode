package algorithm

import "testing"

func TestHelloTalkInterview_Permutation(t *testing.T) {
	type fields struct {
		permutations []string
		used         map[int]bool
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "test 1",
			fields: fields{
				permutations: []string{},
				used:         make(map[int]bool),
			},
			args: args{
				s: "B1vn2x1f",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Algorithm{
				used: tt.fields.used,
			}
			a.Permutation(tt.args.s)
		})
	}
}
