package main

func swap(a int, b *int) int {
	tmp := a

	a = *b
	*b = tmp

	return a
}
