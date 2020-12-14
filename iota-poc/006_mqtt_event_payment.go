////////////////////////////////////////////////
// Send a microtransaction
////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"

	"github.com/iotaledger/iota.go/api"
	"github.com/iotaledger/iota.go/bundle"
	"github.com/iotaledger/iota.go/trinary"
)

// Replace this seed with the one that owns the address you used to get free test tokens
const seed = trinary.Trytes("SIEGWQTEKSXYKGBVKSFTNVJEQZHSUZTJFOZYSEVFPAKTQLPAMGLTUKFNYLYKPDNSH9NXIJJS9BWQDOMY9")

const minimumWeightMagnitude = 9
const depth = 3

func main() {
	// Connect to a node
	apiRef, err := api.ComposeAPI(api.HTTPClientSettings{URI: os.Args[1]})
	must(err)

	// Define an address to which to send IOTA tokens
	address := trinary.Trytes("L9TXCAXAXMPJDUJJRKKDQEUBNGFDMQHLEAENJYRVBVSBXOZKJLMOAX9RNVGTLDIZOJWORFLGNGKUAXS9BMFVYHRUTD")

	// Define an input transaction object
	// that sends 1 i to your new address
	transfers := bundle.Transfers{
		{
			Address: address,
			Value:   1,
		},
	}

	fmt.Println("Sending 1 i to " + address)

	trytes, err := apiRef.PrepareTransfers(seed, transfers, api.PrepareTransfersOptions{})
	must(err)

	myBundle, err := apiRef.SendTrytes(trytes, depth, minimumWeightMagnitude)
	must(err)

	fmt.Println(myBundle)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
