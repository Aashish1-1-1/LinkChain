package BlockChain

import(
	"fmt"
	"crypto/sha256"
	"strconv"
	"strings"
	"time"
	"encoding/json"
)

type Block struct{
	data map[string]interface{}
	hash  string
	previousHash string
	timestamp time.Time
	pow int
}

type Blockchain struct{
	genesis Block
	chain []Block
	difficulty int
}
type Jsondata struct{
	Platform string `json:platform`
	Link interface{} `json:link`
}

func (b Block) calculateHash()string{
	data, _:=json.Marshal(b.data)
	blockData := b.previousHash + string(data) + b.timestamp.String() + strconv.Itoa(b.pow)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

func (b *Block) mine(difficulty int){
	for !strings.HasPrefix(b.hash, strings.Repeat("0",difficulty)){
		b.pow++
		b.hash=b.calculateHash()
	}
}

func CreateBlockchain(difficulty int) Blockchain{
	genesisBlock:=Block{
		hash: "0",
		timestamp: time.Now(),
	}
	return Blockchain{
		genesisBlock,
		[]Block{genesisBlock},
		difficulty,
	}
}

func (b *Blockchain) AddBlock(platform string,Link string){
	blockData :=map[string]interface{}{
		"Platform": platform,
		"Link": Link,
	}
	lastBlock := b.chain[len(b.chain)-1]
	newBlock := Block{
		data: blockData,
		previousHash: lastBlock.hash,
		timestamp: time.Now(),
	}
	newBlock.mine(b.difficulty)
	b.chain = append(b.chain,newBlock)
}

func (b Blockchain) IsValid() bool{
	for i:=range b.chain[1:]{
		previousBlock:=b.chain[i]
		currentBlock:=b.chain[i+1]
		if currentBlock.hash != currentBlock.calculateHash()||currentBlock.previousHash != previousBlock.hash {
			return false
		}
	}
	return true
}
func (b Blockchain) GetBlockchainData() []Jsondata {
	var jsonData []Jsondata
	for _, block := range b.chain {
		platform := ""
		link := ""
		if platformVal, ok := block.data["Platform"]; ok {
			platform = platformVal.(string)
		}
		if linkVal, ok := block.data["Link"]; ok {
			link = linkVal.(string)
		}
		jsonData = append(jsonData, Jsondata{Platform: platform, Link: link})
	}
	return jsonData
}
//func main(){
//	blockChain:= CreateBlockchain(3)
//	blockChain.addBlock("Github","https://github.com/Aashish1-1-1")
//	blockChain.addBlock("Linkedln","https://linkedln.com/Aashish1-1-1")
//	blockChain.addBlock("Instagram","https://instagram.com/Blessadhikari")
//	blockChain.addBlock("FaceBook","https://facebook.com/Blessadhikari")
//	blockChain.printBlockchain()
//	fmt.Println(blockChain.isValid())
//}
