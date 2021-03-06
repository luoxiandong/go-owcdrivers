package polkadotTransaction

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/blocktree/go-owcdrivers/polkadotTransaction/codec"
)

type TxStruct struct {
	//MethodName string `json:"method_name"`
	//Version int `json:"version"`
	SenderPubkey string `json:"sender_pubkey"`
	RecipientPubkey string `json:"recipient_pubkey"`
	Amount uint64 `json:"amount"`
	Nonce uint64 `json:"nonce"`
	Fee uint64 `json:"fee"`
	BlockHeight uint64 `json:"block_height"`
	BlockHash string `json:"block_hash"`
	GenesisHash string `json:"genesis_hash"`
	SpecVersion uint32 `json:"spec_version"`
}


func (tx TxStruct) NewTxPayLoad() (*TxPayLoad, error) {
	var tp TxPayLoad
	method, err := NewMethodTransfer(tx.RecipientPubkey, tx.Amount)
	if err != nil {
		return nil, err
	}

	tp.Method, err = method.ToBytes()
	if err != nil {
		return  nil, err
	}

	if tx.BlockHeight == 0 {
		return nil, errors.New("invalid block height")
	}

	tp.Era = GetEra(tx.BlockHeight)

	if tx.Nonce == 0 {
		tp.Nonce = []byte{0}
	} else {
		nonce, err := codec.Encode(Compact_U32, uint64(tx.Nonce))
		if err != nil {
			return nil, err
		}
		tp.Nonce, _ = hex.DecodeString(nonce)
	}

	if tx.Fee == 0 {
		//return nil, errors.New("a none zero fee must be payed")
		tp.Fee = []byte{0}
	} else {
		fee, err := codec.Encode(Compact_U32, uint64(tx.Fee))
		if err != nil {
			return nil, err
		}
		tp.Fee, _ = hex.DecodeString(fee)
	}

	specv := make([]byte, 4)
	binary.LittleEndian.PutUint32(specv, tx.SpecVersion)
	tp.SpecVersion = specv

	genesis, err := hex.DecodeString(tx.GenesisHash)
	if err != nil || len(genesis) != 32 {
		return nil, errors.New("invalid genesis hash")
	}

	tp.GenesisHash = genesis

	block, err := hex.DecodeString(tx.BlockHash)
	if err != nil || len(block) != 32 {
		return nil, errors.New("invalid block hash")
	}

	tp.BlockHash = block

	return &tp, nil
}

func (tx TxStruct) ToJSONString() string {
	j, _ := json.Marshal(tx)
	
	return string(j)
}

func NewTxStructFromJSON(j string) (*TxStruct, error) {

	ts := TxStruct{}

	err := json.Unmarshal([]byte(j), &ts)

	if err != nil {
		return nil, err
	}

	return &ts, nil
}

func (ts TxStruct) GetSignedTransaction (signature string) (string, error) {

	signed := make([]byte, 0)

	signed = append(signed, SigningBitV4)

	if AccounntIDFollow {
		signed = append(signed, 0xff)
	}

	from, err := hex.DecodeString(ts.SenderPubkey)
	if err != nil || len(from) != 32 {
		return "", nil
	}

	signed = append(signed, from...)

	signed = append(signed, 0x00) // ed25519

	sig, err := hex.DecodeString(signature)
	if err != nil || len(sig) != 64 {
		return "", nil
	}
	signed = append(signed, sig...)

	if ts.BlockHeight == 0 {
		return "", errors.New("invalid block height")
	}

	signed = append(signed, GetEra(ts.BlockHeight)...)

	if ts.Nonce == 0 {
		signed = append(signed, 0)
	} else {
		nonce, err := codec.Encode(Compact_U32, uint64(ts.Nonce))
		if err != nil {
			return "", err
		}

		nonceBytes, _ := hex.DecodeString(nonce)
		signed = append(signed, nonceBytes...)
	}

	feeBytes := make([]byte, 0)
	if ts.Fee == 0 {
		//return "", errors.New("a none zero fee must be payed")
		feeBytes = []byte{0}
	} else {
		fee, err := codec.Encode(Compact_U32, uint64(ts.Fee))
		if err != nil {
			return "", err
		}
		feeBytes, _ = hex.DecodeString(fee)
	}

	signed = append(signed, feeBytes...)

	method, err := NewMethodTransfer(ts.RecipientPubkey, ts.Amount)
	if err != nil {
		return "", err
	}

	methodBytes, err := method.ToBytes()
	if err != nil {
		return "", err
	}

	signed = append(signed, methodBytes...)

	length, err := codec.Encode(Compact_U32, uint64(len(signed)))
	if err != nil {
		return "", err
	}
	lengthBytes, _ := hex.DecodeString(length)
	lengthBytes[0] += 1
	return "0x" + hex.EncodeToString(lengthBytes) + hex.EncodeToString(signed), nil
}