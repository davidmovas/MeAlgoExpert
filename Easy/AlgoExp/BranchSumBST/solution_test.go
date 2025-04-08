package BranchSumBST

import (
	"reflect"
	"testing"
)

func Test_FindClosetsValueInBST(t *testing.T) {
	testFunctions := map[string]func(node *Node) []int{
		"FindClosetsValueInBST_V1": CountBranchSumBST,
	}

	type args struct {
		node *Node
	}

	tests := []struct {
		args args
		want []int
	}{
		{
			args: args{
				node: &Node{
					value: 1,
					left: &Node{
						value: 2,
						left: &Node{
							value: 4,
							left: &Node{
								value: 8,
							},
							right: &Node{
								value: 9,
							},
						},
						right: &Node{
							value: 5,
							left: &Node{
								value: 10,
							},
						},
					},
					right: &Node{
						value: 3,
						left: &Node{
							value: 6,
						},
						right: &Node{
							value: 7,
						},
					},
				},
			},
			want: []int{15, 16, 18, 10, 11},
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.node); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
