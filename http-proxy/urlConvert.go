package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"io/ioutil"
)

func DecodeBase58Url(url string) string {
	decoded := base58.Decode(url)
	fmt.Println("Base58 Decoded Data:", string(decoded))
	return string(decoded)
}

func EncodeBase58Url(url string) string {
	encoded := base58.Encode([]byte(url))
	fmt.Println("Base58 Encoded Data:", encoded)
	return encoded
}

func generateKeyPair(bits int) {
	// This method requires a random number of bits.
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		fmt.Println("Error in generating key: ", err)
	}
	fmt.Printf("Private key: %v\n", privateKey)
	fmt.Printf("Public Key: %v\n", privateKey.PublicKey)

	privKeyPem := string(pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		},
	))
	pubKeyPem := string(pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(&privateKey.PublicKey), // The public key is part of the PrivateKey struct
		},
	))
	// save private and public key as a string in PEM format in file
	ioutil.WriteFile("privkey.pem", []byte(privKeyPem), 0400)
	ioutil.WriteFile("pubkey.pem", []byte(pubKeyPem), 0400)
}

func encryptUrl(url string) string {

	pubKeyPEM, _ := ioutil.ReadFile("pubkey.pem")
	block, _ := pem.Decode(pubKeyPEM) // Decode public key struct from PEM string
	pubKey, _ := x509.ParsePKCS1PublicKey(block.Bytes)

	cipherText, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, pubKey, []byte(url), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Encrypted message: ", string(cipherText))
	cipherUrl := EncodeBase58Url(string(cipherText)) // encode with Base58 before encrypting
	return cipherUrl
}

func decryptUrl(url string) string {
	url = DecodeBase58Url(url) // decode the previously encoded Base58 layer

	privKeyPEM, _ := ioutil.ReadFile("privkey.pem")
	block, _ := pem.Decode(privKeyPEM)
	privKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)

	decMessage, _ := rsa.DecryptOAEP(sha256.New(), rand.Reader, privKey, []byte(url), nil)
	fmt.Printf("Decrypted message: %s\n", string(decMessage))
	return string(decMessage)
}

func mainTest() {
	// Generate 2048-bits keys, if keys are not present in current dir
	generateKeyPair(1024)

	//encryptUrl("/users/3") //363SSTsNXVSjHVmApEfMkLHHVDkmZ2Nujor1ZbpWR3kuX5g8fCE5n31hYVmABgL1KxjEo1Ad8Ctq7V5JxyCeiUSR5x49zx9ZxPAFUSAjJSCLfnh1pyETcGoXtFqgQqwt8GJ5m244xsYDEApXUVJDHrwF3eF8qzZTnqKdxdZyfZYocsMX
	decryptUrl("63SSTsNXVSjHVmApEfMkLHHVDkmZ2Nujor1ZbpWR3kuX5g8fCE5n31hYVmABgL1KxjEo1Ad8Ctq7V5JxyCeiUSR5x49zx9ZxPAFUSAjJSCLfnh1pyETcGoXtFqgQqwt8GJ5m244xsYDEApXUVJDHrwF3eF8qzZTnqKdxdZyfZYocsMX")
}
