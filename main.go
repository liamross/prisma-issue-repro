package main

import (
	"context"
	"fmt"
	"log"

	"github.com/liamross/prisma-issue-repro/prisma/db"
)

func main() {
	database := db.NewClient()
	if err := database.Prisma.Connect(); err != nil {
		log.Fatalf("error connecting to database")
	}

	defer func() {
		if err := database.Prisma.Disconnect(); err != nil {
			log.Fatalf("error disconnecting from database")
		}
	}()

	tv1, err := database.TestValueTwo.CreateOne(
		db.TestValueTwo.TestValueOne.Link(db.TestValueOne.ID.Equals("")),
	).Exec(context.Background())
	if err != nil {
		log.Fatalf("unable to create CustomerSession 1")
	}

	tv2, err := database.TestValueTwo.CreateOne(
		db.TestValueTwo.TestValueOne.Link(db.TestValueOne.ID.Equals("")),
	).Exec(context.Background())
	if err != nil {
		log.Fatalf("unable to create CustomerSession 2")
	}

	ids := []string{tv1.ID, tv2.ID}
	sessions, err := database.TestValueTwo.FindMany(

		// This is the line that breaks: `In` does not exist.
		db.TestValueTwo.ID.In(ids),
	).Exec(context.Background())
	if err != nil {
		log.Fatalf("unable to find many CustomerSession")
	}

	fmt.Println(sessions)
}
