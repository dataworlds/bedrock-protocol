package packets

import "github.com/dataworlds/bedrock-protocol/types"

const LocatorBarID uint32 = 341

type LocatorBar struct {
	Waypoints []types.LocatorBarWaypointPayload
}
