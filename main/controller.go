package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"pixiv_api/pixiv"
	"strconv"
	"time"
)

func main() {
	refreshToken := flag.String("r", "", "Refresh Token")
	host := flag.String("h", ":9630", "Port")
	flag.Parse()
	coverFromEnv(refreshToken, host)

	client := pixiv.Client{Cxt: pixiv.NewContext("", *refreshToken)}

	client.Login()

	logger := log.New(os.Stdout, "http", log.LstdFlags)

	http.HandleFunc("/illust", func(writer http.ResponseWriter, request *http.Request) {
		logger.Printf("handling: /illust")

		pid, err := strconv.ParseInt(request.URL.Query().Get("pid"), 10, 32)
		if err != nil {
			writer.WriteHeader(400)
			return
		}
		b, _ := json.Marshal(client.Illust(int(pid)))
		writer.Header().Set("content-type", "application/json")
		_, _ = writer.Write(b)
	})

	http.HandleFunc("/related", func(writer http.ResponseWriter, request *http.Request) {
		logger.Printf("handling: /related")

		pid, err := strconv.ParseInt(request.URL.Query().Get("pid"), 10, 32)
		if err != nil {
			writer.WriteHeader(400)
			return
		}
		b, _ := json.Marshal(client.Related(int(pid), 0))
		writer.Header().Set("content-type", "application/json")
		_, _ = writer.Write(b)
	})

	http.HandleFunc("/user", func(writer http.ResponseWriter, request *http.Request) {
		logger.Printf("handling: /user")

		id, err := strconv.ParseInt(request.URL.Query().Get("id"), 10, 32)
		if err != nil {
			writer.WriteHeader(400)
			return
		}
		b, _ := json.Marshal(client.Member(int(id)))
		writer.Header().Set("content-type", "application/json")
		_, _ = writer.Write(b)
	})

	http.HandleFunc("/userIllust", func(writer http.ResponseWriter, request *http.Request) {
		logger.Printf("handling: /userIllust")

		id, err := strconv.ParseInt(request.URL.Query().Get("id"), 10, 32)
		if err != nil {
			writer.WriteHeader(400)
			return
		}
		b, _ := json.Marshal(client.MemberIllusts(int(id), 0))
		writer.Header().Set("content-type", "application/json")
		_, _ = writer.Write(b)
	})

	http.HandleFunc("/rank", func(writer http.ResponseWriter, request *http.Request) {
		logger.Printf("handling: /rank")

		mode := request.URL.Query().Get("mode")
		if mode == "" {
			writer.WriteHeader(400)
			return
		}
		yesterday := time.Now().AddDate(0, 0, -1)
		b, _ := json.Marshal(client.Rank(mode, 0, &yesterday))
		writer.Header().Set("content-type", "application/json")
		_, _ = writer.Write(b)
	})

	http.HandleFunc("/searchByTitle", func(writer http.ResponseWriter, request *http.Request) {
		logger.Printf("handling: /searchByTitle")

		title := request.URL.Query().Get("title")
		if title == "" {
			writer.WriteHeader(400)
			return
		}
		b, _ := json.Marshal(client.SearchByTitle(title))
		writer.Header().Set("content-type", "application/json")
		_, _ = writer.Write(b)
	})

	http.HandleFunc("/searchByTags", func(writer http.ResponseWriter, request *http.Request) {
		logger.Printf("handling: /searchByTags")

		tag := request.URL.Query().Get("tag")
		if tag == "" {
			writer.WriteHeader(400)
			return
		}
		b, _ := json.Marshal(client.SearchByTags(tag))
		writer.Header().Set("content-type", "application/json")
		_, _ = writer.Write(b)
	})

	logger.Println("Server listening on " + *host)

	_ = http.ListenAndServe(*host, nil)
}

func coverFromEnv(refreshToken, host *string) {
	r := os.Getenv("refresh_token")
	h := os.Getenv("host")
	if r != "" {
		*refreshToken = r
	}
	if h != "" {
		*host = h
	}
}
