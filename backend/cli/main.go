package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	dest   string
	source string
)

var rootCmd = &cobra.Command{
	Use:   "Copy your file from source to destination",
	Short: "Copy souce to destination",
	RunE: func(cmd *cobra.Command, args []string) error {
		if dest != "" && source != "" {
			err := saveFile(dest, source)
			if err != nil {
				fmt.Printf("An error occured", err)
			}
		} else {
			fmt.Println("Please Pass a Valid destination and source")
		}
		return nil
	},
}

func saveFile(dest, source string) error {
	dat, err := os.Open(source)
	if err != nil {
		return err
	}

	dst, err := os.Create(filepath.Join(dest, filepath.Base(source)))

	if err != nil {
		return err
	}

	defer dst.Close()

	_, err = io.Copy(dst, dat)
	if err != nil {
		return err
	}

	return nil

}

func init() {
	rootCmd.PersistentFlags().StringVarP(&dest, "dest", "d", "", "pass a destination where you want to copy")
	rootCmd.PersistentFlags().StringVarP(&source, "source", "s", "", "pass a destination from where you want to copy ")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("An error occured", err)
		os.Exit(1)
	}
}
