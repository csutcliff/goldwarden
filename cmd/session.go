package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/quexten/goldwarden/ipc/messages"
	"github.com/spf13/cobra"
)

// sessionCmd represents the run command
var sessionCmd = &cobra.Command{
	Use:    "session",
	Hidden: true,
	Short:  "Starts a new session",
	Long:   `Starts a new session.`,
	Run: func(cmd *cobra.Command, args []string) {
		for {
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			text = strings.TrimSuffix(text, "\n")
			args := strings.Split(text, " ")
			rootCmd.SetArgs(args)
			rootCmd.Execute()
		}
	},
}

var pinentry = &cobra.Command{
	Use:    "pinentry",
	Hidden: true,
	Short:  "Registers as a pinentry program",
	Long:   `Registers as a pinentry program.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var authenticateSession = &cobra.Command{
	Use:    "authenticate-session",
	Hidden: true,
	Short:  "Authenticates a session",
	Long:   `Authenticates a session.`,
	Run: func(cmd *cobra.Command, args []string) {
		token := args[0]
		response, err := commandClient.SendToAgent(messages.SessionAuthRequest{Token: token})
		if err != nil {
			panic(err)
		}
		fmt.Println(response.(messages.SessionAuthResponse).Verified)
	},
}

func init() {
	rootCmd.AddCommand(sessionCmd)
	rootCmd.AddCommand(pinentry)
	rootCmd.AddCommand(authenticateSession)
}
