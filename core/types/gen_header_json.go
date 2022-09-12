// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/spruce-solutions/go-quai/common"
	"github.com/spruce-solutions/go-quai/common/hexutil"
)

var _ = (*headerMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (h Header) MarshalJSON() ([]byte, error) {
	var enc struct {
		ParentHash  []common.Hash    `json:"parentHash"       gencodec:"required"`
		UncleHash   []common.Hash    `json:"sha3Uncles"       gencodec:"required"`
		Coinbase    []common.Address `json:"miner"            gencodec:"required"`
		Root        []common.Hash    `json:"stateRoot"        gencodec:"required"`
		TxHash      []common.Hash    `json:"transactionsRoot" gencodec:"required"`
		ReceiptHash []common.Hash    `json:"receiptsRoot"     gencodec:"required"`
		Bloom       []Bloom          `json:"logsBloom"        gencodec:"required"`
		Difficulty  []*hexutil.Big   `json:"difficulty"       gencodec:"required"`
		Number      []*hexutil.Big   `json:"number"           gencodec:"required"`
		GasLimit    []hexutil.Uint64 `json:"gasLimit"         gencodec:"required"`
		GasUsed     []hexutil.Uint64 `json:"gasUsed"          gencodec:"required"`
		BaseFee     []*hexutil.Big   `json:"baseFeePerGas"    gencodec:"required"`
		Location 	common.Location  `json:"location"         gencodec:"required"`
		Time        hexutil.Uint64   `json:"timestamp"        gencodec:"required"`
		Extra       hexutil.Bytes    `json:"extraData"        gencodec:"required"`
		Nonce       BlockNonce       `json:"nonce"`
		Hash        common.Hash      `json:"hash"`
	}
	copy(enc.ParentHash, h.ParentHashArray())
	copy(enc.UncleHash, h.UncleHashArray())
	copy(enc.Coinbase, h.CoinbaseArray())
	copy(enc.Root, h.RootArray())
	copy(enc.TxHash, h.TxHashArray())
	copy(enc.ReceiptHash, h.ReceiptHashArray())
	copy(enc.Bloom, h.BloomArray())
	for i := 0; i < common.HierarchyDepth; i++ {
		enc.Difficulty[i] = (*hexutil.Big)(h.Difficulty(i))
		enc.Number[i] = (*hexutil.Big)(h.Number(i))
		enc.GasLimit[i] = hexutil.Uint64(h.GasLimit(i))
		enc.GasUsed[i] = hexutil.Uint64(h.GasUsed(i))
		enc.BaseFee[i] = (*hexutil.Big)(h.BaseFee(i))
	}
	enc.Location = h.Location()
	enc.Time = hexutil.Uint64(h.Time())
	enc.Extra = h.Extra()
	enc.Nonce = h.Nonce()
	enc.Hash = h.Hash()
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (h *Header) UnmarshalJSON(input []byte) error {
	var dec struct {
		ParentHash  []*common.Hash    `json:"parentHash"       gencodec:"required"`
		UncleHash   []*common.Hash    `json:"sha3Uncles"       gencodec:"required"`
		Coinbase    []*common.Address `json:"miner"            gencodec:"required"`
		Root        []*common.Hash    `json:"stateRoot"        gencodec:"required"`
		TxHash      []*common.Hash    `json:"transactionsRoot" gencodec:"required"`
		ReceiptHash []*common.Hash    `json:"receiptsRoot"     gencodec:"required"`
		Bloom       []*Bloom          `json:"logsBloom"        gencodec:"required"`
		Difficulty  []*hexutil.Big   `json:"difficulty"       gencodec:"required"`
		Number      []*hexutil.Big   `json:"number"           gencodec:"required"`
		GasLimit    []*hexutil.Uint64 `json:"gasLimit"         gencodec:"required"`
		GasUsed     []*hexutil.Uint64 `json:"gasUsed"          gencodec:"required"`
		BaseFee     []*hexutil.Big   `json:"baseFeePerGas"    gencodec:"required"`
		Location *common.Location 	 `json:"location"         gencodec:"required"`
		Time        *hexutil.Uint64  `json:"timestamp"        gencodec:"required"`
		Extra       *hexutil.Bytes   `json:"extraData"        gencodec:"required"`
		Nonce       *BlockNonce      `json:"nonce"`
	}
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.ParentHash == nil {
		return errors.New("missing required field 'parentHash' for Header")
	}
	if dec.UncleHash == nil {
		return errors.New("missing required field 'sha3Uncles' for Header")
	}
	if dec.Coinbase == nil {
		return errors.New("missing required field 'miner' for Header")
	}
	if dec.Root == nil {
		return errors.New("missing required field 'stateRoot' for Header")
	}
	if dec.TxHash == nil {
		return errors.New("missing required field 'transactionsRoot' for Header")
	}
	if dec.ReceiptHash == nil {
		return errors.New("missing required field 'receiptsRoot' for Header")
	}
	if dec.Bloom == nil {
		return errors.New("missing required field 'logsBloom' for Header")
	}
	if dec.Difficulty == nil {
		return errors.New("missing required field 'difficulty' for Header")
	}
	if dec.Number == nil {
		return errors.New("missing required field 'number' for Header")
	}
	if dec.GasLimit == nil {
		return errors.New("missing required field 'gasLimit' for Header")
	}
	if dec.GasUsed == nil {
		return errors.New("missing required field 'gasUsed' for Header")
	}
	if dec.BaseFee == nil {
		return errors.New("missing required field 'baseFee' for Header")
	}
	if dec.Time == nil {
		return errors.New("missing required field 'timestamp' for Header")
	}
	if dec.Extra == nil {
		return errors.New("missing required field 'extraData' for Header")
	}
	for i := 0; i < common.HierarchyDepth; i++ {
		h.SetParentHash(*dec.ParentHash[i], i)
		h.SetUncleHash(*dec.UncleHash[i], i)
		h.SetCoinbase(*dec.Coinbase[i], i)
		h.SetRoot(*dec.Root[i], i)
		h.SetTxHash(*dec.TxHash[i], i)
		h.SetReceiptHash(*dec.ReceiptHash[i], i)
		h.SetBloom(*dec.Bloom[i], i)
		h.SetDifficulty((*big.Int)(dec.Difficulty[i]), i)
		h.SetNumber((*big.Int)(dec.Number[i]), i)
		h.SetGasLimit(uint64(*dec.GasLimit[i]), i)
		h.SetGasUsed(uint64(*dec.GasUsed[i]), i)
	}
	h.SetLocation(dec.Location)
	h.SetTime(uint64(*dec.Time))
	h.SetExtra(*dec.Extra)
	h.SetNonce(*dec.Nonce)
	return nil
}
