package FinalValueOfVariableAfterPerformingOperations_2011

func v1(operations []string) int {
	var x int

	for _, oper := range operations {
		if oper[1] == '+' {
			x++
		} else {
			x--
		}
	}

	return x
}
