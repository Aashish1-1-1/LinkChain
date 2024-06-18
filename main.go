package main

import(
	"fmt"
	"net/http"
	"encoding/json"
	"html/template"

	"Linkchain/BlockChain"
)

var throw = make(chan int)
var jsondata = make(chan []BlockChain.Jsondata)

func main(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/aashish",ThrowLinks)
	go MakeBlockChain()
	
    	fmt.Println("Server started on :8080")
    	err := http.ListenAndServe(":8080", nil)
    	if err != nil {
       		 panic("Error starting server: " + err.Error())
   	}
}

func MakeBlockChain(){
	blockChain := BlockChain.CreateBlockchain(3)
	blockChain.AddBlock("Github","https://github.com/Aashish1-1-1")
	blockChain.AddBlock("Linkedin","https://www.linkedin.com/in/aashish-adhikari-92a958212/")
	blockChain.AddBlock("Instagram","https://instagram.com/aashish_em")
	blockChain.AddBlock("Twitter","https://x.com/AashishAdh9")
	for{
		<-throw
		data := blockChain.GetBlockchainData()
		fmt.Println(data)
	 	if(blockChain.IsValid()){
			jsondata <- data
		}
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl,err:= template.ParseFiles("index.html")
	if err!=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w,nil)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
}

func ThrowLinks(w http.ResponseWriter, r *http.Request){
	throw <- 1
	data := <- jsondata
	jsonData,err := json.Marshal(data)
	if err != nil {
       		 http.Error(w, err.Error(), http.StatusInternalServerError)
        	return
   	 }
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
