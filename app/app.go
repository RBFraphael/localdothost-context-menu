package app

import (
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/urfave/cli"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
)

func Init() *cli.App {
	app := cli.NewApp()
	app.Name = "Local.Host Context Menu"
	app.Usage = "CLI tool to manage Local.Host context menu options"

	app.Commands = []cli.Command{
		{
			Name:  "git-bash",
			Usage: "Manage Git Bash context menu",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "path",
					Value: "C:\\local.host\\modules\\git\\git-bash.exe",
					Usage: "Path to git-bash.exe",
				},
				cli.StringFlag{
					Name:  "icon",
					Value: "C:\\local.host\\modules\\git\\git-bash.exe",
					Usage: "Path to Git Bash icon file",
				},
			},
			Subcommands: []cli.Command{
				{
					Name:   "add",
					Usage:  "Add Git Bash to context menu",
					Action: addGitBash,
				},
				{
					Name:   "remove",
					Usage:  "Remove Git Bash from context menu",
					Action: removeGitBash,
				},
			},
		},
		{
			Name:  "git-gui",
			Usage: "Manage Git GUI context menu",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "path",
					Value: "C:\\local.host\\modules\\git\\cmd\\git-gui.exe",
					Usage: "Path to git-gui.exe",
				},
				cli.StringFlag{
					Name:  "icon",
					Value: "C:\\local.host\\modules\\git\\cmd\\git-gui.exe",
					Usage: "Path to Git GUI icon file",
				},
			},
			Subcommands: []cli.Command{
				{
					Name:   "add",
					Usage:  "Add Git GUI to context menu",
					Action: addGitGui,
				},
				{
					Name:   "remove",
					Usage:  "Remove Git GUI from context menu",
					Action: removeGitGui,
				},
			},
		},
		{
			Name:  "git",
			Usage: "Manage Git GUI and Git Bash context menu",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "gui-path",
					Value: "C:\\local.host\\modules\\git\\cmd\\git-gui.exe",
					Usage: "Path to git-gui.exe",
				},
				cli.StringFlag{
					Name:  "bash-path",
					Value: "C:\\local.host\\modules\\git\\cmd\\git-gui.exe",
					Usage: "Path to git-gui.exe",
				},
			},
			Subcommands: []cli.Command{
				{
					Name:   "add",
					Usage:  "Add Git GUI to context menu",
					Action: addGitAll,
				},
				{
					Name:   "remove",
					Usage:  "Remove Git GUI from context menu",
					Action: removeGitAll,
				},
			},
		},
		{
			Name:  "symlink",
			Usage: "Manage Local.Host Symlink Action context menu",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "path",
					Value: "C:\\local.host\\modules\\tools\\symlink.exe",
					Usage: "Path to git-gui.exe",
				},
				cli.StringFlag{
					Name:  "icon",
					Value: "C:\\local.host\\modules\\gui\\Local.Host.exe",
					Usage: "Path to Local.Host icon file",
				},
			},
			Subcommands: []cli.Command{
				{
					Name:   "add",
					Usage:  "Add Local.Host Symlink Action to context menu",
					Action: addSymlink,
				},
				{
					Name:   "remove",
					Usage:  "Remove Local.Host Symlink Action from context menu",
					Action: removeSymlink,
				},
			},
		},
	}

	return app
}

func addGitBash(c *cli.Context) error {
	fmt.Println("Adding Git Bash to context menu")

	paths := []string{
		"Directory\\shell\\git_shell",
		"Directory\\Background\\shell\\git_shell",
	}

	elevate()
	for _, path := range paths {
		registry.DeleteKey(registry.CLASSES_ROOT, path)

		key, _, keyErr := registry.CreateKey(registry.CLASSES_ROOT, path, registry.ALL_ACCESS)
		if keyErr != nil {
			fmt.Println(keyErr)
			return keyErr
		}
		key.SetStringValue("", "Open Git Bash Here")
		key.SetStringValue("Icon", c.String("icon"))
		key.Close()

		commandKey, _, commandKeyErr := registry.CreateKey(registry.CLASSES_ROOT, path+"\\command", registry.ALL_ACCESS)
		if commandKeyErr != nil {
			fmt.Println(commandKeyErr)
			return commandKeyErr
		}
		commandKey.SetStringValue("", c.String("path")+" --cd=%V")
		commandKey.Close()
	}

	return nil
}

