/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
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

		database.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("todo"))
			if b == nil {
				return fmt.Errorf("get bucket: FAILED")
			}
			// we need cursor for iteration
			c := b.Cursor()
			for k, v := c.First(); k != nil; k, v = c.Next() {
				fmt.Println("Key: ", string(k), " Value: ", string(v))
			}
			// should return nil to complete the transaction
			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
