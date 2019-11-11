package model

import (
//	"gopkg.in/mgo.v2/bson"
	"strconv"
	"fmt"
)

type AccessOptionDf1 struct {
	BaudRate	int    `bson:"baudRate" json:"baudrate"`
	DataBits	int    `bson:"dataBits" json:"databits"`
	StopBits	int    `bson:"stopBits" json:"stopbits"`
	Parity		int `bson:"parity" json:"parity"`
	SlaveId		int	`bson:"slaveId" json:"slaveid"`
	Mode		string `bson:"mode" json:"mode"`
	Sip			string	`bson:"sip" json:"sip"`
	Protocol    string        `bson:"protocol" json:"protocol"`
}

type ReadAccessDataDf1 struct {
	DataName string `bson:"dataName" json:"dataName"`
	DataType string `bson:"dataType" json:"dataType"`
	DataAddr string `bson:"dataAddr" json:"dataAddr"`
}

type WriteAccessDataDf1 struct {
	DataName string `bson:"dataName" json:"dataName"`
	DataType string `bson:"dataType" json:"dataType"`
	DataAddr string `bson:"dataAddr" json:"dataAddr"`
}

type ConfigDf1 struct {
	DeviceName    string        `bson:"deviceName" json:"devicename"`
	ConfigName    string        `bson:"configName" json:"configname"`
	Endpoint      string        `bson:"endpoint" json:"endpoint"`
	Protocol      string        `bson:"protocol" json:"protocol"`
	AccessOptions AccessOptionDf1  `bson:"accessOptions" json:"accessoptions"`
	ReadAccessDatas []ReadAccessDataDf1 `bson:"readAccessDatas" json:"readaccessdatas"`
	WriteAccessDatas []WriteAccessDataDf1 `bson:"writeAccessDatas" json:"writeaccessdatas"`
}

/*
func FindByNameDf1(name string) (*ConfigDf1, error) {
	c := ConfigDf1{}
	err := ConfigCollection().Find(bson.M{
		"deviceName": name,
	}).One(&c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
*/
// args: device mode speed databits parity stopbits 
func GetServerArgs(config *ConfigDf1)(args[]string, err error){
	args = make([]string, 6)
	args[0] = config.Endpoint	
	args[1] = config.AccessOptions.Mode
	args[2] = strconv.Itoa(config.AccessOptions.BaudRate)
	args[3] = strconv.Itoa(config.AccessOptions.DataBits)
	args[4] = strconv.Itoa(config.AccessOptions.Parity)
	args[5] = strconv.Itoa(config.AccessOptions.StopBits)
	
	var flag uint8 = 0
	errStr := "DF1 server config: "
	for i := 0; i< len(args); i++ {
		if len(args[i]) == 0 {
			errStr += "param"+ strconv.Itoa(i) +","	
			flag = 1
		}
	}
	if flag == 0{
		return args, nil
	} else {
		errStr += " are null"
		err = fmt.Errorf(errStr)
	}
	
	return nil, err 
}

// args: sip addr=value 
func GetClientArgs(config *ConfigDf1)(args[]string, err error){
	args = make([]string, 1)
	args[0] = config.AccessOptions.Sip
	
	errStr := "DF1 client config:"
	if len(args[0]) == 0 {
		errStr += "param" + "0" + " are null"
		return nil, fmt.Errorf(errStr)
	}
	return args, nil
}

