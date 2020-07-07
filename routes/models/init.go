package models

import (
	"fmt"
	"github.com/kentpon/LetsGO/utils"
)

var DB = utils.DB

func init() {
	// Create uuid function
	DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	DB.AutoMigrate(&Note{})
	fmt.Println("Done DB migration")
}
