package main

import (
	"log"
	"time"

	"github.com/PumpkinSeed/sqlfuzz/drivers"
	"github.com/PumpkinSeed/sqlfuzz/pkg/connector"
	"github.com/PumpkinSeed/sqlfuzz/pkg/descriptor"
	"github.com/PumpkinSeed/sqlfuzz/pkg/flags"
	"github.com/PumpkinSeed/sqlfuzz/pkg/fuzzer"
	"github.com/brianvoe/gofakeit/v5"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	f := flags.Get()
	gofakeit.Seed(0)
	db := connector.Connection(drivers.New(flags.Get().Driver))
	defer db.Close()

	var tables []string
	if f.Table == "" {
		var err error
		tables, err = descriptor.ShowTables(db)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		tables = []string{f.Table}
	}
	for _, table := range tables {
		f.Table = table
		fields, err := descriptor.Describe(db, table)
		if err != nil {
			log.Fatal(err.Error())
		}

		t := time.Now()
		if err := fuzzer.Run(db, fields, f); err != nil {
			log.Fatal(err.Error())
		}
		log.Printf("Fuzzing %s table taken: %v \n", table, time.Since(t))
	}
}
