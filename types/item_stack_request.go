package types

type ItemStackRequest struct {
	ClientRequestId ItemStackRequestId
	Actions         []struct {
		ControlValue                                        uint32
		ItemStackRequestTakeAction                          *ItemStackRequestTakeAction
		ItemStackRequestPlaceAction                         *ItemStackRequestPlaceAction
		ItemStackRequestSwapAction                          *ItemStackRequestSwapAction
		ItemStackRequestDropAction                          *ItemStackRequestDropAction
		ItemStackRequestDestroyAction                       *ItemStackRequestDestroyAction
		ItemStackRequestConsumeAction                       *ItemStackRequestConsumeAction
		ItemStackRequestCreateAction                        *ItemStackRequestCreateAction
		ItemStackRequestLabTableCombineAction               *ItemStackRequestLabTableCombineAction
		ItemStackRequestBeaconPaymentAction                 *ItemStackRequestBeaconPaymentAction
		ItemStackRequestMineBlockAction                     *ItemStackRequestMineBlockAction
		ItemStackRequestCraftRecipeAction                   *ItemStackRequestCraftRecipeAction
		ItemStackRequestCraftRecipeAutoAction               *ItemStackRequestCraftRecipeAutoAction
		ItemStackRequestCraftCreativeAction                 *ItemStackRequestCraftCreativeAction
		ItemStackRequestCraftRecipeOptionalAction           *ItemStackRequestCraftRecipeOptionalAction
		ItemStackRequestCraftRepairAndDisenchantAction      *ItemStackRequestCraftRepairAndDisenchantAction
		ItemStackRequestCraftLoomAction                     *ItemStackRequestCraftLoomAction
		ItemStackRequestReservedAction                      *ItemStackRequestReservedAction
		ItemStackRequestCraftNonImplementedDeprecatedAction *ItemStackRequestCraftNonImplementedDeprecatedAction
		ItemStackRequestCraftResultsDeprecatedAction        *ItemStackRequestCraftResultsDeprecatedAction
	}
	StringsToFilter       []string
	StringsToFilterOrigin TextProcessingEventOrigin
}
