package MiddleNode

import (
	. "algoexpert/models"
)

func MiddleNode(head *LinkedList[int]) *LinkedList[int] {
	var resultNode, current *LinkedList[int]
	var nodeMap = make(map[int]*LinkedList[int])
	var idx int
	current = head

	for current != nil {
		node := current
		nodeMap[idx] = node
		current = current.Next
		idx++
	}

	length := len(nodeMap)
	if length%2 == 0 {
		resultNode = nodeMap[length/2]
	} else {
		resultNode = nodeMap[(length/2)-1]
	}

	return resultNode
}
