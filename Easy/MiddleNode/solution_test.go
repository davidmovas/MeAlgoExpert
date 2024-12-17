package MiddleNode

import (
	. "algoexpert/Easy/models"
	. "algoexpert/models"
	"reflect"
	"testing"
)

func Test_MiddleNode(t *testing.T) {
	testFunctions := map[string]func(head *LinkedList[int]) *LinkedList[int]{
		"MiddleNode_V1": MiddleNode,
	}

	type args struct {
		head *LinkedList[int]
	}

	tests := []struct {
		args args
		want *LinkedList[int]
	}{
		{
			args: args{
				head: &LinkedList[int]{
					Value: 2,
					Next: &LinkedList[int]{
						Value: 7,
						Next: &LinkedList[int]{
							Value: 3,
							Next: &LinkedList[int]{
								Value: 5,
							},
						},
					},
				},
			},
			want: &LinkedList[int]{
				Value: 3,
			},
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.head); !reflect.DeepEqual(got.Value, tt.want.Value) {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			}
		})
	}
}
