package snow

import (
	"device-cobra-cli/util/parse"
	"fmt"
	"github.com/goburrow/modbus"
	"strings"
	"time"
	"github.com/sirupsen/logrus"
)

const (
	RTUSnowDevice = "/dev/ttyUSB3"
)

func GetSnowData(rtuDevice string){
	handler := modbus.NewRTUClientHandler(rtuDevice)
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "N"
	handler.StopBits = 1
	handler.SlaveId = 1
	err := handler.Connect()
	if err != nil {
		logrus.Infof("访问串口%v异常",rtuDevice)
	}
	defer handler.Close()
	client := modbus.NewClient(handler)
	for i:=1;i<=5;i++{
		results, err := client.ReadHoldingRegisters(22, 1)
		if err != nil || results == nil {
			logrus.Info("访问雨雪设备异常,无法获取雨雪数据,尝试重新访问设备")
			continue
		}
		var snow float64
		s1 := strings.Replace(fmt.Sprintf("%v", results), "[", "", -1)
		s2 := strings.Replace(s1, "]", "", -1)
		ss1 := strings.Split(s2, " ")[0]
		ss2 := strings.Split(s2, " ")[1]
		snow = parse.Hex2Dec(ss1, ss2)
		if snow==0{
			logrus.Infof("snow: %v【当前无雨雪】",snow)
		}else if snow==1{
			logrus.Infof("snow: %v【当前有雨雪】",snow)
		}else{
			logrus.Infof("snow: %v【无法从雨雪设备获取有效的数据】",snow)
		}
		time.Sleep(time.Second * 1)
		if i==5{
			logrus.Infof("串口:%v以及设备一切正常",rtuDevice)
		}
	}
}
