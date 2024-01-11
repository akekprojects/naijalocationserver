package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/muhammadolammi/naijalocationserver/internal/database"
)

type Config struct {
	PORT string
	DB   *database.Queries
}

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Println(err)
	// }
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	log.Println("there is no port provided kindly provide a port.")
	// 	return
	// }

	// dbURL := os.Getenv("DB_URL")
	// if dbURL == "" {
	// 	log.Println("empty dbURL")
	// 	return
	// }
	port := "8080"
	dbURL := "postgresql://postgres:akek@localhost:5432/naijalocation?sslmode=disable"
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	dbQueries := database.New(db)

	apiConfig := Config{
		PORT: port,
		DB:   dbQueries,
	}

	server(&apiConfig)

}
