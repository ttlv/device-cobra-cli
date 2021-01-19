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

var shutter2TTYSerial string

// shutter2Cmd represents the shutter2 command
var shutter2Cmd = &cobra.Command{
	Use:   "shutter2",
	Short: "百叶窗2设备(pm2.5,pm10,噪音)",
	Long: `百叶窗设备检测,将要测试采集是否正常,百叶窗设备的数据项:
    1. PM2.5
    2. PM10
    3. 噪音`,
	Run: func(cmd *cobra.Command, args []string) {
		if shutter2TTYSerial==""{
			snowTTYSerial = shutter.RTUShutter2Device
		}
		shutter.GetShutter2Data(shutter2TTYSerial)
	},
}

func init() {
	rootCmd.AddCommand(shutter2Cmd)
	shutter2Cmd.Flags().StringVarP(&shutter2TTYSerial,"ttySerial","t","","shutter2 device tty Serial")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// shutter2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// shutter2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
