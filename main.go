package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
	"uuid"
)

var coins, stash int

type batch struct {
	CoinsAmount    int
	UserId         int
	LastUpdatedAt  time.Time
	IdempotencyKey uuid.UUID
}

var pseudoDB = make(map[uuid.UUID]batch)
var mx = sync.Mutex{}

func main() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)

	http.HandleFunc("/batch/send", SendBatch)
	http.HandleFunc("/coin/add", AddCoin)
	http.HandleFunc("/coin/get", GetCoins)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func AddCoin(w http.ResponseWriter, req *http.Request) {
	coins++
	io.WriteString(w, "Coin accepted\n")
}

func GetCoins(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Coins in stash:"+strconv.Itoa(stash)+"\n")
}

func SendBatch(w http.ResponseWriter, req *http.Request) {
	mx.Lock()
	defer mx.Unlock()

	var b = batch{
		CoinsAmount:    coins,
		UserId:         1,
		LastUpdatedAt:  time.Now(),
		IdempotencyKey: uuid.New(),
	}

	if _, found := pseudoDB[b.IdempotencyKey]; !found {
		pseudoDB[b.IdempotencyKey] = b
		stash += b.CoinsAmount
		coins = 0
		io.WriteString(w, "Successfully added to pseudoDB\n")
	} else {
		io.WriteString(w, "Duplicated value, nothing added\n")
	}
}
