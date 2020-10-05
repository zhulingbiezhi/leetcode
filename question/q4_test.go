package question

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test 1",
			args: args{
				nums1: []int{1, 3},
				nums2: []int{2},
			},
			want: 2,
		},
		{
			name: "test 2",
			args: args{
				nums1: []int{1, 2, 3},
				nums2: []int{2, 4},
			},
			want: 2,
		},
		{
			name: "test 3",
			args: args{
				nums1: []int{1, 2, 3},
				nums2: []int{10, 14},
			},
			want: 3,
		},
		{
			name: "test 4",
			args: args{
				nums1: []int{1, 3},
				nums2: []int{2},
			},
			want: 2,
		},
		{
			name: "test 5",
			args: args{
				nums1: []int{1, 2, 6},
				nums2: []int{4, 10, 14},
			},
			want: 5,
		},
		{
			name: "test 6",
			args: args{
				nums1: []int{1},
				nums2: []int{2},
			},
			want: 1.5,
		},
		{
			name: "test 7",
			args: args{
				nums1: []int{1},
				nums2: []int{},
			},
			want: 1,
		},
		{
			name: "test 8",
			args: args{
				nums1: []int{1, 3},
				nums2: []int{2, 7},
			},
			want: 2.5,
		},
		{
			name: "test 9",
			args: args{
				nums1: []int{},
				nums2: []int{2, 3},
			},
			want: 2.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays1() = %v, want %v", got, tt.want)
			}
		})
	}
}
