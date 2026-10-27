package packets

const NetworkStackLatencyID uint32 = 115

type NetworkStackLatency struct {
	CreationTime uint64
	IsFromServer bool
}
