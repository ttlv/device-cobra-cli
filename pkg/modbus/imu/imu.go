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

func GetIMUData(rtuDevice string){
	handler := modbus.NewRTUClientHandler(rtuDevice)
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "N"
	handler.StopBits = 1
	handler.SlaveId = 80
	err := handler.Connect()
	if err != nil {
		logrus.Infof("访问串口%v异常",rtuDevice)
	}
	defer handler.Close()
	client := modbus.NewClient(handler)
	for i:=1;i<=5;i++{
		results, err := client.ReadHoldingRegisters(52, 3)
		if err != nil || results == nil {
			logrus.Info("访问IMU设备异常,无法获取加速度数据,尝试重新访问设备")
		}else {
			s1 := strings.Replace(fmt.Sprintf("%v", results), "[", "", -1)
			s2 := strings.Replace(s1, "]", "", -1)
			ss1 := strings.Split(s2, " ")[0]
			ss2 := strings.Split(s2, " ")[1]
			ss3 := strings.Split(s2, " ")[2]
			ss4 := strings.Split(s2, " ")[3]
			ss5 := strings.Split(s2, " ")[4]
			ss6 := strings.Split(s2, " ")[5]
			axh,_:=strconv.Atoi(ss1)
			axl,_:=strconv.Atoi(ss2)
			ayh,_:=strconv.Atoi(ss3)
			ayl,_:=strconv.Atoi(ss4)
			azh,_:=strconv.Atoi(ss5)
			azl,_:=strconv.Atoi(ss6)
			k:= 16.0
			accX := float64(axh << 8 | axl) / 32768.0 * k
			accY := float64(ayh << 8 | ayl) / 32768.0 * k
			accZ := float64(azh << 8 | azl) / 32768.0 * k
			if accX >= k{
				accX -= 2 * k
			}
			if accY>=k{
				accY -= 2 * k
			}
			if accZ>=k{
				accZ -= 2 * k
			}
			logrus.Infof("accX: %v m/s2【加速度X】",accX)
			logrus.Infof("accY: %v m/s2【加速度Y】",accY)
			logrus.Infof("accZ: %v m/s2【加速度Z】",accZ)
		}
		results, err = client.ReadHoldingRegisters(55, 3)
		if err != nil || results == nil {
			logrus.Info("访问IMU设备异常,无法获取角速度数据,尝试重新访问设备")
		}else{
			s1 := strings.Replace(fmt.Sprintf("%v", results), "[", "", -1)
			s2 := strings.Replace(s1, "]", "", -1)
			ss1 := strings.Split(s2, " ")[0]
			ss2 := strings.Split(s2, " ")[1]
			ss3 := strings.Split(s2, " ")[2]
			ss4 := strings.Split(s2, " ")[3]
			ss5 := strings.Split(s2, " ")[4]
			ss6 := strings.Split(s2, " ")[5]
			wxh,_:=strconv.Atoi(ss1)
			wxl,_:=strconv.Atoi(ss2)
			wyh,_:=strconv.Atoi(ss3)
			wyl,_:=strconv.Atoi(ss4)
			wzh,_:=strconv.Atoi(ss5)
			wzl,_:=strconv.Atoi(ss6)
			k:= 2000.0
			wX := float64(wxh << 8 | wxl) / 32768.0 * k
			wY := float64(wyh << 8 | wyl) / 32768.0 * k
			wZ := float64(wzh << 8 | wzl) / 32768.0 * k
			if wX >= k{
				wX -= 2 * k
			}
			if wY>=k{
				wY -= 2 * k
			}
			if wZ>=k{
				wZ -= 2 * k
			}
			logrus.Infof("WX: %v °/s【角速度X】",wX)
			logrus.Infof("WY: %v °/s【角速度X】",wY)
			logrus.Infof("WZ: %v °/s【角速度X】",wZ)
		}
		results, err = client.ReadHoldingRegisters(61, 3)
		if err != nil || results == nil {
			logrus.Info("访问IMU设备异常,无法获取角度数据,尝试重新访问设备")
		}else{
			s1 := strings.Replace(fmt.Sprintf("%v", results), "[", "", -1)
			s2 := strings.Replace(s1, "]", "", -1)
			ss1 := strings.Split(s2, " ")[0]
			ss2 := strings.Split(s2, " ")[1]
			ss3 := strings.Split(s2, " ")[2]
			ss4 := strings.Split(s2, " ")[3]
			ss5 := strings.Split(s2, " ")[4]
			ss6 := strings.Split(s2, " ")[5]
			rollH,_:=strconv.Atoi(ss1)
			rollL,_:=strconv.Atoi(ss2)
			pitchH,_:=strconv.Atoi(ss3)
			pitchL,_:=strconv.Atoi(ss4)
			yawH,_:=strconv.Atoi(ss5)
			YawL,_:=strconv.Atoi(ss6)
			k:= 180.0
			roll := float64(rollH << 8 | rollL) / 32768.0 * k
			pitch := float64(pitchH << 8 | pitchL) / 32768.0 * k
			yaw := float64(yawH << 8 | YawL) / 32768.0 * k
			if roll >= k{
				roll -= 2 * k
			}
			if pitch>=k{
				pitch -= 2 * k
			}
			if yaw>=k{
				yaw -= 2 * k
			}
			logrus.Infof("roll: %v °【角度X轴】",roll)
			logrus.Infof("pitch: %v °【角度Y轴】",pitch)
			logrus.Infof("yaw: %v °【角度Z轴】",yaw)
		}
		results, err = client.ReadHoldingRegisters(58, 4)
		if err != nil || results == nil {
			logrus.Info("访问IMU设备异常,无法获取磁场数据,尝试重新访问设备")
		}else{
			s1 := strings.Replace(fmt.Sprintf("%v", results), "[", "", -1)
			s2 := strings.Replace(s1, "]", "", -1)
			ss1 := strings.Split(s2, " ")[0]
			ss2 := strings.Split(s2, " ")[1]
			ss3 := strings.Split(s2, " ")[2]
			ss4 := strings.Split(s2, " ")[3]
			ss5 := strings.Split(s2, " ")[4]
			ss6 := strings.Split(s2, " ")[5]
			hxH,_:=strconv.Atoi(ss1)
			hxL,_:=strconv.Atoi(ss2)
			hyH,_:=strconv.Atoi(ss3)
			hyL,_:=strconv.Atoi(ss4)
			hzH,_:=strconv.Atoi(ss5)
			hzL,_:=strconv.Atoi(ss6)
			k:= 1.0
			hX := float64(hxH << 8 | hxL)
			hY := float64(hyH << 8 | hyL)
			hZ := float64(hzH << 8 | hzL)
			if hX >= k{
				hX -= 2 * k
			}
			if hY>=k{
				hY -= 2 * k
			}
			if hZ>=k{
				hZ -= 2 * k
			}
			logrus.Infof("hX: %v H【磁场X轴】",hX)
			logrus.Infof("hY: %v H【磁场Y轴】",hY)
			logrus.Infof("hZ: %v H【磁场Z轴】",hZ)
		}
		results, err = client.ReadHoldingRegisters(81, 4)
		if err != nil || results == nil {
			logrus.Info("访问IMU设备异常,无法获取四元素数据,尝试重新访问设备")
		}else{
			s1 := strings.Replace(fmt.Sprintf("%v", results), "[", "", -1)
			s2 := strings.Replace(s1, "]", "", -1)
			ss1 := strings.Split(s2, " ")[0]
			ss2 := strings.Split(s2, " ")[1]
			ss3 := strings.Split(s2, " ")[2]
			ss4 := strings.Split(s2, " ")[3]
			ss5 := strings.Split(s2, " ")[4]
			ss6 := strings.Split(s2, " ")[5]
			ss7 := strings.Split(s2, " ")[6]
			ss8 := strings.Split(s2, " ")[7]
			q0H,_:=strconv.Atoi(ss1)
			q0L,_:=strconv.Atoi(ss2)
			q1H,_:=strconv.Atoi(ss3)
			q1L,_:=strconv.Atoi(ss4)
			q2H,_:=strconv.Atoi(ss5)
			q2L,_:=strconv.Atoi(ss6)
			q3H,_:=strconv.Atoi(ss7)
			q3L,_:=strconv.Atoi(ss8)
			k:= 1.0
			q0 := float64(q0H << 8 | q0L)/32768.0
			q1 := float64(q1H << 8 | q1L)/32768.0
			q2 := float64(q2H << 8 | q2L)/32768.0
			q3 := float64(q3H << 8 | q3L)/32768.0
			if q0 >= k{
				q0 -= 2 * k
			}
			if q1>=k{
				q1 -= 2 * k
			}
			if q2>=k{
				q2 -= 2 * k
			}
			if q3>=k{
				q3 -= 2 * k
			}
			logrus.Infof("Q0: %v H【四元素Q0】",q0)
			logrus.Infof("Q1: %v H【四元素Q1】",q1)
			logrus.Infof("Q2: %v H【四元素Q2】",q2)
			logrus.Infof("Q3: %v H【四元素Q3】",q3)
		}
	}
}
