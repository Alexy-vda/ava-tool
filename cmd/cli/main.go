package main

import (
	"fmt"
	"go-init-service/internal/commands"
)

func main() {
	fmt.Println(`
	 _____           _     
	|  __ \         | |    
	| |  | | ___  __| |___ 
	| |  | |/ _ \/ _' / __|
	| |__| |  __/ (_| \__ \
	|_____/ \___|\__,_|___/
	Simple Tool CLI
	`)

	// Appeler la fonction qui gère la création d'un nouveau service
	commands.CreateServiceCommand()
}