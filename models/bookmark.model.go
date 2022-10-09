package models

import (
	"context"
	"crypto/rand"
	"strconv"
	"math/big"

	"encore.dev/storage/sqldb"
)

type Bookmarks struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Info      string  `json:"info"`
	Owner     int
}

func Concat(id int, c context.Context) int {
	concStr := strconv.Itoa(getRowsCnt(id, c)) + "0" + strconv.FormatInt(getRandId(), 10)
	println("concStr:", concStr)
	resStr, _ := strconv.Atoi(concStr)
	println("resStr:", resStr)
	return resStr
}

func getRowsCnt(id int, c context.Context) int {
	var (
		// ctx context.Context
		cnt int
	)

	e := sqldb.QueryRow(c, `select count(*) from Bookmark where "owner" = $1`, id).Scan(&cnt)

	if e != nil {
		panic(e)
	} else {
		println("cnt:", cnt)
		return cnt+1
	}
}

func getRandId() int64 {
	// min := 1000
	// max := 9000
	rndBigInt, e := rand.Int(rand.Reader, big.NewInt(9000))
	if e != nil {
		panic(e)
	} else {
		rndInt := rndBigInt.Int64()
		println("rndInt:", rndInt)
		return rndInt
	}
}