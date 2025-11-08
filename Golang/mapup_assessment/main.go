package main

import (
	"fmt"
	"log"
	"mongoserver/router"
	"net/http"
)

func main() {
	fmt.Println("🚀 MongoDB API")
	r := router.Router()

	fmt.Println("✅ Server is starting at http://0.0.0.0:4000")

	// Start server (blocking call)
	err := http.ListenAndServe("0.0.0.0:4000", r)
	if err != nil {
		log.Fatal("❌ Server failed:", err)
	}
}
