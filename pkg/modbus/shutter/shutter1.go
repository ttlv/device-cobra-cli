package shutter

import (
	"device-cobra-cli/util/parse"
	"fmt"
	"github.com/goburrow/modbus"
	"github.com/sirupsen/logrus"
	"strings"
)

const (
	RTUShutter1Device = "/dev/ttyUSB1"
)

func GetShutter1Data(rtuDevice string) {
	handler := modbus.NewRTUClientHandler(rtuDevice)
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "N"
	handler.StopBits = 1
	handler.SlaveId = 1
	err := handler.Connect()
	if err != nil {
		logrus.Infof("访问串口%v异常", RTUShutter1Device)
		return
	}
	defer handler.Close()
	client := modbus.NewClient(handler)
	var lux, co2, pressure, temperature, humidity float64
	results, err := client.ReadHoldingRegisters(0, 23)
	s1 := strings.Replace(fmt.Sprintf("%v", results), "[", "", -1)
	s2 := strings.Replace(s1, "]", "", -1)
	splitS2 := strings.Split(s2, " ")
	ss1 := splitS2[0]
	ss2 := splitS2[1]
	if ss1 == "255" && ss2 == "255" {
		humidity = 0
		logrus.Info("访问设备异常,无法获取湿度数据")
	} else {
		humidity = parse.Hex2Dec(ss1, ss2) / 10
		logrus.Infof("成功获取湿度数据,当前的湿度: %v %v", humidity, "%RH")
	}
	// 温度
	ss3 := splitS2[2]
	ss4 := splitS2[3]
	if ss3 == "255" && ss4 == "255" {
		temperature = 0
		logrus.Info("访问设备异常,无法获取温度数据")
	} else {
		temperature = parse.Hex2Dec(ss3, ss4) / 10
		logrus.Info("成功获取温度数据,当前温度: %v ℃", temperature)
	}
	// 光强
	ss5 := splitS2[4]
	ss6 := splitS2[5]
	ss7 := splitS2[6]
	ss8 := splitS2[7]
	if ss5 == "255" && ss6 == "255" && ss7 == "255" && ss8 == "255" {
		lux = 0
		logrus.Info("访问设备异常,无法获取光强数据")
	} else {
		lux = parse.Hex2Dec(ss5, ss6, ss7, ss8)
		logrus.Info("成功获取光强数据, 当前光照强度: %v Lux", lux)
	}
	// CO2
	ss11 := splitS2[14]
	ss12 := splitS2[15]
	if ss11 == "255" && ss12 == "255" {
		co2 = 0
		logrus.Info("访问设备异常,无法获取二氧化碳浓度数据")
	} else {
		co2 = parse.Hex2Dec(ss11, ss12)
		logrus.Info("成功获取二氧化碳浓度数据,当前二氧化碳浓度: %v ppm", co2)
	}
	// 大气压强
	ss13 := splitS2[22]
	ss14 := splitS2[23]
	if ss13 == "255" && ss14 == "255" {
		pressure = 0
		logrus.Info("访问设备异常,无法获取大气压强数据")
	} else {
		pressure = parse.Hex2Dec(ss13, ss14) / 10
		logrus.Info("成功获取大气压强数据,当前大气压强: %v hpa", pressure)
	}
}
