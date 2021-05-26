package db

import (
	"fmt"
	"log"

	"github.com/DaisukeMatsumoto0925/sample_pj/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type SQLHandler struct {
	Conn *gorm.DB
}

func New() *SQLHandler {
	connectionString := genConnectString()
	fmt.Println("connectionString: ", connectionString)
	conn, err := gorm.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("db connection error: ", err)
	}
	fmt.Println("Connection OK")
	conn.LogMode(true) // debug mode
	sqlHandler := &SQLHandler{
		Conn: conn,
	}

	return sqlHandler
}

func genConnectString() string {
	conf := config.Get()
	USER := conf.Db.User
	PASSWORD := conf.Db.Password
	PROTOCOL := fmt.Sprintf("tcp(db:%s)", conf.Db.Port)
	DBNAME := conf.Db.Name
	QUERY := "?charset=utf8mb4&parseTime=True&loc=Local"

	return USER + ":" + PASSWORD + "@" + PROTOCOL + "/" + DBNAME + QUERY
}
