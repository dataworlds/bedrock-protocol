package types

type GameRule struct {
	RuleName          string
	RuleCanBeModified bool
	RuleValue         struct {
		ControlValue uint32
		Variant0     *any
		Variant1     *bool
		Variant2     *int32
		Variant3     *float32
	}
}
