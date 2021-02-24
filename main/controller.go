package main

import (
	"encoding/json"
	"flag"
	"net/http"
	"pixiv_api/pixiv"
	"strconv"
	"time"
)

func main() {
	token := flag.String("a", "", "Access Token")
	refreshToken := flag.String("r", "", "Refresh Token")
	host := flag.String("h", ":9630", "Port")
	proxy := flag.String("p", "", "Proxy")

	flag.Parse()

	client := pixiv.Client{Cxt: pixiv.NewContext(*token, *refreshToken)}

	if *proxy != "" {
		client.Cxt.Proxy = *proxy
	}

	client.Login()

	http.HandleFunc("/illust", func(writer http.ResponseWriter, request *http.Request) {
		pid, err := strconv.ParseInt(request.URL.Query().Get("pid"), 10, 32)
		if err != nil {
			writer.WriteHeader(400)
			return
		}
		b, _ := json.Marshal(client.Illust(int(pid)))
		writer.Header().Set("Context-Type", "application/json")
		_, _ = writer.Write(b)
	})

	http.HandleFunc("/related", func(writer http.ResponseWriter, request *http.Request) {
		pid, err := strconv.ParseInt(request.URL.Query().Get("pid"), 10, 32)
		if err != nil {
			writer.WriteHeader(400)
			return
		}
		b, _ := json.Marshal(client.Related(int(pid), 0))
		writer.Header().Set("Context-Type", "application/json")
		_, _ = writer.Write(b)
	})

	http.HandleFunc("/user", func(writer http.ResponseWriter, request *http.Request) {
		id, err := strconv.ParseInt(request.URL.Query().Get("id"), 10, 32)
		if err != nil {
			writer.WriteHeader(400)
			return
		}
		b, _ := json.Marshal(client.Member(int(id)))
		writer.Header().Set("Context-Type", "application/json")
		_, _ = writer.Write(b)
	})

	http.HandleFunc("/userIllust", func(writer http.ResponseWriter, request *http.Request) {
		id, err := strconv.ParseInt(request.URL.Query().Get("id"), 10, 32)
		if err != nil {
			writer.WriteHeader(400)
			return
		}
		b, _ := json.Marshal(client.MemberIllusts(int(id), 0))
		writer.Header().Set("Context-Type", "application/json")
		_, _ = writer.Write(b)
	})

	http.HandleFunc("/rank", func(writer http.ResponseWriter, request *http.Request) {
		mode := request.URL.Query().Get("mode")
		if mode == "" {
			writer.WriteHeader(400)
			return
		}
		yesterday := time.Now().AddDate(0, 0, -1)
		b, _ := json.Marshal(client.Rank(mode, 0, &yesterday))
		writer.Header().Set("Context-Type", "application/json")
		_, _ = writer.Write(b)
	})

	http.HandleFunc("/searchByTitle", func(writer http.ResponseWriter, request *http.Request) {
		title := request.URL.Query().Get("title")
		if title == "" {
			writer.WriteHeader(400)
			return
		}
		b, _ := json.Marshal(client.SearchByTitle(title))
		writer.Header().Set("Context-Type", "application/json")
		_, _ = writer.Write(b)
	})

	http.HandleFunc("/searchByTags", func(writer http.ResponseWriter, request *http.Request) {
		tag := request.URL.Query().Get("tag")
		if tag == "" {
			writer.WriteHeader(400)
			return
		}
		b, _ := json.Marshal(client.SearchByTags(tag))
		writer.Header().Set("Context-Type", "application/json")
		_, _ = writer.Write(b)
	})

	_ = http.ListenAndServe(*host, nil)
}
