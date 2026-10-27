package types

type MobKilled struct {
	InstigatorActorID         int64
	TargetActorID             int64
	InstigatorSChildActorType ActorType
	DamageSource              int32
	TradeTier                 int32
	TraderName                string
}