func removeGitBash(c *cli.Context) error {
	fmt.Println("Removing Git Bash from context menu")

	paths := []string{
		"Directory\\shell\\git_shell",
		"Directory\\Background\\shell\\git_shell",
	}

	elevate()
	for _, path := range paths {
		key := path + "\\command"
		err := registry.DeleteKey(registry.CLASSES_ROOT, key)
		if err != nil {
			fmt.Println("Error deleting "+key, err)
			return err
		}

		key = path
		err = registry.DeleteKey(registry.CLASSES_ROOT, key)
		if err != nil {
			fmt.Println("Error deleting "+key, err)
			return err
		}
	}

	return nil
}

func addGitGui(c *cli.Context) error {
	fmt.Println("Adding Git GUI to context menu")

	paths := []string{
		"Directory\\shell\\git_gui",
		"Directory\\Background\\shell\\git_gui",
	}

	elevate()
	for _, path := range paths {
		registry.DeleteKey(registry.CLASSES_ROOT, path)

		key, _, keyErr := registry.CreateKey(registry.CLASSES_ROOT, path, registry.ALL_ACCESS)
		if keyErr != nil {
			fmt.Println(keyErr)
			return keyErr
		}
		key.SetStringValue("", "Open Git GUI Here")
		key.SetStringValue("Icon", c.String("icon"))
		key.Close()

		commandKey, _, commandKeyErr := registry.CreateKey(registry.CLASSES_ROOT, path+"\\command", registry.ALL_ACCESS)
		if commandKeyErr != nil {
			fmt.Println(commandKeyErr)
			return commandKeyErr
		}
		commandKey.SetStringValue("", c.String("path")+" --cd=%V")
		commandKey.Close()
	}

	return nil
}

func removeGitGui(c *cli.Context) error {
	fmt.Println("Removing Git GUI from context menu")

	paths := []string{
		"Directory\\shell\\git_gui",
		"Directory\\Background\\shell\\git_gui",
	}

	elevate()
	for _, path := range paths {
		key := path + "\\command"
		err := registry.DeleteKey(registry.CLASSES_ROOT, key)
		if err != nil {
			fmt.Println("Error deleting "+key, err)
			return err
		}

		key = path
		err = registry.DeleteKey(registry.CLASSES_ROOT, key)
		if err != nil {
			fmt.Println("Error deleting "+key, err)
			return err
		}
	}

	return nil
}

func addGitAll(c *cli.Context) error {
	fmt.Println("Adding Git Bash and Git GUI to context menu")

	elevate()

	paths := []string{
		"Directory\\shell\\git_shell",
		"Directory\\Background\\shell\\git_shell",
	}

	for _, path := range paths {
		registry.DeleteKey(registry.CLASSES_ROOT, path)

		key, _, keyErr := registry.CreateKey(registry.CLASSES_ROOT, path, registry.ALL_ACCESS)
		if keyErr != nil {
			fmt.Println(keyErr)
			return keyErr
		}
		key.SetStringValue("", "Open Git Bash Here")
		key.SetStringValue("Icon", c.String("bash-path"))
		key.Close()

		commandKey, _, commandKeyErr := registry.CreateKey(registry.CLASSES_ROOT, path+"\\command", registry.ALL_ACCESS)
		if commandKeyErr != nil {
			fmt.Println(commandKeyErr)
			return commandKeyErr
		}
		commandKey.SetStringValue("", c.String("bash-path")+" --cd=%V")
		commandKey.Close()
	}

	paths = []string{
		"Directory\\shell\\git_gui",
		"Directory\\Background\\shell\\git_gui",
	}

	for _, path := range paths {
		registry.DeleteKey(registry.CLASSES_ROOT, path)

		key, _, keyErr := registry.CreateKey(registry.CLASSES_ROOT, path, registry.ALL_ACCESS)
		if keyErr != nil {
			fmt.Println(keyErr)
			return keyErr
		}
		key.SetStringValue("", "Open Git GUI Here")
		key.SetStringValue("Icon", c.String("gui-path"))
		key.Close()

		commandKey, _, commandKeyErr := registry.CreateKey(registry.CLASSES_ROOT, path+"\\command", registry.ALL_ACCESS)
		if commandKeyErr != nil {
			fmt.Println(commandKeyErr)
			return commandKeyErr
		}
		commandKey.SetStringValue("", c.String("gui-path")+" --cd=%V")
		commandKey.Close()
	}

	return nil
}

