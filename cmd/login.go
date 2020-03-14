package cmd

import (
	"fmt"

	"github.com/idoubi/onepub/platform"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use: "login",
	Run: func(cmd *cobra.Command, args []string) {
		plat := platform.New("juejin")
		err := plat.Login()
		fmt.Println(err)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
