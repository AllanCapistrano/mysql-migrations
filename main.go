package main

import (
	"fmt"

	"github.com/AllanCapistrano/cnx-migrations/services"
)

// "github.com/AllanCapistrano/cnx-migrations/cmd"

func main() {
	// cmd.Execute()

	fmt.Println(services.GetDatabases())
}
