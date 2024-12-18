package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server...")
	http.HandleFunc("/api/test", getTest)
	http.ListenAndServe("127.0.0.1:8080", nil)

	//go func() {
	//	if err := http.ListenAndServe(":8000", nil); !errors.Is(err, http.ErrServerClosed) {
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

func getTest(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"test": "works!",
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(data)
}
