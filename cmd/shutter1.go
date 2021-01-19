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
	"device-cobra-cli/pkg/modbus/shutter"
	"github.com/spf13/cobra"
)

var shutter1TTYSerial string

// shutter1Cmd represents the shutter1 command
var shutter1Cmd = &cobra.Command{
	Use:   "shutter1",
	Short: "百叶窗1设备(光强,二氧化碳,大气压强,温湿度)",
	Long: `百叶窗设备检测,将要测试采集是否正常,百叶窗设备的数据项:
    1. 光强
    2. 二氧化碳
    3. 大气压强
    4. 温度
    5 湿度
`,
	Run: func(cmd *cobra.Command, args []string) {
		if shutter1TTYSerial==""{
			snowTTYSerial = shutter.RTUShutter1Device
		}
		shutter.GetShutter1Data(shutter1TTYSerial)
	},
}

func init() {
	rootCmd.AddCommand(shutter1Cmd)
	shutter1Cmd.Flags().StringVarP(&shutter1TTYSerial,"ttySerial","t","","shutter1 device tty Serial")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// shutter1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// shutter1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
