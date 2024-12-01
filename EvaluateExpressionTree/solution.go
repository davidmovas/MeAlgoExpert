package EvaluateExpressionTree

type Node struct {
	value int
	left  *Node
	right *Node
}

func EvaluateExpressionTree_V1(root *Node) int {
	return recEvaluation(root.left, root.value) + recEvaluation(root.right, root.value)
}

func recEvaluation(node *Node, parentVal int) int {
	var leftVal, rightVal int
	if node == nil {
		return 0
	}

	if node.value > 0 {
		return node.value
	} else {
		leftVal = recEvaluation(node.left, node.value)
		rightVal = recEvaluation(node.right, node.value)
	}

	return makeResultOnOperation(node.value, leftVal, rightVal)
}

func makeResultOnOperation(operation, left, right int) int {
	switch operation {
	case -1:
		return left + right
	case -2:
		return left - right
	case -3:
		return left / right
	default:
		return left * right
	}
}
