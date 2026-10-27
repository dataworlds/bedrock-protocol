package types

type MolangVersion int16

const (
	MolangVersionInvalid                                MolangVersion = -1
	MolangVersionBeforeVersioning                       MolangVersion = 0
	MolangVersionInitial                                MolangVersion = 1
	MolangVersionFixedItemRemainingUseDurationQuery     MolangVersion = 2
	MolangVersionExpressionErrorMessages                MolangVersion = 3
	MolangVersionUnexpectedOperatorErrors               MolangVersion = 4
	MolangVersionConditionalOperatorAssociativity       MolangVersion = 5
	MolangVersionComparisonAndLogicalOperatorPrecedence MolangVersion = 6
	MolangVersionDivideByNegativeValue                  MolangVersion = 7
	MolangVersionFixedCapeFlapAmountQuery               MolangVersion = 8
	MolangVersionQueryBlockPropertyRenamedToState       MolangVersion = 9
	MolangVersionDeprecateOldBlockQueryNames            MolangVersion = 10
	MolangVersionDeprecatedSnifferAndCamelQueries       MolangVersion = 11
	MolangVersionLeafSupportingInFirstSolidBlockBelow   MolangVersion = 12
	MolangVersionNumValidVersions                       MolangVersion = 14
	MolangVersionLatest                                 MolangVersion = 13
	MolangVersionHardcodedMolang                        MolangVersion = 13
)
