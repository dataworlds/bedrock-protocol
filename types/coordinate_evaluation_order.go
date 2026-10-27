package types

type CoordinateEvaluationOrder int32

const (
	CoordinateEvaluationOrderXYZ CoordinateEvaluationOrder = 0
	CoordinateEvaluationOrderXZY CoordinateEvaluationOrder = 1
	CoordinateEvaluationOrderYXZ CoordinateEvaluationOrder = 2
	CoordinateEvaluationOrderYZX CoordinateEvaluationOrder = 3
	CoordinateEvaluationOrderZXY CoordinateEvaluationOrder = 4
	CoordinateEvaluationOrderZYX CoordinateEvaluationOrder = 5
)
