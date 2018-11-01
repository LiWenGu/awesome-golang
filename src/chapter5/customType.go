package chapter5

type user struct {
	name  string
	email string
}
type admin struct {
	person user
	level  string
}

func main() {
	list := user{
		name:  "list",
		email: "list@gmail.com",
	}

	list.email = "xxx"

	fred := admin{
		person: user{
			name:  "fred",
			email: "fred@gmail.com",
		},
		level: "2",
	}
	fred.level = "3"
}
