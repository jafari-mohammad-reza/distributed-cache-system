package client

import (
	"fmt"
	"log"

	"github.com/jafari-mohammad-reza/distributed-cache-system/pb"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var commandClient pb.CommandClient

func init() {
	conn, err := grpc.NewClient(fmt.Sprintf("localhost:%d", 6090), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to broker: %v", err)
	}
	commandClient = pb.NewCommandClient(conn)
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

func InitClient() error {

	return rootCmd.Execute()
}
