/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"github.com/idoubi/onepub/platform"
	"github.com/idoubi/onepub/util"
	"github.com/spf13/cobra"
)

// pubCmd represents the pub command
var checkLoginCmd = &cobra.Command{
	Use:     "checklogin",
	Aliases: []string{"cl"},
	Short:   "check platform login status. default check all platform",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		// 获取推送平台，默认所有平台
		var publishPlatform []string
		p, err := cmd.Flags().GetString("platform")
		if err != nil {
			fmt.Println(err)
			return
		}

		if p != "" && !util.InSlice(p, platform.AllPlatform()) {
			fmt.Println("invalid platform " + p)
			return
		}

		if p != "" {
			publishPlatform = append(publishPlatform, p)
		} else {
			publishPlatform = platform.AllPlatform()
		}

		// 推送...
		for _, k := range publishPlatform {
			fmt.Println("检验 " + k + " 平台登录态...")

			err := platform.New(k).IsLogin()

			if err != nil {
				fmt.Println("认证失败，原因: " + err.Error())
			} else {
				fmt.Println("认证" + k + " 平台成功")
			}
		}
	},
}

func init() {
	checkLoginCmd.Flags().StringP("platform", "p", "", "default all platform")
	rootCmd.AddCommand(checkLoginCmd)
}
