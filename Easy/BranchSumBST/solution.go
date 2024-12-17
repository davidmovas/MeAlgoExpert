package BranchSumBST

type Node struct {
	value int
	left  *Node
	right *Node
}

func CountBranchSumBST(node *Node) []int {
	result := make([]int, 0)

	sum := node.value

	countSum(node.left, sum, &result)
	countSum(node.right, sum, &result)

	return result
}

func countSum(node *Node, sum int, array *[]int) {
	sum += node.value

	if node.left != nil {
		countSum(node.left, sum, array)
	}

	if node.right != nil {
		countSum(node.right, sum, array)
	}

	if node.left == nil && node.right == nil {
		*array = append(*array, sum)
	}
}
