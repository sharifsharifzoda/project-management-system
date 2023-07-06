package db

import (
	"fmt"
	"github.com/sharifsharifzoda/project-management-system/configs"
	"github.com/sharifsharifzoda/project-management-system/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func GetDBConnection(cfg configs.DatabaseConnConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Dushanbe",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DbName, cfg.SSLMode)
	//fmt.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error connecting database. Error is %v", err.Error())
		panic(err.Error())
	}

	log.Printf("Connection success host:%s port:%s", cfg.Host, cfg.Port)

	return db
}

func Close(db *gorm.DB) {
	conn, err := db.DB()
	err = conn.Close()
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func Init(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{}, &models.Project{}, &models.ProjectParticipant{}, &models.Task{})
	if err != nil {
		log.Fatal(err)
	}

	var superuser = models.User{
		Firstname: "Sharif",
		Lastname:  "Sharifzoda",
		Email:     "sharifzodastudy@gmail.com",
		Password:  "sharifzoda123",
		Role:      "superuser",
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(superuser.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("failed to hash the superuser' s password. Error is: ", err.Error())
	}

	superuser.Password = string(hash)
	if err := db.Model(&models.User{}).Create(&superuser).Error; err != nil {
		log.Println("Error is: ", err.Error())
	}
}
