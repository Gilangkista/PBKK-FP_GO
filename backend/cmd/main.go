package main

import (
	"FP_GO_PBKK-D/internal/controller"
	"FP_GO_PBKK-D/internal/infrastructure"
	"FP_GO_PBKK-D/internal/repositories"
	"FP_GO_PBKK-D/internal/routes"
	"FP_GO_PBKK-D/internal/usecases"

	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Koneksi ke database MySQL
	dsn := "root:@tcp(127.0.0.1:3306)/?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	dbName := "fpgolang"
	// Drop database if exist
	err = db.Exec("DROP DATABASE IF EXISTS " + dbName).Error
	if err != nil {
		log.Fatalf("Failed to drop database: %v", err)
	}
	log.Printf("Database %s dropped successfully.", dbName)

	// Create new database
	err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName).Error
	if err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}
	log.Printf("Database %s created successfully.", dbName)

	dsn = "root:@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Jalankan migrasi
	infrastructure.MigrateDB(db)

	// Inisialisasi repositories
	artistRepo := &repositories.ArtistRepository{DB: db}
	songRepo := &repositories.SongRepository{DB: db}

	// Inisialisasi usecases
	artistUsecase := &usecases.ArtistUsecase{Repo: artistRepo}
	songUsecase := &usecases.SongUsecase{Repo: songRepo}

	// Inisialisasi controllers
	artistController := &controller.ArtistController{Usecase: artistUsecase}
	songController := &controller.SongController{Usecase: songUsecase}

	// Setup Gin Router
	router := gin.Default()

	// Menambahkan middleware CORS untuk mengizinkan permintaan dari frontend
	router.Use(cors.Default()) // Default CORS policy

	// Register routes
	routes.ArtistRoutes(router, artistController)
	routes.SongRoutes(router, songController)
	// Tambahkan routes lain di sini jika ada

	// Jalankan server
	port := ":8080"
	log.Printf("Server is running on http://localhost%s", port)
	if err := router.Run(port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
