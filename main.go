package main

import (
    "log"
    "os"

    "github.com/spf13/cobra"
    "gopkg.in/telebot.v3"
)

func main() {
    // Create a new Cobra root command
    var rootCmd = &cobra.Command{
        Use:   "telegram-bot",
        Short: "A Telegram bot developed in Golang",
        Run: func(cmd *cobra.Command, args []string) {
            // Retrieve bot token from environment variable
            token := os.Getenv("TELE_TOKEN")
            if token == "" {
                log.Fatal("TELE_TOKEN environment variable not set")
            }

            // Create a new Telegram bot
            bot, err := telebot.NewBot(telebot.Settings{
                Token: token,
            })
            if err != nil {
                log.Fatal(err)
            }

            // Add message handler for text messages
            bot.Handle(telebot.OnText, func(ctx telebot.Context) error {
                msg := ctx.Message()
                // Echo the received message
                _, err := bot.Send(msg.Chat, msg.Text)
                if err != nil {
                    log.Println(err)
                }
                return nil
            })

            // Start the bot
            bot.Start()
        },
    }

    // Execute the root command
    if err := rootCmd.Execute(); err != nil {
        log.Fatal(err)
    }
}

