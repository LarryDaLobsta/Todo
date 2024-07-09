/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (

	"fmt"
	"log"
	"time"
	"github.com/spf13/cobra"
)

// timeCmd represents the time command
var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "time test",
	Long: `Time test thing`,
	Run: getTimeInTimezone,
}




func getTimeInTimezone(cmd *cobra.Command, args []string){
 timezoneInput, _ := cmd.Flags().GetString("timezone")
 location, err := time.LoadLocation(timezoneInput)
 if err != nil {
  log.Fatalln("Error parsing timezone. Check to make sure the correct timezone was given")
 }
 currentTime := time.Now().In(location)
 fmt.Println(currentTime.Format(time.RFC1123))
}

func init() {
	rootCmd.AddCommand(timeCmd)

	// Here you will define your flags and configuration settings.

	timeCmd.Flags().StringP("timezone", "t", "EST", "Desired timezone")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// timeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// timeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
