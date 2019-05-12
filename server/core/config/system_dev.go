// +build !prod

package config

import (
	"os"
	"path"
)

const (

	//ServerConfigPath config of server listening and logs
	ServerConfigPath = "serverConfig/application.yaml"

	//ProjectHome env path to amco
	ProjectHome = "AMCO_HOME"

	//PermissionLogDir permission logs directory
	PermissionLogDir = 0755

	//PermissionlogFire  permisstion log files
	PermissionlogFire = 0644

	//Timeout  If the client has some requests but The server responded
	//slowly over the specified time AMCO will respond timeout message
	Timeout = 30

	FormatTimeLogPattern = "2006-01-02 15:04:05"
)

var (
	currentPath, _ = os.Getwd()

	InfoPath = path.Join(currentPath, "logs/info/info.log")

	DebugPath = path.Join(currentPath, "logs/debug/debugs.log")
)
