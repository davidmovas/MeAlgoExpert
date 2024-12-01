package EvaluateExpressionTree

import (
	"reflect"
	"testing"
)

func Test_EvaluateExpressionTree(t *testing.T) {
	testFunctions := map[string]func(root *Node) int{
		"EvaluateExpressionTree_V1": EvaluateExpressionTree_V1,
	}

	type args struct {
		node *Node
	}

	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				node: &Node{
					value: -1,
					left: &Node{
						value: -2,
						left: &Node{
							value: -4,
							left: &Node{
								value: 2,
							},
							right: &Node{
								value: 3,
							},
						},
						right: &Node{
							value: 2,
						},
					},
					right: &Node{
						value: -3,
						left: &Node{
							value: 8,
						},
						right: &Node{
							value: 3,
						},
					},
				},
			},
			want: 6,
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
