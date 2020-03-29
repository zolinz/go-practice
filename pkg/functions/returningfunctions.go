package functions





func MyReturningFunction() func() int {
	return func() int{
		return 45
	}
}
