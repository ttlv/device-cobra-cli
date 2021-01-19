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
	"device-cobra-cli/pkg/modbus/shutter"
	"device-cobra-cli/pkg/modbus/snow"
	"github.com/spf13/cobra"
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "测试所有设备",
	Long: `执行all命令可直接测试采集所有的设备的数据
    1. IMU从/dev/ttyUSB0
    2. 百叶窗1从/dev/ttyUSB1
    3. 百叶窗2从/dev/ttyUSB2
    4. 雨雪从/dev/ttyUSB3
`,
	Run: func(cmd *cobra.Command, args []string) {
		imu.GetIMUData(imu.RTUIMUDevice)
		shutter.GetShutter1Data(shutter.RTUShutter1Device)
		shutter.GetShutter2Data(shutter.RTUShutter2Device)
		snow.GetSnowData(snow.RTUSnowDevice)
	},
}

func init() {
	rootCmd.AddCommand(allCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// allCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// allCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
