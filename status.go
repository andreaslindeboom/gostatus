package main

import (
    "os"
    "fmt"
    "log"
    "net/http"
)

func main() {
    port := os.Args[1]
    http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
        fmt.Fprintf(writer, "{ \"status\": \"ok\" }")
    })

    log.Fatal(http.ListenAndServe(":" + port, nil))
}
