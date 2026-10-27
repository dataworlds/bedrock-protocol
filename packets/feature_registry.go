package packets

import "github.com/dataworlds/bedrock-protocol/types"

const FeatureRegistryID uint32 = 191

type FeatureRegistry struct {
	FeaturesDataList []types.FeatureRegistryFeatureBinaryJsonFormat
}
