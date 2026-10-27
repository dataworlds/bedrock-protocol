package types

type RandomDistributionType int32

const (
	RandomDistributionTypeSingleValued    RandomDistributionType = 0
	RandomDistributionTypeUniform         RandomDistributionType = 1
	RandomDistributionTypeGaussian        RandomDistributionType = 2
	RandomDistributionTypeInverseGaussian RandomDistributionType = 3
	RandomDistributionTypeFixedGrid       RandomDistributionType = 4
	RandomDistributionTypeJitteredGrid    RandomDistributionType = 5
	RandomDistributionTypeTriangle        RandomDistributionType = 6
)
