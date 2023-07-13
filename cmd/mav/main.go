package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/bluenviron/gomavlib/v2"
	"github.com/bluenviron/gomavlib/v2/pkg/dialects/ardupilotmega"
	"github.com/fairytale5571/mav-cli/pkg/mavlink"
	"github.com/fairytale5571/mav-cli/pkg/simulator"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:              "cli",
	Short:            "CLI application for MAVLink messages",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {},
	Run:              func(cmd *cobra.Command, args []string) {},
}

func listCmd(simulators simulator.SimulatorsInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "px4-list",
		Short: "List all connected simulators",
		Run: func(cmd *cobra.Command, args []string) {
			sims := simulators.GetAll()
			if len(sims) == 0 {
				fmt.Println("simulators not found")
				return
			}
			for id := range sims {
				fmt.Println("Simulator ID:", id)
			}
		},
	}
}

func statusCmd(simulators simulator.SimulatorsInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "status [id] [number]",
		Short: "Print position of the drone with id drone and message number",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			id, _ := strconv.Atoi(args[0])
			n, _ := strconv.Atoi(args[1])
			if n < 0 {
				log.Println(",message number must be positive")
				return
			}
			if n > 100 {
				log.Println("message number must be less than 100")
				return
			}

			sim, exists := simulators.Get(id)
			if !exists {
				log.Println("Simulator not found")
				return
			}

			sim.Messages.Do(func(value interface{}) {
				if n == 0 {
					if msg, ok := value.(*ardupilotmega.MessageGpsRawInt); ok {
						log.Printf("SIM ID: %d - Lattitude: %d Longtitude: %d Altitude: %d\n", id, msg.Lat, msg.Lon, msg.Alt)
					} else {
						log.Println("No message available")
					}
				}
				n--
			})
		},
	}
}

func ExecuteCommands(simulators simulator.SimulatorsInterface) {
	rootCmd.AddCommand(listCmd(simulators))
	rootCmd.AddCommand(statusCmd(simulators))
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

func getSystemIDbyAddress(address string) int {
	addressBytes := []byte(address)
	hash := sha256.Sum256(addressBytes)
	systemID := binary.LittleEndian.Uint32(hash[:4])
	systemID %= 255
	systemID++
	return int(systemID)
}

func mustReadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func main() {
	mustReadConfig()

	simulators := simulator.NewSimulators()
	var node *gomavlib.Node
	var handler *mavlink.Handler

	for _, address := range viper.GetStringSlice("nodes") {
		node = mavlink.InitNode(address, getSystemIDbyAddress(address))
		handler = mavlink.NewHandler(node, simulators)
		go handler.HandleMessages(getSystemIDbyAddress(address))
	}

	go ExecuteCommands(simulators)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("\nEnter command: ")
		cmdString, _ := reader.ReadString('\n')
		args := strings.Fields(cmdString)
		rootCmd.SetArgs(args)
		err := rootCmd.Execute()
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
	}
}
