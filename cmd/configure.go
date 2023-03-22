package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

func init() {
	rootCmd.AddCommand(configureCmd)
}

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "add your keys",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var openAIKey string
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		configPath := home + "/.mico"
		if _, err = os.Stat(configPath); os.IsNotExist(err) {
			err = os.MkdirAll(configPath, os.ModePerm)
			cobra.CheckErr(err)
		} else {
			err = os.RemoveAll(configPath)
			cobra.CheckErr(err)
			err = os.MkdirAll(configPath, os.ModePerm)
			cobra.CheckErr(err)
		}

		fmt.Print("OPENAI_API_KEY: ")
		fmt.Scan(&openAIKey)
		viper.Set("openai", openAIKey)
		err = viper.SafeWriteConfig()
		cobra.CheckErr(err)
	},
}
