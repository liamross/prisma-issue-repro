# Prisma `v0.12.0` reproduction

1. run `go run github.com/prisma/prisma-client-go generate`
1. check `main.go`, `In` will not exist for `db.TestValueTwo.ID.In(ids)`

## Why?

The presence of `anyOtherArray String[]` in the `TestValueOne` model seems to be the cause. Not sure why, but that's what I've found!
