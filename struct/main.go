package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

func displayUser(user User) string {
	return fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
}

func displayGroup(group Group) {
	fmt.Printf("Name : %s", group.Name)
	fmt.Println("")
	fmt.Printf("Jumlah member : %d", len(group.Users))
	fmt.Println("")
	fmt.Println("Nama User :")
	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}
}

// METHOD
func (user User) Display() string {
	return fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
}

func main() {
	user := User{}
	user.ID = 1
	user.FirstName = "Iqbal"
	user.LastName = "Nugraha"
	user.Email = "iqbal@gmail.com"
	user.IsActive = true

	// user2 := User{
	// 	ID:        2,
	// 	FirstName: "Dwi",
	// 	LastName:  "Shavera",
	// 	Email:     "dwitia@gmail.com",
	// 	IsActive:  true,
	// }

	user3 := User{
		ID:        3,
		FirstName: "Method",
		LastName:  "Panggil",
		Email:     "method@gmail.com",
		IsActive:  true,
	}

	// users := []User{user, user2}

	// group1 := Group{"Gamer", user, users, true}

	// displayGroup(group1)

	// fmt.Println(user)
	// fmt.Println(user2)

	// displayUser1 := displayUser(user)
	// displayUser2 := displayUser(user2)

	// fmt.Println(displayUser1)
	// fmt.Println(displayUser2)

	// panggil method
	result := user3.Display()
	fmt.Println(result)
}
