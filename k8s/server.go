package main

import (
  "fmt"
  "time"
  "io/ioutil"
  "log"
  "net/http"
  "os"
)

var startedAt = time.Now()

func main() {
  http.HandleFunc("/", Hello)
  http.HandleFunc("/healthz", Healthz)
  http.HandleFunc("/configmap", ConfigMap)
  http.HandleFunc("/secret", Secret)
  http.ListenAndServe(":8000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
  name := os.Getenv("NAME")
  age := os.Getenv("AGE")
  fmt.Fprintf(w, "Hello, I'm %s. I'm %s.", name, age)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
  data, err := ioutil.ReadFile("myfamily/family.txt")
  if err != nil {
    log.Fatalf("Error reading file:", err)
  }
  fmt.Fprintf(w, "My Family: %s", string(data))
}

func Secret(w http.ResponseWriter, r *http.Request) {
  user := os.Getenv("USER")
  pass := os.Getenv("PASSWORD")
  fmt.Fprintf(w, "Credentials - %s:%s", user, pass)
}

func Healthz(w http.ResponseWriter, r *http.Request) {
  duration := time.Since(startedAt)

  if duration.Seconds() < 10 || duration.Seconds() > 30 {
    w.WriteHeader(500)
    w.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
  } else {
      w.WriteHeader(200)
      w.Write([]byte("ok"))
  }
}
