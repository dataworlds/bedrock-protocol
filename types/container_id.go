package types

type ContainerID int8

const (
	ContainerIDCONTAINERIDNONE           ContainerID = -1
	ContainerIDCONTAINERIDINVENTORY      ContainerID = 0
	ContainerIDCONTAINERIDFIRST          ContainerID = 1
	ContainerIDCONTAINERIDLAST           ContainerID = 100
	ContainerIDCONTAINERIDOFFHAND        ContainerID = 119
	ContainerIDCONTAINERIDARMOR          ContainerID = 120
	ContainerIDCONTAINERIDSELECTIONSLOTS ContainerID = 122
	ContainerIDCONTAINERIDPLAYERONLYUI   ContainerID = 124
	ContainerIDCONTAINERIDREGISTRY       ContainerID = 125
)
