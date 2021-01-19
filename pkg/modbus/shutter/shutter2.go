package shutter

import (
	"device-cobra-cli/util/parse"
	"fmt"
	"github.com/goburrow/modbus"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

const (
	RTUShutter2Device = "/dev/ttyUSB2"
)

func GetShutter2Data(rtuDevice string){
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
		var pm2point5, pm10, noise float64
		results, err := client.ReadHoldingRegisters(20, 2)
		if err != nil || results == nil {
			logrus.Info("访问百叶窗1设备异常,无法获取PM2.5和PM10的数据,尝试重新访问设备")
		}else{
			s1 := strings.Replace(fmt.Sprintf("%v", results), "[", "", -1)
			s2 := strings.Replace(s1, "]", "", -1)
			ss1 := strings.Split(s2, " ")[0]
			ss2 := strings.Split(s2, " ")[1]
			ss3 := strings.Split(s2, " ")[2]
			ss4 := strings.Split(s2, " ")[3]
			if  ss1=="255" || ss2=="255" || ss3=="255" || ss4=="255"{
				logrus.Info("访问百叶窗1设备异常,无法获取pm2.5与pm10的数据,尝试重新访问设备")
			}else{
				pm2point5 = parse.Hex2Dec(ss1, ss2)
				logrus.Infof("PM2.5: %v ug/m3【PM2.5】",pm2point5)
				pm10 = parse.Hex2Dec(ss3, ss4)
				logrus.Infof("PM10: %v ug/m3【PM10】",pm10)
			}
		}

		results, err = client.ReadHoldingRegisters(4, 1)
		if err != nil || results == nil {
			logrus.Info("访问百叶窗2设备异常,无法获取噪音数据,尝试重新访问设备")
		}else{
			s1 := strings.Replace(fmt.Sprintf("%v", results), "[", "", -1)
			s2 := strings.Replace(s1, "]", "", -1)
			ss1 := strings.Split(s2, " ")[0]
			ss2 := strings.Split(s2, " ")[1]
			if  ss1=="255" || ss2=="255"{
				logrus.Info("访问百叶窗1设备异常,无法噪音数据,尝试重新访问设备")
			}else{
				noise = parse.Hex2Dec(ss1, ss2)
				logrus.Infof("noise: %v db【噪音】",noise/10)
			}
		}
		time.Sleep(time.Second * 1)
	}

}

