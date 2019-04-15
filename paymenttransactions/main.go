package main

import "fmt"

type transaction struct {
	TransactionID int
	SenderID      int
	ReceiverID    int
	Amount        int
}

type user struct {
	Uuid   int
	Wallet int
}

func main() {

	// Init the slice of users assuming everyone starts with $10
	users := []user{{Uuid: 001, Wallet: 10}, {Uuid: 002, Wallet: 10}, {Uuid: 003, Wallet: 10}}

	// load the transactions
	loadedtransactions := loadTransactions()

	// process the transactions
	processedtransactions := process(loadedtransactions, users)

	for _, x := range processedtransactions {
		fmt.Println(x.Uuid)
		fmt.Println(x.Wallet)
	}
}

func loadTransactions() []transaction {

	load := make([]transaction, 100)

	lunch := transaction{TransactionID: 001, SenderID: 001, ReceiverID: 002, Amount: 10}

	load = append(load, lunch)

	sticker := transaction{TransactionID: 002, SenderID: 002, ReceiverID: 003, Amount: 3}

	load = append(load, sticker)

	soda := transaction{TransactionID: 003, SenderID: 002, ReceiverID: 001, Amount: 1}

	load = append(load, soda)

	return load
}

// The overall Big-O would be O(n)
func process(in []transaction, users []user) []user {
	for _, x := range in {
		if x.SenderID == users[0].Uuid {
			users[0].Wallet = users[0].Wallet - x.Amount
		} else if x.ReceiverID == users[0].Uuid {
			users[0].Wallet = users[0].Wallet + x.Amount
		} else if x.SenderID == users[1].Uuid {
			users[1].Wallet = users[1].Wallet - x.Amount
		} else if x.ReceiverID == users[1].Uuid {
			users[1].Wallet = users[1].Wallet + x.Amount
		} else if x.SenderID == users[2].Uuid {
			users[2].Wallet = users[2].Wallet - x.Amount
		} else if x.ReceiverID == users[2].Uuid {
			users[2].Wallet = users[2].Wallet + x.Amount
		} else {
			return []user{}
		}
	}
	return users
}
