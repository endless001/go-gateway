package loadbalancer

type BalanceType int

const (
	BalanceTypeRandom BalanceType = iota
	BalanceTypeRound
	BalanceTypeWeight
	BalanceTypeHash
)

func LoadBanlanceFactory(balanceType BalanceType) LoadBalance {
	switch balanceType {
	case BalanceTypeRandom:
		return &RandomBalance{}
	case BalanceTypeRound:
		return &RoundBalance{}
	case BalanceTypeWeight:
		return &WeightBalance{}
	case BalanceTypeHash:
		return &HashBalance{}
	default:
		return &RandomBalance{}
	}
}
