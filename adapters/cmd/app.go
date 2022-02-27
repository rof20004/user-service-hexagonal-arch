package main

import (
	"fmt"
	"log"

	userRepository "user-service/adapters/database/memory"
	userApp "user-service/application/user"
	userPort "user-service/ports/user"
)

func main() {
	var (
		userDb = userRepository.NewUserMemoryDatabase()
		userSv = userPort.NewUserService(userDb)
	)

	var dto = userApp.CreateUserDto{
		Name:     "Rodolfo",
		Email:    "rof20004@gmail.com",
		Password: "123",
	}

	view, err := userSv.Create(dto)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Saved:", view)

	v, err := userSv.GetById("1")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Get by id:", v)
}
