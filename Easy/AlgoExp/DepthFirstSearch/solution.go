package DepthFirstSearch

type Node struct {
	name     string
	children []*Node
}

func DepthFirstSearch(root *Node, array []string) []string {
	c := &cache{cacheArray: array}

	recCollection(c, root)

	return c.cacheArray
}

func recCollection(c *cache, node *Node) {
	c.addName(node.name)

	for _, child := range node.children {
		recCollection(c, child)
	}
}

type cache struct {
	cacheArray []string
}

func (c *cache) addName(name string) {
	c.cacheArray = append(c.cacheArray, name)
}
