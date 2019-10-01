package utils

import (
	"fmt"
	"log"
	"net"
	"os"
	"projects/goBlockChain/app"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// GetHost returns localhost address.
func GetHost() string {
	host, _ := os.Hostname()
	addr, _ := net.LookupHost(host)
	return addr[0]
}

func isFoundHost(target, port string) bool {
	addr := target + ":" + port
	_, err := net.DialTimeout("tcp", addr, time.Duration(1)*time.Second)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// FindNeighbours search neighbours IP address and returns slice of it.
func FindNeighbours(myHost string, myPort, startIPRange, endIPRange, startPort, endPort int) []string {
	var neighbours []string
	regexIP := regexp.MustCompile("^(\\d{1,3}.\\d{1,3}.\\d{1,3}.)(\\d{1,3})$")
	result := regexIP.FindStringSubmatch(myHost)
	if len(result) == 0 {
		return neighbours
	}
	startIP := result[1]
	lastIP, _ := strconv.Atoi(result[2])

	addr := myHost + ":" + strconv.Itoa(myPort)

	for guessPort := startPort; guessPort <= endPort; guessPort++ {
		for ipRange := startIPRange; ipRange <= endIPRange; ipRange++ {
			guessHost := startIP + strconv.Itoa(lastIP+ipRange)
			guessAddress := guessHost + ":" + strconv.Itoa(guessPort)
			if isFoundHost(guessHost, strconv.Itoa(guessPort)) && guessAddress != addr {
				neighbours = append(neighbours, guessAddress)
			}
		}
	}
	return neighbours
}

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
			fmt.Printf(formatB, "SenderBlockchainAddress", v.SenderAddress)
			fmt.Printf(formatB, "RecipientAddress", v.RecipientAddress)
			fmt.Printf(formatB, "Value", v.Value)
			fmt.Println(strings.Repeat("-", 50))
		}
	}
	fmt.Printf("%s\n\n\n", strings.Repeat("*", 50))
}
