package main

import (
	"fmt"

	"github.com/iotaledger/iota.go/api"
	"github.com/iotaledger/iota.go/bundle"
	"github.com/iotaledger/iota.go/converter"
	"github.com/iotaledger/iota.go/trinary"
)

var network = "https://nodes.devnet.iota.org"

// define depth and minimum weight (tangle properties)
const depth = 3
const minimumWeightMagnitude = 9

// send to address does not have to belong to anyone. just needs to have 81 trytes.
const address = trinary.Trytes("ZLGVEQ9JUZZWCZXLWVNTHBDX9G9KZTJP9VEERIIFHY9SIQKYBVAHIMLHXPQVE9IXFDDXNHQINXJDRPFDXNYVAPLZAW")

// seed not used in zero-value transaction.  However, the library expects it so use a random string
const seed = trinary.Trytes("JBN9ZRCOH9YRUGSWIQNZWAIFEZUBDUGTFPVRKXWPAUCEQQFS9NHPQLXCKZKRHVCCUZNF9CZZWKXRZVCWQ")

func main() {

	// api settings req'd for connecting to a network
	apiRet, err := api.ComposeAPI(api.HTTPClientSettings{URI: network})
	must(err)

	// creat JSON formatted message (will be easier to read message in exercise part 2)
	var data = "{'message' : 'Hello world'}"
	// converted to trytes
	message, err := converter.ASCIIToTrytes(data)
	must(err)

	// define zero-value transaction
	transfers := bundle.Transfers{
		{
			Address: address,
			Value:   0,
			Message: message,
		},
	}

	// prepare transfers object as trytes
	trytes, err := apiRet.PrepareTransfers(seed, transfers, api.PrepareTransfersOptions{})
	must(err)

	// send trytes handles tip selection, remote proof of work, and sending the bundle to network
	myBundle, err := apiRet.SendTrytes(trytes, depth, minimumWeightMagnitude)
	must(err)

	// bundle hash of the transaction just sent.
	fmt.Println(bundle.TailTransactionHash(myBundle))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
