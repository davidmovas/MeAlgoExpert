package RemoveDuplicatesFromLinkedList

import (
	"fmt"
	"testing"
)

func Test_RemoveDuplicatesFromLinkedList(t *testing.T) {
	testFunctions := map[string]func(head *ListNode) *ListNode{
		"RemoveDuplicatesFromLinkedList_V1": RemoveDuplicatesFromLinkedList,
	}

	type args struct {
		head *ListNode
	}

	tests := []struct {
		args args
		want *ListNode
	}{
		{
			args: args{
				head: &ListNode{
					value: 1,
					next: &ListNode{
						value: 1,
						next: &ListNode{
							value: 3,
							next: &ListNode{
								value: 4,
								next: &ListNode{
									value: 4,
									next: &ListNode{
										value: 4,
										next: &ListNode{
											value: 5,
											next: &ListNode{
												value: 6,
												next: &ListNode{
													value: 6,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			want: &ListNode{
				value: 1,
				next: &ListNode{
					value: 3,
					next: &ListNode{
						value: 4,
						next: &ListNode{
							value: 5,
							next: &ListNode{
								value: 6,
							},
						},
					},
				},
			},
		},
	}

	for name, fn := range testFunctions {
		t.Run(name, func(t *testing.T) {
			for _, tt := range tests {
				if got := fn(tt.args.head); !CompareList(tt.args.head, got) {
					t.Errorf("%s = %v, want %v", name, got.value, tt.want.value)
				}
			}
		})
	}
}

func CompareList(first, second *ListNode) bool {
	currentF := first
	currentS := second

	for currentF != nil && currentS != nil {
		fmt.Printf("F: %d S: %d\n", currentF.value, currentS.value)
		if currentF.value != currentS.value {
			return false
		}
		currentF = currentF.next
		currentS = currentS.next
	}

	return true
}
