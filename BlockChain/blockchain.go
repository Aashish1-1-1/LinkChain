package BlockChain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Somerandstruct struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Score    int32  `json:"score"`
}
type Block struct {
	data         map[string]interface{}
	hash         string
	previousHash string
	timestamp    time.Time
	pow          int
}

type Blockchain struct {
	genesis    Block
	chain      []Block
	difficulty int
}
type Jsondata struct {
	Title string `json:title`
	Link  string `json:link`
}

func (b Block) calculateHash() string {
	data, _ := json.Marshal(b.data)
	blockData := b.previousHash + string(data) + b.timestamp.String() + strconv.Itoa(b.pow)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

func (b *Block) mine(difficulty int) {
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
		b.pow++
		b.hash = b.calculateHash()
	}
}

func CreateBlockchain(difficulty int) Blockchain {
	genesisBlock := Block{
		hash:      "0",
		timestamp: time.Now(),
	}
	return Blockchain{
		genesisBlock,
		[]Block{genesisBlock},
		difficulty,
	}
}

func (b *Blockchain) AddBlock(Title string, Link string) {
	blockData := map[string]interface{}{
		"Title": Title,
		"Link":  Link,
	}
	lastBlock := b.chain[len(b.chain)-1]
	newBlock := Block{
		data:         blockData,
		previousHash: lastBlock.hash,
		timestamp:    time.Now(),
	}
	newBlock.mine(b.difficulty)
	b.chain = append(b.chain, newBlock)
}

func (b Blockchain) IsValid() bool {
	for i := range b.chain[1:] {
		previousBlock := b.chain[i]
		currentBlock := b.chain[i+1]
		if currentBlock.hash != currentBlock.calculateHash() || currentBlock.previousHash != previousBlock.hash {
			return false
		}
	}
	return true
}
func (b Blockchain) GetBlockchainData() []Jsondata {
	var jsonData []Jsondata
	for _, block := range b.chain {
		Title := ""
		link := ""
		if TitleVal, ok := block.data["Title"]; ok {
			Title = TitleVal.(string)
		}
		if linkVal, ok := block.data["Link"]; ok {
			link = linkVal.(string)
		}
		jsonData = append(jsonData, Jsondata{Title: Title, Link: link})
	}
	return jsonData
}
func randStr(n int) string {

	var charset = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func MakeBlockChain() {
	blockChain := CreateBlockchain(3)
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
