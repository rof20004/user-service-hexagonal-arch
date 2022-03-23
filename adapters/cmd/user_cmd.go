package main

import (
	"fmt"
	"log"
	userRepository "user-service/adapters/database/memory/user"
	userService "user-service/application/services"
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
