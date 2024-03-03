package app

import (
	"github.com/urfave/cli"
)

type registrySpec struct {
	Path    string
	Command string
}

func Init() *cli.App {
	app := cli.NewApp()

	app.Name = "Local.Host Context Menu Manager"
	app.Usage = "Manage the context menu for Local.Host"
	app.Version = "1.0.0"

	app.Commands = []cli.Command{
		{
			Name:  "symlink",
			Usage: "Manage the symlink context menu",
			Subcommands: []cli.Command{
				{
					Name:   "add",
					Usage:  "Add the Symmlink context menu",
					Action: func(c *cli.Context) error { return Symlink("add") },
				},
				{
					Name:   "remove",
					Usage:  "Remove the Symmlink context menu",
					Action: func(c *cli.Context) error { return Symlink("remove") },
				},
			},
		},
		{
			Name:  "git-gui",
			Usage: "Manage the Git GUI context menu",
			Subcommands: []cli.Command{
				{
					Name:   "add",
					Usage:  "Add the Git GUI context menu",
					Action: func(c *cli.Context) error { return GitGUI("add") },
				},
				{
					Name:   "remove",
					Usage:  "Remove the Git GUI context menu",
					Action: func(c *cli.Context) error { return GitGUI("remove") },
				},
			},
		},
		{
			Name:  "git-bash",
			Usage: "Manage the Git Bash context menu",
			Subcommands: []cli.Command{
				{
					Name:   "add",
					Usage:  "Add the Git Bash context menu",
					Action: func(c *cli.Context) error { return GitShell("add") },
				},
				{
					Name:   "remove",
					Usage:  "Remove the Git Bash context menu",
					Action: func(c *cli.Context) error { return GitShell("remove") },
				},
			},
		},
		{
			Name:  "git",
			Usage: "Manage both Git Bash and Git GUI context menus",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "Add Git Bash and Git GUI to context menu",
					Action: func(c *cli.Context) error {
						err := GitShell("add")
						if err != nil {
							return err
						}
						return GitGUI("add")
					},
				},
				{
					Name:  "remove",
					Usage: "Remove Git Bash and Git GUI to context menu",
					Action: func(c *cli.Context) error {
						err := GitShell("remove")
						if err != nil {
							return err
						}
						return GitGUI("remove")
					},
				},
			},
		},
	}

	return app
}
