package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

var wg sync.WaitGroup

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	num, _ := strconv.Atoi(os.Getenv("NUM_OF_REQUESTS"))

	wg.Add(num)

	start := time.Now()
	for i := 0; i < num; i++ {
		go test()
	}

	wg.Wait()

	fmt.Println(time.Now().Sub(start))
}

func test() {
	http.Get(os.Getenv("WEBSITE_URL"))
	wg.Done()
}
