package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server...")
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/api/test", getTest)
	http.ListenAndServe(":8080", nil)

	//go func() {
	//	if err := http.ListenAndServe(":8080", nil); !errors.Is(err, http.ErrServerClosed) {
	//		log.Fatalln(err)
	//	}
	//}()

	//// graceful shutdown
	//quit := make(chan os.Signal, 1)
	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//<-quit
	//log.Println("Shutdown Server ...")

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//if err := server.Shutdown(ctx); err != nil {
	//	log.Fatalln(err)
	//}
	//log.Println("Server exiting")
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	data := map[string]string{
		"hello": "world",
	}
	json.NewEncoder(w).Encode(data)
}

func getTest(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"test": "changes",
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(data)
}
