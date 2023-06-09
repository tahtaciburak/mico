package utils

import (
	"context"
	"github.com/PullRequestInc/go-gpt3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

func constructPrompt(input string) string {
	return `Act as a kubernetes command generator within given prompts, you only get a description from user and you should return the kubectl command in following format. No explanation. No other words. No details. 
{"command":"<RESULT>"} Don't use \t and \n in response. Give the result in one line. You are an intelligent assistant so you need to understand and predict user needs. Here is the description: ` + input
}

func GetKubectlCommand(prompt string) string {
	ctx := context.Background()
	apikey := viper.GetString("openai")
	client := gpt3.NewClient(apikey)
	systemMessage := gpt3.ChatCompletionRequestMessage{
		Role:    "system",
		Content: "You are a kubectl command helper you will help me to create kubectl commands according to my prompt.",
	}
	message := gpt3.ChatCompletionRequestMessage{
		Role:    "user",
		Content: constructPrompt(prompt)}

	var messages []gpt3.ChatCompletionRequestMessage
	messages = append(messages, systemMessage)
	messages = append(messages, message)

	resp, err := client.ChatCompletion(ctx, gpt3.ChatCompletionRequest{
		Model:    gpt3.GPT3Dot5Turbo0301,
		Messages: messages,
	})
	cobra.CheckErr(err)
	response := strings.TrimSpace(resp.Choices[0].Message.Content)
	command := strings.ReplaceAll(response, "{\"command\":\"", "")
	command = strings.ReplaceAll(command, "{\"command\": \"", "")
	return command[:len(command)-2]
}
