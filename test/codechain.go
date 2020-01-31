package main

import (
	"encoding/hex"
	"fmt"

	"github.com/CodeChain-io/codechain-sdk-go/key"
)

func main() {

	networkID := "tc"
	ecdsa, err := key.GenerateEcdsa()
	if err != nil {
		fmt.Println("ECDSA key pair generation error")
		return
	}

	platformAddress, err := key.CreatePlatformAddress(ecdsa, networkID)

	if err != nil {
		fmt.Println("PlatformAddress creation error: ", err)
		return
	}
	fmt.Println("platformAddress: ", platformAddress.Value)
	fmt.Println("accountID: ", platformAddress.AccountID.ToHexString())
	fmt.Println("private key: ", hex.EncodeToString(ecdsa.GetPrivateKey()))
	fmt.Println("public key: ", hex.EncodeToString(ecdsa.GetPublicKey()))
}
