package fate

const (
	opContains    = "contains"
	opNotContains = "!contains"
)

func compareSet(leftValue interface{}, rightValue interface{}, op string) bool {
	switch left := leftValue.(type) {
	case []interface{}:
		return sliceContains(left, rightValue, op)
	default:
		panic(errorOperationMsg(op, left, rightValue))
	}
}

func sliceContains(left []interface{}, rightValue interface{}, op string) bool {
	switch right := rightValue.(type) {
	case []interface{}:
		if op == opContains {
			return sliceContainsSlice(left, right)
		}
		return !sliceContainsSlice(left, right)
	default:
		if op == opContains {
			return sliceContainsInterface(left, right)
		}
		return !sliceContainsInterface(left, right)
	}
}

func sliceContainsSlice(left []interface{}, right []interface{}) bool {
	if len(left) < len(right) {
		return false
	}
	for _, rightValue := range right {
		if !sliceContainsInterface(left, rightValue) {
			return false
		}
	}
	return true
}

func sliceContainsInterface(left []interface{}, right interface{}) bool {
	for _, leftValue := range left {
		if leftValue == right {
			return true
		}
	}
	return false
}
