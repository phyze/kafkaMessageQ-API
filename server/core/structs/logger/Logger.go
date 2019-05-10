package logger

type Logger struct {
	LogType   string `json:"logType"`
	TimeStamp string `json:"timeStamp"`
	Data      string `json:"data"`
	Version   string `json:"version"`
	Release   string `json:"release"`
	IAM       string `json:"iam"`
}


func (lg *Logger) ToMap() map[string]interface{} {
	mapLog := make(map[string]interface{})
	mapLog["logType"] = lg.LogType
	mapLog["timeStamp"] = lg.TimeStamp
	mapLog["data"] = lg.Data
	mapLog["version"] = lg.Version
	mapLog["iam"] = lg.IAM
	mapLog["release"] = lg.Release
	return mapLog
}