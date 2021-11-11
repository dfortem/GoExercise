package main

import (
  "fmt"
  "monitoring"
)

func main() {
  status, err := monitoring.GetStatus("http://localhost:9200")
  if err != nil {
    fmt.Println("AN ERROR OCCURED! Please check logs for more details.")
  } else {
    fmt.Println("Cluster's status: " + status)
  }
}
