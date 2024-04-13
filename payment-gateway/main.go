package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/Setsu548/Desuna-payment-platform/tree/master/bank-simulator-backend/db/sqlc"
	"github.com/Setsu548/Desuna-payment-platform/tree/master/bank-simulator-backend/db/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load the project config with error: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to Database with error: ", err)
	}

	store := sqlc.NewStore(conn)
	server, err :=

		log.Println(conn)
}
