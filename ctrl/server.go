package ctrl

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/r3labs/sse/v2"
	"github.com/rs/cors"
)

var informUIChan = make(chan []byte, 1)

type payload struct {
	Name  string `json:"name"`
	Value any    `json:"value"`
	CType string `json:"cType"`
}

type basePayload struct {
	Name  string
	CType string
}

func set(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(405)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(400)
		return
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	var bp basePayload
	err = json.Unmarshal(b, &bp)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	switch bp.CType {
	case "float":
		// for some reason `v, ok := p.Value.(float32)` didn't work here
		var p struct{ Value float32 }
		err = json.Unmarshal(b, &p)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		setFloatInternal(bp.Name, p.Value)
	case "int":
		var p struct{ Value int32 }
		err = json.Unmarshal(b, &p)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		setIntInternal(bp.Name, p.Value)
	case "string":
		var p struct{ Value string }
		err = json.Unmarshal(b, &p)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		setStringInternal(bp.Name, p.Value)
	case "bool":
		var p struct{ Value bool }
		err = json.Unmarshal(b, &p)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		setBoolInternal(bp.Name, p.Value)
	}
}

func RunServer() {
	server := sse.New()
	server.CreateStream("messages")

	server.Publish("messages", &sse.Event{Data: []byte("hello")}) // if I don't publish anything up front, SSE client on the browser will immediately disconnect
	go func() {
		for {
			b := <-informUIChan
			server.Publish("messages", &sse.Event{Data: b})
		}
	}()

	http.HandleFunc("/events", server.ServeHTTP)
	http.HandleFunc("/set", set)

	c := cors.Default()
	handler := c.Handler(http.DefaultServeMux)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
