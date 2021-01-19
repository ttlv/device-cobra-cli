/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"device-cobra-cli/pkg/modbus/imu"
	"github.com/spf13/cobra"
)

var imuTTYSerial string

// imuCmd represents the imu command
var imuCmd = &cobra.Command{
	Use:   "imu",
	Short: "IMU设备",
	Long: `imu 设备检测,将要测试采集是否正常,IMU设备的数据项:
    1. 加速度
    2. 角速度
    3. 角度
    4. 磁场
    5. 四元素
`,
	Run: func(cmd *cobra.Command, args []string) {
		if imuTTYSerial==""{
			imuTTYSerial = imu.RTUIMUDevice
		}
		imu.GetIMUData(imuTTYSerial)
	},
}

func init() {
	rootCmd.AddCommand(imuCmd)
	imuCmd.Flags().StringVarP(&imuTTYSerial,"ttySerial","t","","imu device tty Serial")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// imuCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// imuCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
