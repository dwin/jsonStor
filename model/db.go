package model

import (
	"crypto/tls"
	"log"
	"time"

	"github.com/dwin/jsonStor/config"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

var (
	db *pg.DB
)

func init() {

}

// PostgresConnect opens db connection and creates schema
func PostgresConnect() {
	cert, err := tls.LoadX509KeyPair(config.Viper.GetString("postgres.cert"), config.Viper.GetString("postgres.key"))
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}
	opts := pg.Options{
		MaxRetries:         2,
		IdleCheckFrequency: time.Second * 15,
		Addr:               config.Viper.GetString("postgres.address"),
		User:               config.Viper.GetString("postgres.user"),
		Password:           config.Viper.GetString("postgres.password"),
		Database:           config.Viper.GetString("postgres.database"),
		DialTimeout:        time.Second * 30,

		TLSConfig: &tls.Config{
			Certificates:       []tls.Certificate{cert},
			InsecureSkipVerify: true,
		},
	}
	db = pg.Connect(&opts)

	// Automatically create tables from models
	if err := createSchema(db); err != nil {
		panic(err)
	}

	// Log Queries
	/*
		db.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
			query, err := event.FormattedQuery()
			if err != nil {
				panic(err)
			}

			fmt.Printf("%s %s", time.Since(event.StartTime), query)
		})
	*/
}

// PostgresClose closes database connection
func PostgresClose() {
	db.Close()
}
func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{&User{}} {
		err := db.CreateTable(model, &orm.CreateTableOptions{IfNotExists: true})
		if err != nil {
			log.Fatalf("Unable to create table for model: %T error: %s", model, err)
			return err
		}
	}
	return nil
}
