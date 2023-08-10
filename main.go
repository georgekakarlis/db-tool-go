package main

import (
	"fmt"

	"tastetable.com/recipe-db/database"
)




func main() {

	database.ConnectDB()



	fmt.Println("Tables created successfully")
}

