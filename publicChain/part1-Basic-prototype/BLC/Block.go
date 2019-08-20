package BLC

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Height        int64
	PrevBlockHash []byte;
	Data          []byte
	Timestamp     int64
	Hash          []byte
}

//NewBlock 创建新的区块
func NewBlock(data string, height int64, prevBlockHash []byte) *Block {
	return &Block{
		Height:        height,
		PrevBlockHash: prevBlockHash,
		Data:          []byte(data),
		Timestamp:     time.Now().Unix(),
		Hash:          nil,
	}
}

func (block *Block) SetHash() {
	// 1 height 转 []byte
	heightBytes := IntToHex(block.Height)

	// 2 时间戳转[]byte
	timeString := strconv.FormatInt(block.Timestamp, 2)
	timeBytes := []byte(timeString)

	// 3 拼接所有属性
	blockBytes := bytes.Join([][]byte{heightBytes, block.PrevBlockHash, block.Data, timeBytes, block.Hash}, []byte{})

	// 4 生成 hash
	hash := sha256.Sum256(blockBytes)
	block.Hash = hash[:]
}
