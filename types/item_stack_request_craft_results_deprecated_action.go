package types

type ItemStackRequestCraftResultsDeprecatedAction struct {
	ActionType   ItemStackRequestActionType
	CraftResults []ItemStackRequestNetworkItemInstanceDescriptor
	NumCrafts    uint8
}
