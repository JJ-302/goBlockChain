package app

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"math/big"

	"golang.org/x/crypto/ripemd160"
)

// Wallet represents wallet for bitcoin.
type Wallet struct {
	PrivateKey        ecdsa.PrivateKey
	PublicKey         ecdsa.PublicKey
	BlockchainAddress string
}

// Signature is result it signs using the privateKey.
type Signature struct {
	X, Y *big.Int
}

// WalletPool is an instead of DB for wallet.
var WalletPool = make(map[string]Wallet)

// CreateWallet create wallet for bitcoin.
func CreateWallet() *Wallet {
	privateKey := createKeyPair()
	wallet := Wallet{
		PrivateKey:        *privateKey,
		PublicKey:         privateKey.PublicKey,
		BlockchainAddress: generateBlockchainAddress(*privateKey),
	}
	WalletPool[wallet.BlockchainAddress] = wallet
	return &wallet
}

func createKeyPair() *ecdsa.PrivateKey {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	return privateKey
}

func generateBlockchainAddress(privateKey ecdsa.PrivateKey) string {
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

// GenerateSignature retuens signature.
func (wallet *Wallet) GenerateSignature(recBA string, val float64) *Signature {
	tx := CreateTransaction(wallet.BlockchainAddress, recBA, val)
	message := tx.hash()
	var sign Signature
	sign.X, sign.Y, _ = ecdsa.Sign(rand.Reader, &wallet.PrivateKey, message)
	return &sign
}

func (signature *Signature) verifyTransactionSignature(senPubKey *ecdsa.PublicKey, tx Transaction) bool {
	message := tx.hash()
	return ecdsa.Verify(senPubKey, message, signature.X, signature.Y)
}
