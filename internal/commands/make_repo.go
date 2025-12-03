package commands

import (
	"fmt"
	"path/filepath"

	"github.com/fsjorgeluis/sazerac/internal"
	"github.com/fsjorgeluis/sazerac/internal/templates"
	"github.com/spf13/cobra"
)

func NewMakeRepoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "make repo <Entity>",
		Short: "Generate a dummy repository and its MySQL implementation",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			entity := args[0]

			// Repository interface
			outInterface := filepath.Join(
				"internal/repository",
				internal.ToSnake(entity)+"_repository.go",
			)

			// Infrastructure implementation
			outInfra := filepath.Join(
				"infrastructure/database/mysql",
				internal.ToSnake(entity)+"_mysql.go",
			)

			data := map[string]any{
				"Entity": entity,
				"Module": internal.GetModuleName(),
			}

			err := internal.WriteTemplate(templates.FS, "repository/repo_interface.go.tpl", outInterface, data)
			if err != nil {
				return err
			}

			err = internal.WriteTemplate(templates.FS, "repository/repo_mysql.go.tpl", outInfra, data)
			if err != nil {
				return err
			}

			fmt.Println("Repository served 🥃:", outInterface)
			fmt.Println("MySQL dummy implementation served 🥃:", outInfra)
			return nil
		},
	}

	return cmd
}
