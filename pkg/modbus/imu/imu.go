package imu

import (
	"fmt"
	"github.com/goburrow/modbus"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

const (
	RTUIMUDevice = "/dev/ttyUSB0"
)

func GetIMUData(rtuDevice string) {
	handler := modbus.NewRTUClientHandler(rtuDevice)
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "N"
	handler.StopBits = 1
	handler.SlaveId = 80
	err := handler.Connect()
	if err != nil {
		logrus.Infof("访问串口%v异常", RTUIMUDevice)
		return
	}
	defer handler.Close()
	client := modbus.NewClient(handler)
	results, err := client.ReadHoldingRegisters(52, 17)
	s1 := strings.Replace(fmt.Sprintf("%v", results), "[", "", -1)
	s2 := strings.Replace(s1, "]", "", -1)
	splitS2 := strings.Split(s2, " ")
	// acceleration
	ss1 := splitS2[0]
	ss2 := splitS2[1]
	ss3 := splitS2[2]
	ss4 := splitS2[3]
	ss5 := splitS2[4]
	ss6 := splitS2[5]
	if ss1 == "255" && ss2 == "255" && ss3 == "255" && ss4 == "255" && ss5 == "255" && ss6 == "255" {
		logrus.Info("访问设备异常,无法获取加速度数据")
	} else {
		axh, _ := strconv.Atoi(ss1)
		axl, _ := strconv.Atoi(ss2)
		ayh, _ := strconv.Atoi(ss3)
		ayl, _ := strconv.Atoi(ss4)
		azh, _ := strconv.Atoi(ss5)
		azl, _ := strconv.Atoi(ss6)
		k := 16.0
		accX := float64(axh<<8|axl) / 32768.0 * k
		accY := float64(ayh<<8|ayl) / 32768.0 * k
		accZ := float64(azh<<8|azl) / 32768.0 * k
		if accX >= k {
			accX -= 2 * k
		}
		if accY >= k {
			accY -= 2 * k
		}
		if accZ >= k {
			accZ -= 2 * k
		}
		logrus.Infof("accX: %v m/s2", accX)
		logrus.Infof("accY: %v m/s2", accY)
		logrus.Infof("accZ: %v m/s2", accZ)
	}
	// angularVelocity
	ss7 := splitS2[6]
	ss8 := splitS2[7]
	ss9 := splitS2[8]
	ss10 := splitS2[9]
	ss11 := splitS2[10]
	ss12 := splitS2[11]
	if ss7 == "255" && ss8 == "255" && ss9 == "255" && ss10 == "255" && ss11 == "255" && ss12 == "255" {
		logrus.Info("访问设备异常,无法获取角速度数据")
	} else {
		wxh, _ := strconv.Atoi(ss7)
		wxl, _ := strconv.Atoi(ss8)
		wyh, _ := strconv.Atoi(ss9)
		wyl, _ := strconv.Atoi(ss10)
		wzh, _ := strconv.Atoi(ss11)
		wzl, _ := strconv.Atoi(ss12)
		k := 2000.0
		wX := float64(wxh<<8|wxl) / 32768.0 * k
		wY := float64(wyh<<8|wyl) / 32768.0 * k
		wZ := float64(wzh<<8|wzl) / 32768.0 * k
		if wX >= k {
			wX -= 2 * k
		}
		if wY >= k {
			wY -= 2 * k
		}
		if wZ >= k {
			wZ -= 2 * k
		}
		logrus.Infof("wX: %v °/s", wX)
		logrus.Infof("wY: %v °/s", wY)
		logrus.Infof("wX: %v °/s", wZ)
	}
	// magnetic
	ss13 := splitS2[12]
	ss14 := splitS2[13]
	ss15 := splitS2[14]
	ss16 := splitS2[15]
	ss17 := splitS2[16]
	ss18 := splitS2[17]
	if ss13 == "255" && ss14 == "255" && ss15 == "255" && ss16 == "255" && ss17 == "255" && ss18 == "255" {
		logrus.Info("访问设备异常,无法获取磁场数据")
	} else {
		hxH, _ := strconv.Atoi(ss13)
		hxL, _ := strconv.Atoi(ss14)
		hyH, _ := strconv.Atoi(ss15)
		hyL, _ := strconv.Atoi(ss16)
		hzH, _ := strconv.Atoi(ss17)
		hzL, _ := strconv.Atoi(ss18)
		k := 1.0
		hX := float64(hxH<<8 | hxL)
		hY := float64(hyH<<8 | hyL)
		hZ := float64(hzH<<8 | hzL)
		if hX >= k {
			hX -= 2 * k
		}
		if hY >= k {
			hY -= 2 * k
		}
		if hZ >= k {
			hZ -= 2 * k
		}
		logrus.Infof("Hx: %v", hX)
		logrus.Infof("Hy: %v", hY)
		logrus.Infof("Hz: %v", hZ)
	}
	// angular
	ss19 := splitS2[18]
	ss20 := splitS2[19]
	ss21 := splitS2[20]
	ss22 := splitS2[21]
	ss23 := splitS2[22]
	ss24 := splitS2[23]
	if ss19 == "255" && ss20 == "255" && ss21 == "255" && ss22 == "255" && ss23 == "255" && ss24 == "255" {
		logrus.Info("访问设备异常,无法获取角度数据")
	} else {
		rollH, _ := strconv.Atoi(ss19)
		rollL, _ := strconv.Atoi(ss20)
		pitchH, _ := strconv.Atoi(ss21)
		pitchL, _ := strconv.Atoi(ss22)
		yawH, _ := strconv.Atoi(ss23)
		YawL, _ := strconv.Atoi(ss24)
		k := 180.0
		roll := float64(rollH<<8|rollL) / 32768.0 * k
		pitch := float64(pitchH<<8|pitchL) / 32768.0 * k
		yaw := float64(yawH<<8|YawL) / 32768.0 * k
		if roll >= k {
			roll -= 2 * k
		}
		if pitch >= k {
			pitch -= 2 * k
		}
		if yaw >= k {
			yaw -= 2 * k
		}
		logrus.Infof("roll: %v °", roll)
		logrus.Infof("pitch: %v °", pitch)
		logrus.Infof("yaw: %v °", yaw)
	}
	//element
	ss25 := splitS2[len(splitS2)-8]
	ss26 := splitS2[len(splitS2)-7]
	ss27 := splitS2[len(splitS2)-6]
	ss28 := splitS2[len(splitS2)-5]
	ss29 := splitS2[len(splitS2)-4]
	ss30 := splitS2[len(splitS2)-3]
	ss31 := splitS2[len(splitS2)-2]
	ss32 := splitS2[len(splitS2)-1]
	if ss25 == "255" && ss26 == "255" && ss27 == "255" && ss28 == "255" && ss29 == "255" && ss30 == "255" {
		logrus.Info("访问设备异常,无法获取四元素数据")
	} else {
		q0H, _ := strconv.Atoi(ss25)
		q0L, _ := strconv.Atoi(ss26)
		q1H, _ := strconv.Atoi(ss27)
		q1L, _ := strconv.Atoi(ss28)
		q2H, _ := strconv.Atoi(ss29)
		q2L, _ := strconv.Atoi(ss30)
		q3H, _ := strconv.Atoi(ss31)
		q3L, _ := strconv.Atoi(ss32)
		k := 1.0
		q0 := float64(q0H<<8|q0L) / 32768.0
		q1 := float64(q1H<<8|q1L) / 32768.0
		q2 := float64(q2H<<8|q2L) / 32768.0
		q3 := float64(q3H<<8|q3L) / 32768.0
		if q0 >= k {
			q0 -= 2 * k
		}
		if q1 >= k {
			q1 -= 2 * k
		}
		if q2 >= k {
			q2 -= 2 * k
		}
		if q3 >= k {
			q3 -= 2 * k
		}
		logrus.Infof("Q0: %v", q0)
		logrus.Infof("Q1: %v", q1)
		logrus.Infof("Q2: %v", q2)
		logrus.Infof("Q3: %v", q3)
	}
}
