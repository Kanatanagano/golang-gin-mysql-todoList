package main
import(
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Todo struct {
	*gorm.Model
	Context string `json:"content"`
}

type DBConfig struct {
	User string
	Password string
	Host string
	Port int
	Table string
}

func getDBConfig() DBConfig {
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	return DBConfig{
			User: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Host: os.Getenv("DB_HOST"),
			Port: port,
	Table: os.Getenv("DB"),
	}
}

func connectionDB() (*gorm.DB, error) {
	config := getDBConfig()
	// PostgreSQL用のDSN形式に変更
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Tokyo", 
		config.Host, config.User, config.Password, config.Table, config.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}


func main() {
	r := gin.Default()
	db, err := connectionDB()

	if err != nil {
		log.Fatalf("Failed to connect to databases: %v", err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&Todo{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	fmt.Println("Database connection and setup successful")
	r.Run(":8080")
	r.Run()
}