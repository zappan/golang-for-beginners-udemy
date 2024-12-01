package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"myapp/game"
	"strconv"
)

const PORT = 8080


func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}


func playRound(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := game.PlayRound(playerChoice)

	out, err := json.MarshalIndent(result, "", "  ")
	if (err != nil) {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/hson")
	w.Write(out)

	log.Println(result.Winner, result.ComputerChoice, result.RoundResult, result.WisdomNugget)
}


func renderTemplate(w http.ResponseWriter, pathToPage string) {
	tpl, err := template.ParseFiles(pathToPage)
	if err != nil {
		log.Println(err)
		return
	}

	err = tpl.Execute(w, nil)   // nil: a viariable used to pass data to the template
	if err != nil {
		log.Println(err)
		return
	}
}


func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/play", playRound)

	log.Println("Starting web server on port", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}
