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
	"device-cobra-cli/pkg/modbus/snow"
	"github.com/spf13/cobra"
)

var snowTTYSerial string

// snowCmd represents the snow command
var snowCmd = &cobra.Command{
	Use:   "snow",
	Short: "雨雪设备",
	Long: `雨雪设备检测,将要测试采集是否正常,雨雪设备的数据项:
    1. 雨雪`,
	Run: func(cmd *cobra.Command, args []string) {
		if snowTTYSerial==""{
			snowTTYSerial = snow.RTUSnowDevice
		}
		snow.GetSnowData(snowTTYSerial)
	},
}

func init() {
	rootCmd.AddCommand(snowCmd)
	snowCmd.Flags().StringVarP(&snowTTYSerial,"ttySerial","t","","snow device tty Serial")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// snowCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// snowCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
