package types

type ServerWaypointPayload struct {
	UpdateFlag              uint32
	IsVisible               *bool
	WorldPosition           *WorldPosition
	TexturePath             *string
	IconSize                *Vec2
	Color                   *Color
	ClientPositionAuthority *bool
	ActorUniqueID           *ActorUniqueID
}
