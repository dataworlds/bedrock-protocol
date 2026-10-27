package types

type SimulationType uint8

const (
	SimulationTypeGame    SimulationType = 0
	SimulationTypeEditor  SimulationType = 1
	SimulationTypeTest    SimulationType = 2
	SimulationTypeINVALID SimulationType = 3
)
