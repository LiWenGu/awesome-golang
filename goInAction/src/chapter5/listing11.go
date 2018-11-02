package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("Sending User email To %s<%s>\n",
		u.name,
		u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	bill := user{"Bill", "bill@gmail.com"}
	bill.notify()

	list := user{"lisa", "list@gmail.com"}
	list.notify()

	bill.changeEmail("bill@163.com")
	bill.notify()

	list.changeEmail("lisa@163.com")
	list.notify()

	lisa := &user{"lisa", "lisaaaa"}
	lisa.notify()
}
