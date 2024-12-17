package FindNodeDepths

type Node struct {
	value int
	left  *Node
	right *Node
}

func FindNodeDepthsInBTS(node *Node) int {
	return countSum(node.left, 0, 0) + countSum(node.right, 0, 0)
}

func countSum(node *Node, level, sum int) int {
	level += 1

	if node == nil {
		return 0
	} else {
		sum += 1
	}

	sum += countSum(node.left, level, sum) + countSum(node.right, level, sum)

	return sum
}
