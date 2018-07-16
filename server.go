package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"io/ioutil"

	"golang.org/x/net/websocket"
)

func Echo(ws *websocket.Conn) {
	var err error

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)

		msg := "Received:  " + reply
		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}



func GetPage(ws *websocket.Conn) {
	var err error
	
	
	for {
		var reply string
	

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)
		
		u, errP := url.Parse(reply)

		if errP != nil {

		 log.Fatal(errP)
		
		}
		
		resp, err := http.Get(u.String())
		if err != nil {
			fmt.Println(err.Error())
		}
		
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		
		
		
		msg := "Received:  " + reply + "http page: " + string(body)
		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}

func main() {
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.Handle("/echo", websocket.Handler(Echo))
	http.Handle("/GetPage", websocket.Handler(GetPage))

	port := os.Getenv("PORT")
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
