package types

type PlayerDied struct {
	InstigatorActorID    int32
	InstigatorMobVariant int32
	DamageSource         int32
	DiedInRaid           bool
}
