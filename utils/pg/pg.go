package pg

import (
	"fmt"
	"os"
	"urlshortener/repositories/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

type Config struct {
	Host string
	User string
	Db   string
	Pwd  string
	Port string
}

func init() {
	if os.Getenv("ENV") != "test" {
		if err := godotenv.Load(); err != nil {
			fmt.Println("Error loading .env file")
		}
		PGConfig := Config{
			Host: os.Getenv("POSTGRES_HOST"),
			User: os.Getenv("POSTGRES_USER"),
			Db:   os.Getenv("POSTGRES_DB"),
			Pwd:  os.Getenv("POSTGRES_PASSWORD"),
			Port: os.Getenv("POSTGRES_PORT"),
		}
		if _, err := NewPgClient(PGConfig); err != nil {
			fmt.Println("db err", err)
		}
		Db.AutoMigrate(&model.Urls{})
	}
}

func NewPgClient(dbConfig Config) (*gorm.DB, error) {
	var err error
	Db, err = gorm.Open(postgres.Open(GetPgDns(dbConfig)), &gorm.Config{})
	if err != nil || Db.Error != nil {
		panic("database error")
	}
	return Db, err
}

func GetPgDns(conf Config) (dsn string) {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%v sslmode=disable TimeZone=Asia/Taipei",
		conf.Host, conf.User, conf.Pwd, conf.Db, conf.Port)
}
