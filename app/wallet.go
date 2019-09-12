package app

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"

	"golang.org/x/crypto/ripemd160"
)

func createKeyPair() *ecdsa.PrivateKey {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	return privateKey
}

// GenerateBlockchainAddress is return blockchainAddress.
// blockchainAddress is created by publicKey.
func GenerateBlockchainAddress() string {
	privateKey := createKeyPair()
	publicKey := privateKey.PublicKey
	publicKeyByte := elliptic.Marshal(publicKey.Curve, publicKey.X, publicKey.Y)
	sha256Encoder := sha256.New()
	sha256Encoder.Write(publicKeyByte)
	sha256Digest := sha256Encoder.Sum(nil)

	ripemd160Encoder := ripemd160.New()
	ripemd160Encoder.Write(sha256Digest)
	ripemd160Digest := ripemd160Encoder.Sum(nil)

	networkByte := []byte("11")
	networkBitcoinPubKey := append(networkByte, ripemd160Digest...)
	networkBitcoinPubKeyHex := hex.EncodeToString(networkBitcoinPubKey)

	sha256Encoder2 := sha256.New()
	sha256Encoder2.Write(networkBitcoinPubKey)
	sha256Digest = sha256Encoder2.Sum(nil)
	sha256Encoder2.Write(sha256Digest)
	sha256Digest2 := sha256Encoder2.Sum(nil)
	sha256Hex := hex.EncodeToString(sha256Digest2)

	checksum := sha256Hex[:8]

	addressHex := networkBitcoinPubKeyHex + checksum

	blockchainAddress := base64.StdEncoding.EncodeToString([]byte(addressHex))

	return blockchainAddress
}
