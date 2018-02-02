package tplink

type SmartDev interface{
	RawQuery(devIP string, cmd string, args string) ([]byte, error)
	Features() ([]string, error)
	SysInfo() (map[string]interface{}, error)
	Model() (string, error)
	Alias() (string, error)
	SetAlias(string) error

}

