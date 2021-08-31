package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/ozonva/ova-course-api/internal/repo"
	server "github.com/ozonva/ova-course-api/internal/server"
	"github.com/ozonva/ova-course-api/internal/utils"
	api "github.com/ozonva/ova-course-api/pkg/ova-course-api"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

const (
	configPath = "configs/config.yml"
	grpcPort   = ":82"
)

func main() {
	fmt.Println("Welcome to ova-template-api")

	func() {
		fmt.Println("Задание 1 - Разделение слайса на батчи")
		origin := []string{"a", "b", "c", "d", "e"}
		converted := utils.SeparationSlice(origin, 2)
		fmt.Printf("Входной срез: %v\n", origin)
		fmt.Printf("Выходной срез: %v\n", converted)
	}()

	func() {
		fmt.Println("Задание 2 - Обратный ключ ")
		origin := map[int]string{1: "a", 2: "b", 3: "c", 4: "b"}
		revert, err := utils.ReverseKeyView(origin)
		if err != nil {
			fmt.Printf("Ошибка: %s\n", err.Error())
			return
		}
		fmt.Printf("Входной Map: %v\n", origin)
		fmt.Printf("Выходной Map: %v\n", revert)

	}()

	func() {
		fmt.Println("Задание 3 - Фильтрация по захардкоженному списку")
		origin := []string{"map", "rebus", "car", "globus"}
		filtered := utils.FilterSlice(origin)
		fmt.Printf("Входной срез: %v\n", origin)
		fmt.Printf("Выходной срез: %v\n", filtered)
	}()

	// Задание 3 A Используя defer и функтор реализовать открытие и закрытие файла в цикле
	readConfig := func(filename string) (err error) {
		r, err := os.Open(filename)
		if err != nil {
			return
		}
		defer func(r *os.File) {
			err = r.Close()
		}(r)

		// Здесь работа с файлом по необходимости

		return
	}

	for i := 0; i < 100; i++ {
		err := readConfig(configPath)
		if err != nil {
			fmt.Printf("Error: %s ", err.Error())
			break
		}
	}

	if err := godotenv.Load(); err != nil {
		panic(fmt.Sprintf("error loading .env: %s", err.Error()))
	}

	//BD
	db, err := repo.NewDB(repo.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		DBName:   os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  "disabled",
	})

	if err != nil {
		panic(fmt.Sprintf("error init db: %s", err.Error()))
	}

	// сервер gRPC
	s := grpc.NewServer()
	api.RegisterCourseServer(s, server.NewCourseServer(db))

	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
