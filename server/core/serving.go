package core

import (
	"KafkaMessageQ-API/server/core/config"
	"KafkaMessageQ-API/server/core/router"
	"KafkaMessageQ-API/server/core/structs/serverConfig"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"

	"gopkg.in/yaml.v2"
)

var (
	serverConf        serverConfig.ServeConfig
	applicationConfig serverConfig.Application
)

var pathAMCOHome string

func initServer() error {
	serverConfigYaml, err := readConfigFile()
	if err != nil {
		return err
	}

	if err = setConfig(serverConfigYaml); err != nil {
		return err
	}

	if err = createLog(); err != nil {
		return err
	}

	return nil
}

//Start run web server
func Start() error {

	if err := initServer(); err != nil {
		return err
	}

	listen := applicationConfig.Host + ":" + strconv.Itoa(applicationConfig.Port)
	http.Handle("/", router.ServeHandle())
	return http.ListenAndServe(listen, nil)
}

func readConfigFile() ([]byte, error) {

	//get value of Env for set project path
	pathAMCOHome = os.Getenv(config.ProjectHome)
	if pathAMCOHome == "" {
		return []byte(""), errors.New("AMCO_HOME path isn't define on environment variable")
	}

	//read system config yaml file
	serverConfigYaml, err := ioutil.ReadFile(path.Join(pathAMCOHome, config.ServerConfigPath))
	if err != nil {
		return []byte(""), err
	}

	return serverConfigYaml, nil
}

func setConfig(serverConfigYaml []byte) error {

	if err := yaml.Unmarshal(serverConfigYaml, &serverConf); err != nil {
		return err
	}

	//get value of development by specification environment
	for _, v := range serverConf.Serving.Application {
		if serverConf.Serving.Spec == v.Spec {
			applicationConfig = v
		}
	}

	return nil
}

func createLog() error {

	const (
		prefix      = "logs"
		debugLogDir = "debug"
		infoLogDir  = "info"

		debugFile = "debug.log"
		infoFile  = "info.log"
	)

	var (
		debugFilePath = path.Join(prefix, debugLogDir, debugFile)
		infoFilePath  = path.Join(prefix, debugLogDir, infoFile)
	)

	logDir := path.Join(pathAMCOHome, prefix)
	if _, err := os.Stat(logDir); err != nil {
		os.MkdirAll(logDir, config.PermissionLogDir)
	}

	createDebugLogDirPath := path.Join(logDir, debugLogDir)
	if _, err := os.Stat(createDebugLogDirPath); err != nil {
		os.MkdirAll(createDebugLogDirPath, config.PermissionLogDir)

	}

	createInfoLogDirPath := path.Join(logDir, infoLogDir)
	if _, err := os.Stat(createInfoLogDirPath); err != nil {
		os.MkdirAll(createInfoLogDirPath, config.PermissionLogDir)

	}

	//create file logs
	if config.DebugPath == "" {
		_, err := os.Create(debugFilePath)
		if err != nil {
			return err
		}
	} else {
		if _, err := os.Stat(config.DebugPath); err != nil {
			if _, err := os.Create(config.DebugPath); err != nil {
				return err
			}
		}
	}

	if config.InfoPath == "" {
		_, err := os.Create(infoFilePath)
		if err != nil {
			return err
		}
	} else {
		if _, err := os.Stat(config.InfoPath); err != nil {
			if _, err := os.Create(config.InfoPath); err != nil {
				return err
			}
		}
	}

	return nil
}
