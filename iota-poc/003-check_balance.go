////////////////////////////////////////////////
// Check the balance of an address
////////////////////////////////////////////////

package main

import (
	"fmt"

	. "github.com/iotaledger/iota.go/api"
	"github.com/iotaledger/iota.go/trinary"
)

var node = "http://localhost:14265"

// The address whose balance you want to check
const address = trinary.Trytes("KFGDPCWVNOLYX9BKEICEJNZIACBZUOJESWNXSCVGJZLHNSFYHXVDQQFODYFKSRMBNIANVWYCQNRUVSDCAWXHSWAFMW")

func main() {
	// Connect to a node
	api, err := ComposeAPI(HTTPClientSettings{URI: node})
	must(err)

	// Get the confirmed balance from the node
	balances, err := api.GetBalances(trinary.Hashes{address}, "100")
	must(err)
	fmt.Println("\nBalance:", balances.Balances[0])
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
