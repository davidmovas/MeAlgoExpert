package ClosestValueInBST

import (
	"reflect"
	"testing"
)

func Test_FindClosetsValueInBST(t *testing.T) {
	testFunctions := map[string]func(node *Node, target int) int{
		"FindClosetsValueInBST_V1": FindClosestValueInBST_V1,
	}

	type args struct {
		node   *Node
		target int
	}

	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				node: &Node{
					value: 10,
					left: &Node{
						value: 5,
						left: &Node{
							value: 2,
							left: &Node{
								value: 1,
							},
						},
						right: &Node{
							value: 5,
						},
					},
					right: &Node{
						value: 15,
						left: &Node{
							value: 13,
							right: &Node{
								value: 14,
							},
						},
						right: &Node{
							value: 22,
						},
					},
				},
				target: 12,
			},
			want: 13,
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.node, tt.args.target); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
