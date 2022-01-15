/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"strings"
	"time"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		database, err := bolt.Open("my.db", 0600, &bolt.Options{})
		if err != nil {
			log.Fatal(err)
		}
		defer database.Close()

		database.Update(func(t *bolt.Tx) error {
			bucket, err := t.CreateBucketIfNotExists([]byte("todo"))
			if err != nil {
				return err
			}
			bucket.Put([]byte(time.Now().Format(time.RFC822)), []byte(strings.Join(args, " / ")))
			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
