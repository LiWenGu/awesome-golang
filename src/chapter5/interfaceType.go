package main

import "fmt"

type notifier interface {
	notifyz()
}

type user struct {
	name  string
	email string
}

func (u *user) notifyz() {
	fmt.Printf("Sending%s<%s>", u.name, u.email)
}

func main() {
	u := user{"bill", "bill@email.com"}
	sendNotification(&u)
}

func sendNotification(n notifier) {
	n.notifyz()
}
