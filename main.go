package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"Linkchain/BlockChain"
)

var blockChain = BlockChain.CreateBlockchain(3)

func CORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		if r.Method == "OPTIONS" {
			http.Error(w, "No Content", http.StatusNoContent)
			return
		}

		next(w, r)
	}
}
func main() {
	http.HandleFunc("/", CORS(homePage))
	http.HandleFunc("/add", CORS(Addlink))
	http.HandleFunc("/aashish", CORS(ThrowLinks))
	fmt.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Error starting server: " + err.Error())
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func Addlink(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	var data BlockChain.Jsondata
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	_, err = http.Get(data.Link)
	if err != nil {
		http.Error(w, "Enter valid link you mf", http.StatusBadRequest)
		return
	}
	blockChain.AddBlock(data.Title, data.Link)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message": "Success"}`)
}
func ThrowLinks(w http.ResponseWriter, r *http.Request) {
	if blockChain.IsValid() {
		jsondata, err := json.Marshal(blockChain.GetBlockchainData())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsondata)
		return
	}
	http.Error(w, "Ur mom", http.StatusInternalServerError)
	return
}

func MakeBlockChain() {
	blockChain := BlockChain.CreateBlockchain(3)
	blockChain.AddBlock("Github", "https://github.com/")
	blockChain.AddBlock("Linkedin", "https://www.linkedin.com/in")
	blockChain.AddBlock("Instagram", "https://instagram.com/")
	blockChain.AddBlock("Twitter", "https://x.com/")
	for {
		data := blockChain.GetBlockchainData()
		if blockChain.IsValid() {
			fmt.Println(data)
		}
	}
}
