package core

import (
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestEdenLocations(t *testing.T) {
	slotDelegate := []string{"0xac33ff75c19e70fe83507db0d683fd3465c996598dc972688b7ace676c89077b",
		"0xe90b7bceb6e7df5418fb78d8ee546e97c83a08bbccc01a0644d599ccd2a7c2e0",
		"0x679795a0195a1b76cdebb7c51d74e058aee92919b8c3389af86ef24535e8a28c"}

	slotExpiration := []string{"0xa6eef7e35abe7026729641147f7915573c7e97b47efa546f5f6e3230263bcb49",
		"0xcc69885fda6bcc1a4ace058b4a62bf5e179ea78fd58a1ccd71c22cc9b688792f",
		"0xd9d16d34ffb15ba3a3d852f0d403e2ce1d691fb54de27ac87cd2f993f3ec330f"}

	eden := NewEden(3)
	for i:=0; i < slotsAmount; i++ {
		if eden.delegateLocations[i].Hex() != slotDelegate[i] {
			t.Fatalf("slotDelegateLocation error")
		}
		if eden.expireLocations[i].Hex() != slotExpiration[i] {
			t.Fatalf("slotExpireLocation error")
		}
	}
}

func TestEdenContractAddr(t *testing.T)  {
	contractAddr := common.HexToAddress(mainnetEdenProxyAddress)
	if contractAddr == *new(common.Address) {
		t.Fatalf("mainnet")
	}
	ropstenContractAddr := common.HexToAddress(ropstenEdenProxyAddress)
	if ropstenContractAddr == *new(common.Address) {
		t.Fatalf("ropsten")
	}
}
