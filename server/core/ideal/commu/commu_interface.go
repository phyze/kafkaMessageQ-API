package commu

type CommunicationInterface interface {
	Validate() error
	GetFieldsName() []string
	IsCompatible(*map[string]interface{}) bool
}
