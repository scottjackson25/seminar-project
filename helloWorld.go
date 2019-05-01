package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/rs/cors"
)

type Password struct {
	Value string `json:"value"`
}
type Status struct {
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

func main() {
	mux := http.NewServeMux()
	file, err := os.Open("LazyPass.txt")
	if err != nil {
		fmt.Printf("%s ", err)
	}
	defer file.Close()
	var body []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		body = append(body, scanner.Text())
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var stat Status
			decoder := json.NewDecoder(r.Body)
			var p Password
			err := decoder.Decode(&p)
			if err != nil {
				fmt.Fprintf(w, "%s", err)
			}
			length := length(p.Value)
			contains := contains(p.Value)
			containsp := !containsspace(p.Value)
			lazy := !lazy(p.Value, body)
			w.WriteHeader(http.StatusOK)
			if length && contains && containsp && lazy {
				stat.Message = "Good Password!"
			} else {
				stat.Message = "Bad Password!"
				if !length {
					stat.Errors = append(stat.Errors, "Password Failed length")
				}
				if !contains {
					stat.Errors = append(stat.Errors, "Password Failed !contains")
				}
				if !containsp {
					stat.Errors = append(stat.Errors, "Password Failed contains space")
				}
				if !lazy {
					stat.Errors = append(stat.Errors, "Password Failed lazy ")

				}

			}
			if err := json.NewEncoder(w).Encode(stat); err != nil {
				log.Println(err)
			}

		}

	})
	handler := cors.Default().Handler(mux)
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(port, handler))
}

func length(pass string) bool {
	return len(pass) > 10
}

func contains(pass string) bool {
	return strings.Contains(pass, "!")
}
func containsspace(pass string) bool {
	return strings.Contains(pass, " ")

}
func lazy(pass string, str []string) bool {
	for _, word := range str {
		if word == pass {
			return true
		}

	}
	return false
}
