package wallet

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Wallets struct {
	Wallets map[string]*Wallet
}

func NewWallets() (*Wallets, error) {
	wallets := Wallets{}
	wallets.Wallets = make(map[string]*Wallet)

	err := wallets.LoadFromFile()

	return &wallets, err
}

func (ws *Wallets) CreateWallet() string {
	wallet := NewWallet()
	address := fmt.Sprintf("%s", wallet.GetAddress())

	ws.Wallets[address] = wallet

	return address
}

func (ws *Wallets) GetAddresses() []string {
	var addresses []string

	for address := range ws.Wallets {
		addresses = append(addresses, address)
	}

	return addresses
}

func (ws Wallets) GetWallet(address string) Wallet {
	return *ws.Wallets[address]
}

func (ws *Wallets) SaveToFile() {
	var content bytes.Buffer

	for _, wallet := range ws.Wallets {
		walletData := wallet.Serialize()
		content.Write(walletData)
		content.Write([]byte("\n")) // Add newline to separate wallets
	}

	err := os.WriteFile(walletFile, content.Bytes(), 0644)
	if err != nil {
		log.Panic(err)
	}
}

func (ws *Wallets) LoadFromFile() error {
	if _, err := os.Stat(walletFile); os.IsNotExist(err) {
		return err
	}

	fileContent, err := ioutil.ReadFile(walletFile)
	if err != nil {
		log.Panic(err)
	}

	lines := bytes.Split(fileContent, []byte("\n"))
	for _, line := range lines {
		if len(line) > 0 {
			var wallet Wallet
			err := wallet.Deserialize(line)
			if err != nil {
				log.Panic(err)
			}

			address := fmt.Sprintf("%s", wallet.GetAddress())
			ws.Wallets[address] = &wallet
		}
	}

	return nil
}
