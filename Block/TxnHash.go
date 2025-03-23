package Block

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Hash returns the Keccak256 hash of the transaction
func Hash(tx *Transaction) common.Hash {
    var ethTx *types.Transaction
    
    switch {
    case tx.MaxFeePerGas != nil:
        // EIP-1559 transaction
        accessList := convertAccessList(tx.AccessList)
        ethTx = types.NewTx(&types.DynamicFeeTx{
            ChainID:    tx.ChainID,
            Nonce:      tx.Nonce,
            GasTipCap:  tx.MaxPriorityFeePerGas,
            GasFeeCap:  tx.MaxFeePerGas,
            Gas:        tx.GasLimit,
            To:         tx.To,
            Value:      tx.Value,
            Data:       tx.Data,
            AccessList: accessList,
            V:          tx.V,
            R:          tx.R,
            S:          tx.S,
        })
    case tx.AccessList != nil && len(tx.AccessList) > 0:
        // EIP-2930 transaction
        accessList := convertAccessList(tx.AccessList)
        ethTx = types.NewTx(&types.AccessListTx{
            ChainID:    tx.ChainID,
            Nonce:      tx.Nonce,
            GasPrice:   tx.GasPrice,
            Gas:        tx.GasLimit,
            To:         tx.To,
            Value:      tx.Value,
            Data:       tx.Data,
            AccessList: accessList,
            V:          tx.V,
            R:          tx.R,
            S:          tx.S,
        })
    default:
        // Legacy transaction
        ethTx = types.NewTx(&types.LegacyTx{
            Nonce:    tx.Nonce,
            GasPrice: tx.GasPrice,
            Gas:      tx.GasLimit,
            To:       tx.To,
            Value:    tx.Value,
            Data:     tx.Data,
            V:        tx.V,
            R:        tx.R,
            S:        tx.S,
        })
    }
    
    // Return the hash of the transaction
    return ethTx.Hash()
}