package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"hamilton/service"
	"io/ioutil"
	"os"
)

var ticketCmd = &cobra.Command{
	Use:   "ticket",
	Short: "Generate a ticket SVG",
	Long: `Generate a ticket SVG with the given input file as describe by:

	{
		"uuid": "12345678-1234-1234-1234-123456789012",
		"name": "Example User",
		"email": "example@uq.edu.au",
		"concert: {
			"uuid": "12345678-1234-1234-1234-123456789012",
			"name": "Example Concert",
			"date": "2021-01-01",
			"venue": "Example Venue",
		}
	}
`,
	Run: func(cmd *cobra.Command, args []string) {
		input, _ := cmd.Flags().GetString("input")
		output, _ := cmd.Flags().GetString("output")

		rawInfo, _ := ioutil.ReadFile(input)
		var info service.Ticket
		err := json.Unmarshal(rawInfo, &info)
		if err != nil {
			errorAndClose(err, output)
		}

		pencil := service.NewDrawer()
		concert, err := pencil.DrawTicket(info)

		f, err := os.Create(fmt.Sprintf("%s.svg", output))
		if err != nil {
			errorAndClose(err, output)
		}
		defer f.Close()

		_, err = f.WriteString(concert)
		if err != nil {
			errorAndClose(err, output)
			return
		}
	},
}

func errorAndClose(originalError error, path string) {
	fmt.Println(originalError)
	f, err := os.Create(fmt.Sprintf("%s.json", path))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("{\"message\": \"%s\"}", originalError.Error()))
	if err != nil {
		return
	}
	os.Exit(1)
}

func init() {
	generateCmd.AddCommand(ticketCmd)

	ticketCmd.Flags().StringP("input", "i", "input.json", "Path of the input file")
	ticketCmd.Flags().StringP("output", "o", "output", "Path of the output file without extension")
}
