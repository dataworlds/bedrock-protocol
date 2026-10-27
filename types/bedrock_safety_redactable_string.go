package types

type BedrockSafetyRedactableString struct {
	Unredacted string
	Redacted   *string
}
