package app

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"

	"golang.org/x/crypto/ripemd160"
)

// CreateKeyPair is
func CreateKeyPair() *ecdsa.PrivateKey {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	return privateKey
}

// GenerateBlockchainAddress is return blockchainAddress.
// blockchainAddress is created by publicKey.
func GenerateBlockchainAddress(privateKey ecdsa.PrivateKey) string {
	publicKey := privateKey.PublicKey
	publicKeyByte := elliptic.Marshal(publicKey.Curve, publicKey.X, publicKey.Y)
	sha256Encoder := sha256.New()
	sha256Encoder.Write(publicKeyByte)
	sha256Digest := sha256Encoder.Sum(nil)

	ripemd160Encoder := ripemd160.New()
	ripemd160Encoder.Write(sha256Digest)
	ripemd160Digest := ripemd160Encoder.Sum(nil)

	networkByte := []byte("00")
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

// GenerateSignature is
func GenerateSignature(senPriKey ecdsa.PrivateKey, senBA string, recBA string, val float64) string {
	tx := Transaction{
		SenderPrivateKey: senPriKey,
		SenderPublicKey:  senPriKey.PublicKey,
		SenderAddress:    senBA,
		RecipientAddress: recBA,
		Value:            val,
	}
	var opts crypto.SignerOpts
	message := tx.hash()
	privateKeySign, _ := senPriKey.Sign(rand.Reader, message, opts)
	return hex.EncodeToString(privateKeySign)
}
