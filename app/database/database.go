package database

import (
	"fmt"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)
func Connect() *gorm.DB {

  dsn := "docker:docker@tcp(mysql_host)/test_database?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		fmt.Printf("Connection with database failed: %v", err)
	} 
	return db
}
