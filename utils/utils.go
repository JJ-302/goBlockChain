package utils

import (
	"fmt"
	"projects/goBlockChain/app"
	"strconv"
	"strings"
	"time"
)

// Printblock is print blockchain on terminal.
func Printblock() {
	headerLine := strings.Repeat("=", 25)
	formatA := "%-15s : %v\n"
	formatB := "%-25s : %v\n"
	for i, v := range app.Chain {
		fmt.Println(headerLine + "Chain" + strconv.Itoa(i) + headerLine)
		fmt.Printf(formatA, "PreviousHash", v.PreviousHash)
		fmt.Printf(formatA, "Timestamp", v.Timestamp.Format(time.RFC3339))
		fmt.Printf(formatA, "Nonce", v.Nonce)
		fmt.Println("Transactions")
		fmt.Println(strings.Repeat("-", 50))
		for _, v := range v.Transactions {
			fmt.Printf(formatB, "SenderBlockchainAddress", v.SenderBlockchainAddress)
			fmt.Printf(formatB, "RecipientAddress", v.RecipientAddress)
			fmt.Printf(formatB, "Value", v.Value)
		}
	}
	fmt.Printf("%s\n\n\n", strings.Repeat("*", 50))
}
