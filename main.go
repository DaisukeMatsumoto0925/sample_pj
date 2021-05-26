package main

import (
	"github.com/DaisukeMatsumoto0925/sample_pj/src/infrastructure/db"
	"github.com/DaisukeMatsumoto0925/sample_pj/src/infrastructure/logging"
	"github.com/DaisukeMatsumoto0925/sample_pj/src/infrastructure/router"
)

func main() {
	logger := logging.NewLogger("development")
	logger.Debug("debug log")

	sqlhandler := db.New()
	defer sqlhandler.Conn.Close()

	router.Dispatch(sqlhandler)
}
