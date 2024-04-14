package main

import (
	"context"
	"fmt"
	"time"

	pg "github.com/go-pg/pg/v9"
	_ "github.com/go-pg/pg/v9/orm"
)

type Ins struct {
	tableName struct{} `pg:"ins,alias:t,discard_unknown_columns"`

	Timewithout time.Time `pg:"timewithout"`
	Timewith    time.Time `pg:"timewith"`
	Testcase    string    `pg:"testcase"`
}

func main() {
	conn := pg.Connect(&pg.Options{
		Addr:     "127.0.0.1:5432",
		User:     "postgres",
		Password: "Test_1tesT",
		Database: "timetest",
	})
	defer conn.Close()
	//insert

	c1 := time.Date(2024, 03, 20, 12, 00, 00, 00, time.FixedZone("", int(3*60*60)))
	c2 := time.Date(2024, 03, 20, 12, 00, 00, 00, time.FixedZone("", int(-7*60*60)))
	c3 := time.Date(2024, 03, 20, 12, 00, 00, 00, time.FixedZone("", 0))
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)

	_, err := conn.ModelContext(context.Background(), &Ins{
		Timewithout: c1,
		Timewith:    c1,
		Testcase:    "2024-03-20T12:00:00+03",
	}).Insert()
	if err != nil {
		panic(fmt.Errorf("test 1 failed: %w", err))
	}

	_, err = conn.ModelContext(context.Background(), &Ins{
		Timewithout: c2,
		Timewith:    c2,
		Testcase:    "2024-03-20T12:00:00-07",
	}).Insert()
	if err != nil {
		panic(fmt.Errorf("test 2 failed: %w", err))
	}

	_, err = conn.ModelContext(context.Background(), &Ins{
		Timewithout: c3,
		Timewith:    c3,
		Testcase:    "2024-03-20T12:00:00Z",
	}).Insert()
	if err != nil {
		panic(fmt.Errorf("test 3 failed: %w", err))
	}

	// select

	var out2 []Ins

	_, err = conn.Query(&out2, "select * from ins")
	if err != nil {
		panic(fmt.Errorf("select 2 failed: %w", err))
	}
	fmt.Println("====q====")
	for _, v := range out2 {
		fmt.Println("|", v.Timewithout, "|", v.Timewith, "|", v.Testcase, "|")
	}
	fmt.Println("====q====")
}
