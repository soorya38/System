package main

import "profile/handlers"

func main() {
	handlers.RegisterHandlers()

	if err := handlers.StartServer(8080); err != nil {
		panic(err)
	}
}