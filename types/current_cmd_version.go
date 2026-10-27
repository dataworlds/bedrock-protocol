package types

type CurrentCmdVersion int32

const (
	CurrentCmdVersionInvalid                                                                                                               CurrentCmdVersion = -1
	CurrentCmdVersionInitial                                                                                                               CurrentCmdVersion = 1
	CurrentCmdVersionTpRotationClamping                                                                                                    CurrentCmdVersion = 2
	CurrentCmdVersionNewBedrockCmdSystem                                                                                                   CurrentCmdVersion = 3
	CurrentCmdVersionExecuteUsesVec3                                                                                                       CurrentCmdVersion = 4
	CurrentCmdVersionCloneFixes                                                                                                            CurrentCmdVersion = 5
	CurrentCmdVersionUpdateAquatic                                                                                                         CurrentCmdVersion = 6
	CurrentCmdVersionEntitySelectorUsesVec3                                                                                                CurrentCmdVersion = 7
	CurrentCmdVersionContainersDontDropItemsAnymore                                                                                        CurrentCmdVersion = 8
	CurrentCmdVersionFiltersObeyDimensions                                                                                                 CurrentCmdVersion = 9
	CurrentCmdVersionExecuteAndBlockCommandAndSelfSelectorFixes                                                                            CurrentCmdVersion = 10
	CurrentCmdVersionInstantEffectsUseTicks                                                                                                CurrentCmdVersion = 11
	CurrentCmdVersionDontRegisterBrokenFunctionCommands                                                                                    CurrentCmdVersion = 12
	CurrentCmdVersionClearSpawnPointCommand                                                                                                CurrentCmdVersion = 13
	CurrentCmdVersionCloneAndTeleportRotationFixes                                                                                         CurrentCmdVersion = 14
	CurrentCmdVersionTeleportDimensionFixes                                                                                                CurrentCmdVersion = 15
	CurrentCmdVersionCloneUpdateBlockAndTimeFixes                                                                                          CurrentCmdVersion = 16
	CurrentCmdVersionCloneIntersectFix                                                                                                     CurrentCmdVersion = 17
	CurrentCmdVersionFunctionExecuteOrderAndChestSlotFix                                                                                   CurrentCmdVersion = 18
	CurrentCmdVersionNonTickingAreasNoLongerConsideredLoaded                                                                               CurrentCmdVersion = 19
	CurrentCmdVersionSpreadplayersHazardAndResolvePlayerByNameFix                                                                          CurrentCmdVersion = 20
	CurrentCmdVersionNewExecuteCommandSyntaxExperimentAndChestLootTableFixAndTeleportFacingVerticalUnclampedAndLocateBiomeAndFeatureMerged CurrentCmdVersion = 21
	CurrentCmdVersionWaterloggingAddedToStructureCommand                                                                                   CurrentCmdVersion = 22
	CurrentCmdVersionSelectorDistanceFilteredAndRelativeRotationFix                                                                        CurrentCmdVersion = 23
	CurrentCmdVersionNewSummonCommandAddedRotationOptionsAndBubbleColumnCloneFixAndExecuteInDimensionTeleportFixAndNewExecuteRotationFix   CurrentCmdVersion = 24
	CurrentCmdVersionNewExecuteCommandReleaseEnchantCommandLevelFixAndHasItemDataFixAndCommandDeferral                                     CurrentCmdVersion = 25
	CurrentCmdVersionExecuteIfScoreFixes                                                                                                   CurrentCmdVersion = 26
	CurrentCmdVersionReplaceItemAndLootReplaceBlockCommandsDoNotPlaceItemsIntoCauldronsFix                                                 CurrentCmdVersion = 27
	CurrentCmdVersionChangesToCommandOriginRotation                                                                                        CurrentCmdVersion = 28
	CurrentCmdVersionRemoveAuxValueParameterFromBlockCommands                                                                              CurrentCmdVersion = 29
	CurrentCmdVersionVolumeSelectorFixes                                                                                                   CurrentCmdVersion = 30
	CurrentCmdVersionEnableSummonRotation                                                                                                  CurrentCmdVersion = 31
	CurrentCmdVersionSummonCommandDefaultRotation                                                                                          CurrentCmdVersion = 32
	CurrentCmdVersionPositionalDimensionFiltering                                                                                          CurrentCmdVersion = 33
	CurrentCmdVersionCommandSelectorHasItemFilterNoLongerCallsSameItemFunction                                                             CurrentCmdVersion = 34
	CurrentCmdVersionAgentSweepingBlockTest                                                                                                CurrentCmdVersion = 34
	CurrentCmdVersionBlockStateEquals                                                                                                      CurrentCmdVersion = 35
	CurrentCmdVersionCommandPositionFix                                                                                                    CurrentCmdVersion = 35
	CurrentCmdVersionCommandSelectorHasItemFilterUsesDataAsDamageForSelectingDamageableItems                                               CurrentCmdVersion = 36
	CurrentCmdVersionExecuteDetectConditionSubcommandNotAllowNonLoadedBlocks                                                               CurrentCmdVersion = 37
	CurrentCmdVersionRemoveSuicideKeyword                                                                                                  CurrentCmdVersion = 38
	CurrentCmdVersionCloneContainerBlockEntityRemovalFix                                                                                   CurrentCmdVersion = 39
	CurrentCmdVersionStopSoundMusicFix                                                                                                     CurrentCmdVersion = 40
	CurrentCmdVersionSpreadPlayersStuckInGroundFixAndMaxHeightParameter                                                                    CurrentCmdVersion = 41
	CurrentCmdVersionLocateStructureOutput                                                                                                 CurrentCmdVersion = 42
	CurrentCmdVersionPostBlockFlattening                                                                                                   CurrentCmdVersion = 43
	CurrentCmdVersionTestForBlockCommandDoesNotIgnoreBlockState                                                                            CurrentCmdVersion = 44
	CurrentCmdVersionCount                                                                                                                 CurrentCmdVersion = 56
	CurrentCmdVersionLatest                                                                                                                CurrentCmdVersion = 55
)
