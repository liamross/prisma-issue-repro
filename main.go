package main

import (
	"context"

	"github.com/liamross/prisma-issue-repro/db"
)

func main() {
	var (
		database = db.NewClient()
		ids      = []string{}
	)

	database.TestValueTwo.FindMany(

		// This is the line that breaks: `In` does not exist, due to the
		// other array value in the relation `TestValueOne` (the other
		// array item is `anyOtherArray`). If you remove `anyOtherArray`
		// from `prisma/schema.prisma` and re-run the generate command,
		// this will be fixed.
		db.TestValueTwo.ID.In(ids),
		//                 ^^ db.TestValueTwo.ID.In undefined (type
		//                    db.testValueTwoQueryIDString has no field
		//                    or method In)
		//                    ------------------------------------------
		//                    compiler - MissingFieldOrMethod
	).Exec(context.Background())
}
