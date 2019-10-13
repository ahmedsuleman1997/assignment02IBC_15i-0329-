package main

import (
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	//"github.com/assignment02IBC_15i-0329 "
)

type Block struct {
	

	Transaction string
	PrevPointer *Block
	Hash        [32]byte
	PrevHash    [32]byte
}

var blockchain *Block

func DeriveHash(Transaction string) [32]byte {
	return sha256.Sum256([]byte(Transaction))
}

func InsertBlock(Transaction string, chainHead *Block) *Block {
	if chainHead == nil {
		return &Block{Transaction, nil, DeriveHash(Transaction), [32]byte{}}
	}
	return &Block{Transaction, chainHead, DeriveHash(Transaction), DeriveHash(chainHead.Transaction)}

}

func ListBlocks(chainHead *Block) {
	for p := chainHead; p != nil; p = p.PrevPointer {
		fmt.Printf("Transaction: %s, Hash:%x, PrevHash:%x\n", p.Transaction, p.Hash, p.PrevHash)

	}
}
func handleConnection(c net.Conn) {

	log.Println("A client has connected", c.RemoteAddr())
	//block1 := &Blocks{"Satoshis100", nil}
	//block2 := &Blocks{"Satoshi50Alice50", block1}
	gobEncoder := gob.NewEncoder(c)
	err := gobEncoder.Encode(blockchain)
	if err != nil {
		log.Println(err)
	}
}
func main() {
	conn, err := net.Dial("tcp", "localhost:6000")
	if err != nil {
		//handle error
	}
	//var currentblock *Block
	var recvdBlock Block
	dec := gob.NewDecoder(conn)
	err = dec.Decode(&recvdBlock)
	if err != nil {

	//Error handling

	}
	ListBlocks(&recvdBlock)

	/*	ln, err := net.Listen("tcp", ":6005")
		if err != nil {
			log.Fatal(err)
		}
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Println(err)
				continue
			}
			go handleConnection(conn)

		}*/
}
