package client

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var clientService *ClientService

func init() {
	clientService = NewClientService()
}

var rootCmd = &cobra.Command{
	Use:   "dcs",
	Short: "distributed cache system,",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Available Commands:")
		for _, c := range cmd.Commands() {
			fmt.Printf("  %-10s %s\n", c.Name(), c.Short)
		}
	},
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set data with key,value,expirationSeconds",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 3 {
			fmt.Println("Not enough args")
			return
		}
		_, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("invalid ttl")
			return
		}
		fmt.Println("set args", args)
		if err := clientService.Set(args); err != nil {
			log.Fatal(err.Error())
			return
		}
	},
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get data using key",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Not enough args")
			return
		}
		resp, err := clientService.Get(args)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(resp)
	},
}

var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete data using key",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Not enough args")
			return
		}
		err := clientService.Del(args)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
	},
}

func InitClient() error {
	rootCmd.AddCommand(setCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(delCmd)
	return rootCmd.Execute()
}
