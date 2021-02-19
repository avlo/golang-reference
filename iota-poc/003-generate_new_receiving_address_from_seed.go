////////////////////////////////////////////////
// Generate an unspent address
////////////////////////////////////////////////

package main

import (
	"fmt"

	"github.com/iotaledger/iota.go/api"
	"github.com/iotaledger/iota.go/trinary"
)

var node = "http://localhost:14265"

// head /dev/urandom |tr -dc A-Z9|head -c${1:-81}
// generates 81digit number
// The seed that will be used to generate a (receiving) address
const seed = trinary.Trytes("ZMWOEGE9PVYCOANUWDXLLQKKCQBREVPEIDXFBBYDHTHRRRVY9TTKQEBHQDOQHCCHCFDVOGZNIZMQIOKD9")

//////////////////////
// output address:
// L9TXCAXAXMPJDUJJRKKDQEUBNGFDMQHLEAENJYRVBVSBXOZKJLMOAX9RNVGTLDIZOJWORFLGNGKUAXS9BMFVYHRUTD
//
// devnet faucet
// https://faucet.devnet.iota.org/
//
// "Your allocation request has been successfully processed
//  d414221b-6c90-4b69-82ac-603447bb01de
//
// devnet explorer
// https://explorer.iota.org/devnet/address/L9TXCAXAXMPJDUJJRKKDQEUBNGFDMQHLEAENJYRVBVSBXOZKJLMOAX9RNVGTLDIZOJWORFLGNGKUAXS9BMFVYHRUTD
//////////////////////

// Define the security level of the address
const securityLevel = 1

func main() {
	// Connect to a node
	apiRet, err := api.ComposeAPI(api.HTTPClientSettings{URI: node})
	must(err)

	// Generate an unspent address with security level 2
	// If this address is spent, this method returns the next unspent address with the lowest index
	addresses, err := apiRet.GetNewAddress(seed, api.GetNewAddressOptions{Security: securityLevel})
	must(err)

	fmt.Println("\nYour address is: ", addresses[0])
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