func removeGitAll(c *cli.Context) error {
	fmt.Println("Removing Git Bash and Git GUI from context menu")

	elevate()

	paths := []string{
		"Directory\\shell\\git_shell",
		"Directory\\Background\\shell\\git_shell",
		"Directory\\shell\\git_gui",
		"Directory\\Background\\shell\\git_gui",
	}

	for _, path := range paths {
		key := path + "\\command"
		err := registry.DeleteKey(registry.CLASSES_ROOT, key)
		if err != nil {
			fmt.Println("Error deleting "+key, err)
			return err
		}

		key = path
		err = registry.DeleteKey(registry.CLASSES_ROOT, key)
		if err != nil {
			fmt.Println("Error deleting "+key, err)
			return err
		}
	}

	return nil
}

func addSymlink(c *cli.Context) error {
	fmt.Println("Adding Local.Host Symlink Action to context menu")

	paths := []string{
		"Directory\\shell\\localhost_symlink_dir",
		"Directory\\Background\\shell\\localhost_symlink_dir",
	}

	elevate()
	for _, path := range paths {
		registry.DeleteKey(registry.CLASSES_ROOT, path)

		key, _, keyErr := registry.CreateKey(registry.CLASSES_ROOT, path, registry.ALL_ACCESS)
		if keyErr != nil {
			fmt.Println(keyErr)
			return keyErr
		}
		key.SetStringValue("", "Create Local.Host Symbolic Link")
		key.SetStringValue("Icon", c.String("icon"))
		key.Close()

		commandKey, _, commandKeyErr := registry.CreateKey(registry.CLASSES_ROOT, path+"\\command", registry.ALL_ACCESS)
		if commandKeyErr != nil {
			fmt.Println(commandKeyErr)
			return commandKeyErr
		}
		commandKey.SetStringValue("", c.String("path")+" d %1")
		commandKey.Close()
	}

	return nil
}

func removeSymlink(c *cli.Context) error {
	fmt.Println("Removing Local.Host Symlink Action from context menu")

	paths := []string{
		"Directory\\shell\\localhost_symlink_dir",
		"Directory\\Background\\shell\\localhost_symlink_dir",
	}

	elevate()
	for _, path := range paths {
		key := path + "\\command"
		err := registry.DeleteKey(registry.CLASSES_ROOT, key)
		if err != nil {
			fmt.Println("Error deleting "+key, err)
			return err
		}

		key = path
		err = registry.DeleteKey(registry.CLASSES_ROOT, key)
		if err != nil {
			fmt.Println("Error deleting "+key, err)
			return err
		}
	}

	return nil
}

func elevate() {
	if !checkAdmin() {
		runMeElevated()
		os.Exit(0)
	}
}

func runMeElevated() {
	verb := "runas"
	exe, _ := os.Executable()
	cwd, _ := os.Getwd()
	args := strings.Join(os.Args[1:], " ")

	verbPtr, _ := syscall.UTF16PtrFromString(verb)
	exePtr, _ := syscall.UTF16PtrFromString(exe)
	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
	argPtr, _ := syscall.UTF16PtrFromString(args)

	var showCmd int32 = 1 //SW_NORMAL

	err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
	if err != nil {
		fmt.Println(err)
	}
}

func checkAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	return err == nil
}
