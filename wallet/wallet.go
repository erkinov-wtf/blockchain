package wallet

import (
	"blockchain/utils"
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"

	"golang.org/x/crypto/ripemd160"
)

const version = byte(0x00)
const walletFile = "wallet.dat"
const addressChecksumLen = 4

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

func NewWallet() *Wallet {
	private, public := newKeyPair()
	wallet := Wallet{private, public}

	return &wallet
}

func (w Wallet) GetAddress() []byte {
	pubKeyHash := HashPubKey(w.PublicKey)

	versionedPayload := append([]byte{version}, pubKeyHash...)
	checksum := checksum(versionedPayload)

	fullPayload := append(versionedPayload, checksum...)
	address := utils.Base58Encode(fullPayload)

	return address
}

func HashPubKey(pubKey []byte) []byte {
	publicSHA256 := sha256.Sum256(pubKey)

	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	if err != nil {
		log.Panic(err)
	}
	publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)

	return publicRIPEMD160
}

func ValidateAddress(address string) bool {
	pubKeyHash := utils.Base58Decode([]byte(address))
	actualChecksum := pubKeyHash[len(pubKeyHash)-addressChecksumLen:]
	version := pubKeyHash[0]
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-addressChecksumLen]
	targetChecksum := checksum(append([]byte{version}, pubKeyHash...))

	return bytes.Compare(actualChecksum, targetChecksum) == 0
}

func checksum(payload []byte) []byte {
	firstSHA := sha256.Sum256(payload)
	secondSHA := sha256.Sum256(firstSHA[:])

	return secondSHA[:addressChecksumLen]
}

func newKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	return *private, pubKey
}

// Serialize the wallet
func (w *Wallet) Serialize() []byte {
	privateKeyBytes := w.PrivateKey.D.Bytes()
	publicKeyBytes := w.PublicKey

	// Format: privateKeyHex,publicKeyHex
	return []byte(fmt.Sprintf("%x,%x", privateKeyBytes, publicKeyBytes))
}

// Deserialize the wallet
func (w *Wallet) Deserialize(data []byte) error {
	// Format: privateKeyHex,publicKeyHex
	parts := strings.Split(string(data), ",")
	if len(parts) != 2 {
		return errors.New("invalid data format")
	}

	privateKeyBytes, err := hex.DecodeString(parts[0])
	if err != nil {
		return err
	}

	publicKeyBytes, err := hex.DecodeString(parts[1])
	if err != nil {
		return err
	}

	// Reconstruct private key
	curve := elliptic.P256()
	privateKey := ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: curve,
			X:     new(big.Int).SetBytes(publicKeyBytes[:len(publicKeyBytes)/2]),
			Y:     new(big.Int).SetBytes(publicKeyBytes[len(publicKeyBytes)/2:]),
		},
		D: new(big.Int).SetBytes(privateKeyBytes),
	}

	w.PrivateKey = privateKey
	w.PublicKey = publicKeyBytes

	return nil
}
