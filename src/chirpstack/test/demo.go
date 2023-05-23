package test

import (
	_ "fmt"
	"log"
	"os"
	"text/template"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var cfgFiles *[]string // config file
var version string

func Ta() {
}

var configCmd = &cobra.Command{
	Use:   "configfile",
	Short: "Print the ChirpStack Gateway Bridge configuration file",
	RunE: func(cmd *cobra.Command, args []string) error {
		t := template.Must(template.New("config").Parse(configTemplate))
		err := t.Execute(os.Stdout, config.C)
		if err != nil {
			return errors.Wrap(err, "execute config template error")
		}
		return nil
	},
}

var rootCmd = &cobra.Command{
	Use:   "cobra-use",
	Short: "abstracts the packet_forwarder protocol into Protobuf or JSON over MQTT",
	Long: `cobra-long`,
	RunE: run,
}

// Execute executes the root command.
func Execute(v string) {
	version = v
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	// fmt.Println("init")
	rootCmd.AddCommand(configCmd)
}
