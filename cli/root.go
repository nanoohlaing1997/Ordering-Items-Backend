package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	LocalRootFlag      bool
	PersistentRootFlag bool
	RootCmd            = &cobra.Command{
		Use:   "example",
		Short: "Researching and developing cli",
		Long:  `This is the simple example of a cobra program.`,
		// Run: func(cmd *cobra.Command, args []string) {
		// 	fmt.Println("Hello from the root command")
		// },
	}
	EchoCmd = &cobra.Command{
		Use:   "echo [string to echo]",
		Short: "Print given string to stdout",
		Args:  cobra.MinimumNArgs(1),
		Long:  `This is the simple example of a cobra cli with printing string.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo: " + strings.Join(args, ""))
		},
	}
	times   int
	TimeCmd = &cobra.Command{
		Use:   "time [sting to echo]",
		Short: "print give string to stdout multiple times",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if times == 0 {
				return errors.New("time cannot be 0")
			}
			for i := 0; i < times; i++ {
				fmt.Println("Echo times: " + strings.Join(args, " "))
			}
			return nil
		},
	}
)

func init() {
	RootCmd.PersistentFlags().BoolVarP(&PersistentRootFlag, "persistFlag", "p", false, "a persistant root flag")
	RootCmd.Flags().BoolVarP(&LocalRootFlag, "localFlag", "1", false, "a local root flag")
	TimeCmd.Flags().IntVarP(&times, "times", "t", 1, "number of times to echo to stdout")
	TimeCmd.MarkFlagRequired("times")
	RootCmd.AddCommand(EchoCmd)
	EchoCmd.AddCommand(TimeCmd)
}
