package types

type ActorLink struct {
	TargetA                ActorUniqueID
	TargetB                ActorUniqueID
	Type                   ActorLinkType
	Immediate              bool
	PassengerInitiated     bool
	VehicleAngularVelocity float32
}
