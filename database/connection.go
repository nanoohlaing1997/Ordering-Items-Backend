package database

import (
	"log"
	"noh/go-order-items/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

type DatabaseManger struct {
	*models.ItemDB
	*models.ItemCategoryDB
	*models.CustomerOrderDB
}

func Connect(connectionString string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Println("Cannot connect to database")
	}
	log.Println("Connected to database...")
	return db
}

// can be use as migration
func Migrate() {
	db.AutoMigrate(models.Item{})
	db.AutoMigrate(models.ItemCategory{})
	db.AutoMigrate(models.CustomerOrder{})
	log.Println("Database Migration Complete...")
}

func DBManger() *DatabaseManger {
	godotenv.Load()
	dbName := os.Getenv("GO_ORDER")
	connection := dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	return &DatabaseManger{
		models.ItemManager(Connect(connection)),
		models.ItemCategoryManager(Connect(connection)),
		models.CustomerOrderManager(Connect(connection)),
	}
}
