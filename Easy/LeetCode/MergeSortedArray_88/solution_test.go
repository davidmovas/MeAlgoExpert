package MergeSortedArray_88

import (
	"reflect"
	"testing"
)

func Test_MergeSortedArray(t *testing.T) {
	testFunctions := map[string]func(nums1, nums2 []int, m, n int) []int{
		"MergeSortedArray": MergeSortedArray,
	}

	type args struct {
		nums1, nums2 []int
		m, n         int
	}

	tests := []struct {
		args args
		want []int
	}{
		{
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				nums2: []int{2, 5, 6},
				m:     3,
				n:     3,
			},
			want: []int{1, 2, 2, 3, 5, 6},
		},
		{
			args: args{
				nums1: []int{1},
				nums2: []int{},
				m:     1,
				n:     0,
			},
			want: []int{1},
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.nums1, tt.args.nums2, tt.args.m, tt.args.n); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
