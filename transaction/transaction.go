package transaction

type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}

type TXOutput struct {
	Value        int
	ScriptPublic string
}

type TXInput struct {
	Txid      []byte
	Vout      int
	ScriptSig string
}

const subsidy = 10
