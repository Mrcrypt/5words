package domain

import (
	"crypto/rand"
	"testing"
)

func TestGenerate(t *testing.T) {
	var randomData []byte = make([]byte, 64)
	rand.Read(randomData)
	//NewDomain(randomData)
}

func TestMnemonic(t *testing.T) {
	/*
		var randomData []byte = make([]byte, 64)
		rand.Read(randomData)
		dom, _ := NewDomain(randomData)

		fmt.Printf("dom. =\t%016x\n", dom.Name)

		res := dom.ToMnemonic()
		fmt.Printf("res = %s\n", res)
		dom2, _ := NewDomainMnemonic(res)
		fmt.Printf("dom2. =\t%016x\n", dom2.Name)

		isvalid := ValidNonce(dom.PubKey, dom.Nonce)
		fmt.Printf("isvalid = %+v\n", isvalid)*/
}
