package cmd

import (
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version number of this cli tool",
	Long:  "The version number is stored in a config file, loaded by viper. This command will print that version number out",
	Run:   ShowVersion,
}

// create a handler struct
type HttpHandler struct{}

// implement `ServeHTTP` method on `HttpHandler` struct
func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// create response binary data
	data := []byte("Hello World!") // slice of bytes
	// write `data` to response
	res.Write(data)
}

// ShowVersion prints out the cli version number
func ShowVersion(cmd *cobra.Command, args []string) {
	showVersion()
	// create a new handler
	handler := HttpHandler{}
	// listen and serve
	http.ListenAndServe(":9000", handler)
}

// Dumb version of ShowVersion(). Used for testing
func showVersion() {
	// internal theme/style:
	color.Red.Printf("go-cli-boilerplate version: %s \n", viper.GetString("version"))
}
