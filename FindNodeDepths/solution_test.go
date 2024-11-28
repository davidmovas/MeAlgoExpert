package FindNodeDepths

import (
	"reflect"
	"testing"
)

func Test_FindNodeDepths(t *testing.T) {
	testFunctions := map[string]func(node *Node) int{
		"FindNodeDepths_V1": FindNodeDepthsInBTS,
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
			want: 16,
		},
		{
			args: args{
				node: &Node{
					value: 10,
					left: &Node{
						value: 5,
						left: &Node{
							value: 3,
							left: &Node{
								value: 1,
							},
						},
						right: &Node{
							value: 7,
							right: &Node{
								value: 8,
							},
						},
					},
					right: &Node{
						value: 15,
						left: &Node{
							value: 13,
						},
						right: &Node{
							value: 18,
						},
					},
				},
			},
			want: 16,
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
