package main

import (
	"fmt"

	"github.com/AllanCapistrano/cnx-migrations/services"
)

// "github.com/AllanCapistrano/cnx-migrations/cmd"

func main() {
	// cmd.Execute()

    databases := services.GetDatabases()

	fmt.Println(databases)
}
