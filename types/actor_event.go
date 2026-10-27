package types

type ActorEvent uint8

const (
	ActorEventNONE                             ActorEvent = 0
	ActorEventJUMP                             ActorEvent = 1
	ActorEventHURT                             ActorEvent = 2
	ActorEventDEATH                            ActorEvent = 3
	ActorEventSTARTATTACKING                   ActorEvent = 4
	ActorEventSTOPATTACKING                    ActorEvent = 5
	ActorEventTAMINGFAILED                     ActorEvent = 6
	ActorEventTAMINGSUCCEEDED                  ActorEvent = 7
	ActorEventSHAKEWETNESS                     ActorEvent = 8
	ActorEventEATGRASS                         ActorEvent = 10
	ActorEventFISHHOOKBUBBLE                   ActorEvent = 11
	ActorEventFISHHOOKFISHPOS                  ActorEvent = 12
	ActorEventFISHHOOKHOOKTIME                 ActorEvent = 13
	ActorEventFISHHOOKTEASE                    ActorEvent = 14
	ActorEventSQUIDFLEEING                     ActorEvent = 15
	ActorEventZOMBIECONVERTING                 ActorEvent = 16
	ActorEventPLAYAMBIENT                      ActorEvent = 17
	ActorEventSPAWNALIVE                       ActorEvent = 18
	ActorEventSTARTOFFERFLOWER                 ActorEvent = 19
	ActorEventSTOPOFFERFLOWER                  ActorEvent = 20
	ActorEventLOVEHEARTS                       ActorEvent = 21
	ActorEventVILLAGERANGRY                    ActorEvent = 22
	ActorEventVILLAGERHAPPY                    ActorEvent = 23
	ActorEventWITCHHATMAGIC                    ActorEvent = 24
	ActorEventFIREWORKSEXPLODE                 ActorEvent = 25
	ActorEventINLOVEHEARTS                     ActorEvent = 26
	ActorEventSILVERFISHMERGEANIM              ActorEvent = 27
	ActorEventGUARDIANATTACKSOUND              ActorEvent = 28
	ActorEventDRINKPOTION                      ActorEvent = 29
	ActorEventTHROWPOTION                      ActorEvent = 30
	ActorEventPRIMETNTCART                     ActorEvent = 31
	ActorEventPRIMECREEPER                     ActorEvent = 32
	ActorEventAIRSUPPLY                        ActorEvent = 33
	ActorEventDEPRECATEDADDPLAYERLEVELS        ActorEvent = 34
	ActorEventGUARDIANMININGFATIGUE            ActorEvent = 35
	ActorEventAGENTSWINGARM                    ActorEvent = 36
	ActorEventDRAGONSTARTDEATHANIM             ActorEvent = 37
	ActorEventGROUNDDUST                       ActorEvent = 38
	ActorEventSHAKE                            ActorEvent = 39
	ActorEventFEED                             ActorEvent = 57
	ActorEventBABYAGE                          ActorEvent = 60
	ActorEventINSTANTDEATH                     ActorEvent = 61
	ActorEventNOTIFYTRADE                      ActorEvent = 62
	ActorEventLEASHDESTROYED                   ActorEvent = 63
	ActorEventCARAVANUPDATED                   ActorEvent = 64
	ActorEventTALISMANACTIVATE                 ActorEvent = 65
	ActorEventDEPRECATEDUPDATESTRUCTUREFEATURE ActorEvent = 66
	ActorEventPLAYERSPAWNEDMOB                 ActorEvent = 67
	ActorEventPUKE                             ActorEvent = 68
	ActorEventUPDATESTACKSIZE                  ActorEvent = 69
	ActorEventSTARTSWIMMING                    ActorEvent = 70
	ActorEventBALLOONPOP                       ActorEvent = 71
	ActorEventTREASUREHUNT                     ActorEvent = 72
	ActorEventSUMMONAGENT                      ActorEvent = 73
	ActorEventFINISHEDCHARGINGITEM             ActorEvent = 74
	ActorEventACTORGROWUP                      ActorEvent = 76
	ActorEventVIBRATIONDETECTED                ActorEvent = 77
	ActorEventDRINKMILK                        ActorEvent = 78
	ActorEventSHAKEWETNESSSTOP                 ActorEvent = 79
	ActorEventKINETICDAMAGEDEALT               ActorEvent = 80
	ActorEventHURTWITHOUTRECEIVINGDAMAGE       ActorEvent = 81
)
