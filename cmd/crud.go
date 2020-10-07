package cmd

import (
	"github.com/fj-onathan/crupi/models"
	"github.com/fj-onathan/crupi/src"
	"github.com/spf13/cobra"
)

var name string

func init() {
	crudCmd.Flags().StringVar(&name, "name", "", "Name of CRUD application.")
	rootCmd.AddCommand(crudCmd)
}

var crudCmd = &cobra.Command{
	Use:   "crud",
	Short: "Crud management options (create, list, etc)",
	Args:  cobra.ExactArgs(1),
	Run:   crudOptions,
}

func crudOptions(cmd *cobra.Command, args []string) {
	switch args[0] {
	case "create":
		createCrud(name)
	case "list":
		listCruds()
	}
}

func createCrud(name string) {
	// Crud name be empty
	if len(name) <= 0 {
		src.InlineResponse("error", "Need to fill name, use the flag:", "--name")
		return
	}

	// Verify if Crud is available
	if models.IfExistCrud(name) {
		src.InlineResponse("warning", "Crud already exists!")
		return
	}

	// Create new Crud
	crud := models.Crud{
		Name: name,
	}
	crud.AddCrud()
	src.InlineResponse("success", "Crud created successfully with name:", name)
}

func listCruds() {
	ls := models.ListCruds()
	src.InlineItems("Cruds", ls)
}
