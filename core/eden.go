package core

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
)

const (
	// BundleTotalGasLimit is the maximum gas limit for all bundle txs
	BundleTotalGasLimit = 4500000

	// SlotGasLimit is the maximum gas limit for each slot
	SlotGasLimit = 1500000

	// slotsAmount is the max slot amount in one block
	slotsAmount int = 3

	// ropstenEdenProxyAddress is the address that can lookup slot owner address on Ropsten
	ropstenEdenProxyAddress = "0xaa75DE4acC8590CF8299106b24656cDa2357C458"

	// mainnetEdenProxyAddress is the address that can lookup slot owner address on Mainnet
	mainnetEdenProxyAddress = "0x9E3382cA57F4404AC7Bf435475EAe37e87D1c453"
)

type Eden struct {
	contractAddr common.Address
	enable bool
	expireLocations [slotsAmount]common.Hash
	delegateLocations [slotsAmount]common.Hash
}

func NewEden(chainId uint64) *Eden {
	p := new(Eden)
	p.enable = false
	// 1=mainnet, 3=Ropsten, 4=Rinkeby, 5=Goerli
	if chainId == 1 {
		p.contractAddr = common.HexToAddress(mainnetEdenProxyAddress)
		p.enable = true
	} else if chainId == 3 {
		p.contractAddr = common.HexToAddress(ropstenEdenProxyAddress)
		p.enable = true
	} else {
	}

	if p.contractAddr != *new(common.Address) {
		for i := 0; i < slotsAmount; i++ {
			p.expireLocations[i] = p.expireLocation(i)
			p.delegateLocations[i] = p.delegateLocation(i)
		}
	}
	return p
}

func (e *Eden) Enable(londonForked bool) bool {
	if e.contractAddr == *new(common.Address) || !londonForked {
		return false
	}

	if !e.enable {
		e.enable = true
	}
	return true
}

func (e *Eden) SetTransactionsStake(parentStatedb *state.StateDB, txs map[common.Address]types.Transactions) {
	staked := make(map[common.Address]*big.Int)
	for addr, tx := range txs {
		for i,_ := range tx {
			_, contains := staked[addr]
			if !contains {
				staked[addr] = e.GetStakedBalance(parentStatedb, addr)
			}
			tx[i].SetStake(staked[addr])
		}
	}
}

func (e *Eden) GetStakedBalance(parentStatedb *state.StateDB, addr common.Address) *big.Int {
	location := e.stakedBalanceLocation(addr)
	v := parentStatedb.GetState(e.contractAddr, location)
	return new(big.Int).SetBytes(v.Bytes())
}

func (e *Eden) stakedBalanceLocation(addr common.Address) common.Hash {
	b, _ := hex.DecodeString("000000000000000000000000" + addr.String()[2:] +"0000000000000000000000000000000000000000000000000000000000000005")
	storageKey := crypto.Keccak256Hash(b)
	return storageKey
}

func (e *Eden) GetSlotAddress(parentStatedb *state.StateDB, headerTimestamp uint64) (map[common.Address]bool, []common.Address) {
	result := make(map[common.Address]bool)
	var resultSorted []common.Address

	for i := 0; i < slotsAmount; i++ {
		slotDelegateLocation := e.delegateLocations[i]
		slotExpirationLocation := e.expireLocations[i]
		slotAddr, err := e.isSlotExpired(parentStatedb, headerTimestamp, slotExpirationLocation, slotDelegateLocation)
		if err == nil {
			resultSorted = append(resultSorted, slotAddr)
		}
	}
	for _, a := range resultSorted {
		result[a] = true
	}
	return result, resultSorted
}

func (e *Eden) expireLocation(slotNumber int) common.Hash {
	str := fmt.Sprintf("%064x", slotNumber)
	b, _ := hex.DecodeString(str +"0000000000000000000000000000000000000000000000000000000000000001")
	storageKey := crypto.Keccak256Hash(b)
	return storageKey
}

func (e *Eden) delegateLocation(slotNumber int) common.Hash {
	str := fmt.Sprintf("%064x", slotNumber)
	b, _ := hex.DecodeString(str +"0000000000000000000000000000000000000000000000000000000000000002")
	storageKey := crypto.Keccak256Hash(b)
	return storageKey
}

func (e *Eden) isSlotExpired(parentStatedb *state.StateDB, headerTimestamp uint64, slotExpirationLocation, slotDelegateLocation common.Hash) (common.Address, error) {
	expiration := parentStatedb.GetState(e.contractAddr, slotExpirationLocation)
	expirationTimestamp := new(big.Int).SetBytes(expiration.Bytes()).Uint64()
	if headerTimestamp > expirationTimestamp {
		if expirationTimestamp != 0 {
			log.Info("Slot expired", "header", headerTimestamp, "expiration", expirationTimestamp)
		}
		return common.Address{}, errors.New("slot expired")
	}

	slot := parentStatedb.GetState(e.contractAddr, slotDelegateLocation)
	// exclude address 0 if this slot don't have an owner yet
	var addr0 common.Hash
	if slot == addr0 {
		return common.Address{}, errors.New("slot is address 0")
	}
	slotAddr := common.BytesToAddress(slot.Bytes())
	return slotAddr, nil
}
