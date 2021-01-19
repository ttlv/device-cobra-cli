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
	RTUShutter1Device = "/dev/ttyUSB1"
)

func GetShutter1Data(rtuDevice string){
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
		var lux, co2, pressure, temperature, humidity float64
		results, err := client.ReadHoldingRegisters(2, 2)
		if err != nil || results == nil {
			logrus.Info("访问百叶窗1设备异常,无法获取光强数据,尝试重新访问设备")
		}else {
			s1 := strings.Replace(fmt.Sprintf("%v", results), "[", "", -1)
			s2 := strings.Replace(s1, "]", "", -1)
			ss1 := strings.Split(s2, " ")[0]
			ss2 := strings.Split(s2, " ")[1]
			ss3 := strings.Split(s2, " ")[2]
			ss4 := strings.Split(s2, " ")[3]
			lux = parse.Hex2Dec(ss1, ss2, ss3, ss4)
			if ss1=="255" || ss2=="255" || ss3=="255" || ss4=="255"{
				logrus.Info("访问百叶窗1设备异常,无法获取光强数据,尝试重新访问设备")
			}else{
				logrus.Infof("lux: %v Lux【光照强度】",lux)
			}
		}

		results, err = client.ReadHoldingRegisters(7, 1)
		if err != nil || results == nil {
			logrus.Info("访问百叶窗1设备异常,无法获取二氧化碳数据,尝试重新访问设备")
		}else{
			s1 := strings.Replace(fmt.Sprintf("%v", results), "[", "", -1)
			s2 := strings.Replace(s1, "]", "", -1)
			ss1 := strings.Split(s2, " ")[0]
			ss2 := strings.Split(s2, " ")[1]
			if  ss1=="255" || ss2=="255"{
				logrus.Info("访问百叶窗1设备异常,无法获取二氧化碳数据,尝试重新访问设备")
			}else{
				co2 = parse.Hex2Dec(ss1, ss2)
				logrus.Infof("co2: %v ppm【二氧化碳】",co2)
			}
		}

		results, err = client.ReadHoldingRegisters(11, 1)
		if err != nil || results == nil {
			logrus.Info("访问百叶窗1设备异常,无法获取大气压强数据,尝试重新访问设备")
		}else{
			s1 := strings.Replace(fmt.Sprintf("%v", results), "[", "", -1)
			s2 := strings.Replace(s1, "]", "", -1)
			ss1 := strings.Split(s2, " ")[0]
			ss2 := strings.Split(s2, " ")[1]
			if  ss1=="255" || ss2=="255"{
				logrus.Info("访问百叶窗1设备异常,无法获取大气压强数据,尝试重新访问设备")
			}else{
				pressure = parse.Hex2Dec(ss1, ss2)
				logrus.Infof("atmospheric pressure: %v hpa【大气压】",pressure/10)
			}
		}

		results, err = client.ReadHoldingRegisters(0, 2)
		if err != nil || results == nil {
			logrus.Info("访问百叶窗1设备异常,无法获取温湿度数据,尝试重新访问设备")
		}else{
			s1 := strings.Replace(fmt.Sprintf("%v", results), "[", "", -1)
			s2 := strings.Replace(s1, "]", "", -1)
			ss1 := strings.Split(s2, " ")[0]
			ss2 := strings.Split(s2, " ")[1]
			ss3 := strings.Split(s2, " ")[2]
			ss4 := strings.Split(s2, " ")[3]
			if  ss1=="255" || ss2=="255" || ss3=="255" || ss4=="255"{
				logrus.Info("访问百叶窗1设备异常,无法获取温湿度数据,尝试重新访问设备")
			}else{
				temperature = parse.Hex2Dec(ss1, ss2)
				logrus.Infof("temperature: %v C【温度】",temperature/10)
				humidity = parse.Hex2Dec(ss3, ss4)
				logrus.Infof("humidity: %v %v【湿度】",humidity/10,"%RH")
			}
		}
		time.Sleep(time.Second * 1)
	}

}