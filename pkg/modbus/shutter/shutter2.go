package shutter

import (
	"device-cobra-cli/util/parse"
	"fmt"
	"github.com/goburrow/modbus"
	"github.com/sirupsen/logrus"
	"strings"
)

const (
	RTUShutter2Device = "/dev/ttyUSB2"
)

func GetShutter2Data(rtuDevice string) {
	handler := modbus.NewRTUClientHandler(rtuDevice)
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "N"
	handler.StopBits = 1
	handler.SlaveId = 1
	err := handler.Connect()
	defer handler.Close()
	if err != nil {
		logrus.Infof("访问串口%v异常", RTUShutter2Device)
		return
	}
	var pm2point5, pm10, noise float64
	client := modbus.NewClient(handler)
	results, err := client.ReadHoldingRegisters(4, 18)
	s1 := strings.Replace(fmt.Sprintf("%v", results), "[", "", -1)
	s2 := strings.Replace(s1, "]", "", -1)
	splitS2 := strings.Split(s2, " ")
	ss1 := splitS2[0]
	ss2 := splitS2[1]
	if ss1 == "255" && ss2 == "255" {
		noise = 0
		logrus.Info("访问设备异常,无法获取噪音数据")
	} else {
		noise = parse.Hex2Dec(ss1, ss2) / 10
		logrus.Infof("成功获取噪音数据,当前的噪音: %v %v", noise, "DB")
	}
	ss3 := strings.Split(s2, " ")[40]
	ss4 := strings.Split(s2, " ")[41]
	ss5 := strings.Split(s2, " ")[42]
	ss6 := strings.Split(s2, " ")[43]
	if ss3 == "255" && ss4 == "255" {
		pm2point5 = 0
		logrus.Info("访问设备异常,无法获取PM2.5数据")
	} else {
		pm2point5 = parse.Hex2Dec(ss1, ss2)
		logrus.Infof("成功获取PM2.5数据,当前的PM2.5浓度: %v %v", pm2point5, "ug/m3")
	}
	if ss5 == "255" && ss6 == "255" {
		pm10 = 0
		logrus.Info("访问设备异常,无法获取PM10数据")
	} else {
		pm10 = parse.Hex2Dec(ss1, ss2)
		logrus.Infof("成功获取PM2.5数据,当前的PM10浓度: %v %v", pm10, "ug/m3")
	}
}
