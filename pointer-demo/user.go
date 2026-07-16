package main

type user struct {
	Name string
	Age  int
}

func birthday(u *user) {
	u.Age++
}
