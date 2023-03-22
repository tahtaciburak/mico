package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	utils "github.com/tahtaciburak/mico/pkg"
	"os"
)

var prompt = ""

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&prompt, "prompt", "p", "", "Prompt to generate")
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("$HOME/.mico/")
}

var rootCmd = &cobra.Command{
	Use:   "mico",
	Short: "mico [micho] is a AI driven kubectl helper",
	Long:  `Write your prompt and get your kubectl command`,
	Args: func(cmd *cobra.Command, args []string) error {
		p, err := cmd.PersistentFlags().GetString("prompt")
		cobra.CheckErr(err)

		if len(p) < 1 {
			cmd.Help()
			os.Exit(0)
		}
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("You should run `mico configure` before running.")
			os.Exit(1)
		}
		if utils.IsValidPrompt(prompt) {
			return nil
		}
		return fmt.Errorf("invalid prompt : %s", args[0])
	},

	Run: func(cmd *cobra.Command, args []string) {
		comm := utils.GetKubectlCommand(prompt)
		fmt.Println(comm)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
