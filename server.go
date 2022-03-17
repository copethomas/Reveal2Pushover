package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RevealWebhook struct {
	r *http.ServeMux
}

func (s *RevealWebhook) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	Info.Printf("Request: %s", r.URL.String())
	s.r.ServeHTTP(w, r)
}

func RunServer() *http.Server {
	r := http.NewServeMux()
	r.HandleFunc("/reveal2pushover", webhookHandler)
	http.Handle("/", &RevealWebhook{r})
	server := &http.Server{Addr: conf.ServerAddr}
	go func() {
		Info.Printf("Starting WebServer, Listening on %s ...", conf.ServerAddr)
		Info.Printf("Online and Ready")
		if err := server.ListenAndServe(); err != nil {
			// cannot panic, because this probably is an intentional close
			Error.Printf("Webserver Error: %s", err)
		}
	}()
	return server
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	var clientEvent RevealWebHook
	bodyData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Error.Printf("Failed to read Message Body. Reason: %s", err.Error())
		http.Error(w, "internal server error", 500)
		return
	}
	if err = json.Unmarshal(bodyData, &clientEvent); err != nil {
		Error.Printf("Bad Data. Reason: %s", err.Error())
		http.Error(w, "internal server error", 500)
		return
	}
	err = SendPushNotification(fmt.Sprintf("New Website Visitor! - %s ", "test"))
	if err != nil {
		Error.Printf("failed to send push notification, err = %s", err.Error())
		http.Error(w, "internal server error", 500)
		return
	}
	w.Write([]byte("ok"))
}
