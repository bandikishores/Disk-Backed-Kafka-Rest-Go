package main


/*
a) HealthCheck
	curl -X GET "localhost:9001/health"
b) Sleep API
	curl -X GET "localhost:9001/sleep"
c) Post to Kafka
	curl -X GET "localhost:9001/postKafka"

*/

import (
	//"encoding/json"
	// "bufio"
	"fmt"
	"log"
	// "os"
	"time"
	// "strconv"
	"./dependency"
	"net/http"
	"runtime/debug"
)

var err error

func main() {
	fmt.Println("hello world")
	var var1 = 1
	var var2 = "2"
	const var3 = "asdf"

	fmt.Println(var1, var2)
	// scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i <= 3; i++ {
		log.Print(i)
	}

	// Kafka
	dependency.InitKafka()
	dependency.InitDiskQueue()

	/*	for true {
		fmt.Print("Enter text: ")
		scanner.Scan()
		text := scanner.Text()
		i2, _ := strconv.ParseInt(text, 10, 64)

		if i2 < 0 {
			break;
		}

		for i := 0; i <= int(i2); i++ {
			sendMessage("1233", fmt.Sprintf("%s : %d", "test-kafka-go", i))
			// time.Sleep(3 * time.Second)
		}
	} */

	// Http Server
	router := dependency.AddRouter()
	go freeMemory(time.Second * 5)
	// http.ListenAndServe(":9001", router))
	server := &http.Server{Addr: ":9001", Handler: router, IdleTimeout: time.Second * 1}
	log.Fatal(server.ListenAndServe())
	// time.Sleep(3 * time.Hour)

}

func freeMemory(d time.Duration) {
	ticker := time.NewTicker(d)
	var gcStats debug.GCStats
	for {
		select {
		case <-ticker.C:
			debug.ReadGCStats(&gcStats)
			log.Printf("GC Stats: %v", gcStats)
			debug.FreeOSMemory()
		}
	}
}
