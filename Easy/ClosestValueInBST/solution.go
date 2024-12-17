package ClosestValueInBST

type Node struct {
	value int
	left  *Node
	right *Node
}

func FindClosestValueInBST_V1(node *Node, target int) int {
	return findClosestValue(node, node.value, target)
}

func findClosestValue(node *Node, closest, target int) int {
	if absDifference(target, closest) > absDifference(target, node.value) {
		closest = node.value
	}

	if node.left != nil && target < node.value {
		return findClosestValue(node.left, closest, target)
	} else if node.right != nil && target > node.value {
		return findClosestValue(node.right, closest, target)
	}

	return closest
}

func absDifference(first, second int) int {
	if first > second {
		return first - second
	}

	return second - first
}
