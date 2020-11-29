package interview

import "testing"

func TestDistributedTasks(t *testing.T) {
	type args struct {
		key   string
		token string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test 1",
			args: args{
				key:   "B1vn2x1f",
				token: "s",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DistributedTasks(tt.args.key, tt.args.token)
		})
	}
}

func TestDealWithTask(t *testing.T) {
	type args struct {
		keys  string
		token string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test 1",
			args: args{
				keys:  "1,2,3,4,5,6,7,8,9",
				token: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := DealWithTask(tt.args.keys, tt.args.token); (err != nil) != tt.wantErr {
				t.Errorf("DealWithTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
