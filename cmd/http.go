package cmd

import (
	"github.com/fj-onathan/crupi/http"
	"github.com/fj-onathan/crupi/models"
	"github.com/fj-onathan/crupi/src"
	"github.com/spf13/cobra"
)

var crud string

func init() {
	cr := serveCmd.PersistentFlags()

	cr.StringVarP(&crud, "crud", "c", "", "Specific CRUD name (already registered)")
	err := cobra.MarkFlagRequired(cr, "crud")
	if err != nil {
		panic(err)
	}
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Listen and serve Crud",
	Run:   ServeCrud,
}

// ServeCrud starts http server with selected crud
func ServeCrud(cmd *cobra.Command, args []string) {
	// FakeRoutes properly to testing some routing handling
	FakeRoutes := []string{"/fake", "/route", "/path"}

	// Verify if Crud is available
	if !models.IfExistCrud(crud) {
		src.InlineResponse("error", "Crud requested doesn't exists!")
		return
	}

	http.StartServer(FakeRoutes)
}
