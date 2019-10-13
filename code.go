package main

import (
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
	//"github.com/assignment02IBC_15i-0329 "
)

type Block struct {
	//Hash
	//Data

	Transaction string
	PrevPointer *Block
	Hash        [32]byte
	PrevHash    [32]byte
}
type User struct {
	Name       string
	Portnumber int
	Bitcoins   int
}

var blockchain *Block

//var X User

func DeriveHash(Transaction string) [32]byte {
	return sha256.Sum256([]byte(Transaction))
}

func InsertBlock(Transaction string, chainHead *Block) *Block {
	if chainHead == nil {
		return &Block{Transaction, nil, DeriveHash(Transaction), [32]byte{}}
	}
	return &Block{Transaction, chainHead, DeriveHash(Transaction), DeriveHash(chainHead.Transaction)}

}
func checkNode(sendername string, receivername string, minername string, deliveryamount int, reward int, X *User, Y *User, Z *User) {
	if X.Bitcoins >= deliveryamount {
		str1 := strconv.Itoa(deliveryamount)
		str2 := strconv.Itoa(reward)
		blockchain = InsertBlock(sendername+" send "+str1+" to "+receivername+" where miner will be "+minername+" the reward is "+str2+" given to "+minername, blockchain)
		X.Bitcoins = X.Bitcoins - deliveryamount
		Z.Bitcoins = Z.Bitcoins + deliveryamount
		Y.Bitcoins = Y.Bitcoins + reward
		X.Bitcoins = X.Bitcoins - reward
		fmt.Println("The block was added successfully")
	} else {
		fmt.Println(sendername + " doesnot have eenough bitcoins ")
	}
}
func ListBlocks(chainHead *Block) {
	for p := chainHead; p != nil; p = p.PrevPointer {
		fmt.Printf("Transaction: %s, Hash:%x, PrevHash:%x\n", p.Transaction, p.Hash, p.PrevHash)

	}
}



type Blocks struct {
	Transaction string
	PrevPointer *Blocks
}

func handleConnection(c net.Conn) {

	log.Println("A client has connected", c.RemoteAddr())
	
	gobEncoder := gob.NewEncoder(c)
	err := gobEncoder.Encode(blockchain)
	if err != nil {
		log.Println(err)
	}
}

var Xes int 
var Uarray = make([]User, 0)

func checksender(n string, xes int) int {
	var x int
	for i := 0; i <= xes; i++ {
		x = strings.Compare(Uarray[i].Name, n)
		if x == 0 {
			return i
		}
	}
	return -1
}
func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
func main() {
	var sender, receiver, miner string
	var endcondition int
	var amountsend, minerreward int
	fmt.Println("Users in the blockchain = ")
	fmt.Scan(&Xes)
	var name string
	var money int
	var port = 6001
	for i := Xes; i >= 0; i-- {
		fmt.Println("Enter the name of the user")
		fmt.Scan(&name)
		fmt.Println("Enter the bitcoins of the user")
		fmt.Scan(&money)
		Uarray = append(Uarray, User{Name: name, Portnumber: port, Bitcoins: money})
	}
	
	blockchain = InsertBlock("THE GENESIS BLOCK", blockchain)
	
	for endcondition != -1 {
		fmt.Println("Do you want to continue the transaction process??")
		fmt.Scan(&endcondition)
		if endcondition == -1 {
			break
		}
		fmt.Println("Bitcoin Sender Name :")
		fmt.Scan(&sender)
		fmt.Println("Bitcoin Receiver Name")
		fmt.Scan(&receiver)
		fmt.Println("BitCoin Miner Name")
		fmt.Scan(&miner)
		fmt.Println("Bitcoin Amount = " + sender + " you will send " + receiver)
		fmt.Scan(&amountsend)
		fmt.Println("Miner Reward=")
		fmt.Scan(&minerreward)
		checkNode(Uarray[checksender(sender, Xes)].Name, Uarray[checksender(receiver, Xes)].Name, Uarray[checksender(miner, Xes)].Name, amountsend, minerreward, &Uarray[checksender(sender, Xes)], &Uarray[checksender(miner, Xes)], &Uarray[checksender(receiver, Xes)])
		checkNode(Uarray[2].Name, Uarray[1].Name, Uarray[3].Name, 4, 1, &Uarray[2], &Uarray[3], &Uarray[1])
		checkNode(Uarray[2].Name, Uarray[0].Name, Uarray[2].Name, 8, 3, &Uarray[2], &Uarray[2], &Uarray[0])
		checkNode(Uarray[2].Name, Uarray[2].Name, Uarray[1].Name, 7, 2, &Uarray[2], &Uarray[1], &Uarray[2])

		ListBlocks(blockchain)
		fmt.Println(Uarray)
	

		ln, err := net.Listen("tcp", ":6000")
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

		}

		

	}

}
