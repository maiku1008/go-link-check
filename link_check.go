package main

import (
    "encoding/csv"
    "flag"
    "fmt"
    "log"
    "net/http"
    "os"
)

type urlStatus struct {
    url    string
    status bool
    // status_code int
    // redirect_url string
}

func FollowUrl(link []string, c chan urlStatus) {
    url := link[0]
    _, err := http.Get(url)
    if err != nil {
        c <- urlStatus{url, false}
        return
    }
    c <- urlStatus{url, true}
}

func main() {
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

    c := make(chan urlStatus)

    // Follow each link
    for _, link := range links {
        go FollowUrl(link, c)
    }

    result := make([]urlStatus, len(links))
    for i, _ := range result {
        result[i] = <-c
        if result[i].status {
            fmt.Println(result[i].url, "is up.")
        } else {
            fmt.Println(result[i].url, "is down !!")
        }
    }
}
