package users_db

import ( 
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"github.com/joho/godotenv"

)

const (
	mysql_users_username = "mysql_users_username"
	mysql_users_password = "mysql_users_password"
	mysql_users_host = "mysql_users_host"
	mysql_users_schema = "mysql_users_schema"
)

var (
	Client *sql.DB
)

func init() {
	errG := godotenv.Load(".env")

	if errG != nil {
		log.Fatalf("Error during .env file loading")
	}

	username := os.Getenv(mysql_users_username)
	password := os.Getenv(mysql_users_password)
	host := os.Getenv(mysql_users_host)
	schema := os.Getenv(mysql_users_schema)

	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
	username, password, host, schema,
)
	var err error
	Client, err = sql.Open("mysql", datasourceName) 

	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}