package main

import (
	"fmt"
	"log"
	userService "user-service/services"

	userRepository "user-service/adapters/database/memory/user"
)

func main() {
	var (
		database = userRepository.NewUserMemoryDatabase()
		service = userService.NewUserService(database)
	)

	user, err := service.Create("Rodolfo do Nascimento Azevedo", "rof20004@gmail.com", "123")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Saved:", user)
}
