package types

type LocatorBarWaypointPayload struct {
	GroupHandle           WaypointGroupWaypointHandle
	ServerWaypointPayload ServerWaypointPayload
	ActionFlag            ServerWaypointGroupAction
}
