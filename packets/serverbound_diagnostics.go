package packets

import "github.com/dataworlds/bedrock-protocol/types"

const ServerboundDiagnosticsID uint32 = 315

type ServerboundDiagnostics struct {
	AvgFps                    float32
	AvgServerSimTickTimeMS    float32
	AvgClientSimTickTimeMS    float32
	AvgBeginFrameTimeMS       float32
	AvgInputTimeMS            float32
	AvgRenderTimeMS           float32
	AvgEndFrameTimeMS         float32
	AvgRemainderTimePercent   float32
	AvgUnaccountedTimePercent float32
	MemoryCategoryValues      []types.MemoryCategoryCounter
	EntityDiagnostics         []types.EntityDiagnosticTimingInfo
	SystemDiagnostics         []types.SystemDiagnosticTimingInfo
	SystemCategories          *[]types.SystemCategory
	WhiskerScopes             []types.WhiskerScopeDataSummary
}
