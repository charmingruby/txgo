package helper

func If[T any](cond bool, right T, left T) T {
	if cond {
		return right
	}

	return left
}
