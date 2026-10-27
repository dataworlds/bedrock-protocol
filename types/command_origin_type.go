package types

type CommandOriginType uint8

const (
	CommandOriginTypePlayer                   CommandOriginType = 0
	CommandOriginTypeCommandBlock             CommandOriginType = 1
	CommandOriginTypeMinecartCommandBlock     CommandOriginType = 2
	CommandOriginTypeDevConsole               CommandOriginType = 3
	CommandOriginTypeTest                     CommandOriginType = 4
	CommandOriginTypeAutomationPlayer         CommandOriginType = 5
	CommandOriginTypeClientAutomation         CommandOriginType = 6
	CommandOriginTypeDedicatedServer          CommandOriginType = 7
	CommandOriginTypeEntity                   CommandOriginType = 8
	CommandOriginTypeVirtual                  CommandOriginType = 9
	CommandOriginTypeGameArgument             CommandOriginType = 10
	CommandOriginTypeEntityServer             CommandOriginType = 11
	CommandOriginTypePrecompiled              CommandOriginType = 12
	CommandOriginTypeGameDirectorEntityServer CommandOriginType = 13
	CommandOriginTypeScripting                CommandOriginType = 14
	CommandOriginTypeExecuteContext           CommandOriginType = 15
)
