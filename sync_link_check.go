package main

import (
    "encoding/csv"
    "flag"
    "fmt"
    "log"
    "net/http"
    "os"
    "time"
)

func FollowUrl(link []string) {
    url := link[0]
    _, err := http.Get(url)
    // fmt.Println(resp)
    // fmt.Println(resp.StatusCode)
    if err != nil {
        fmt.Println(url + " is down!!")
        return
    }
    fmt.Println(url + " is up.")
}

func main() {
    start := time.Now()

    // Pass CSV file with -csv option. Defaults to urls.csv
    file := flag.String("csv", "urls.csv", "csv file name")
    flag.Parse()

    // Open File
    f, err := os.Open(*file)
    if err != nil {
        log.Fatal(err)
    }

    // Read CSV
    links, err := csv.NewReader(f).ReadAll()
    if err != nil {
        log.Fatal(err)
    }

    // Follow each link
    for _, link := range links {
        FollowUrl(link)
    }

    elapsed := time.Since(start)
    fmt.Println("Sync link checking for", len(links), "links took", elapsed.Seconds(), "seconds")
}
