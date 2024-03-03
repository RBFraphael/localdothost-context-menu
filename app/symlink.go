package app

import (
	"fmt"

	"golang.org/x/sys/windows/registry"
)

func Symlink(action string) error {
	commands := []registrySpec{
		{
			Path:    "Directory\\Background\\shell\\localhost_symlink_dir",
			Command: "\"C:\\local.host\\modules\\tools\\symlink.exe\" \"d\" \"%v.\"",
		},
		{
			Path:    "Directory\\shell\\localhost_symlink_dir",
			Command: "\"C:\\local.host\\modules\\tools\\symlink.exe\" \"d\" \"%1\"",
		},
	}

	if action == "add" {
		fmt.Println("Adding symlink context menu")
		Elevate()

		for _, command := range commands {
			registry.DeleteKey(registry.CLASSES_ROOT, command.Path+"\\command")
			registry.DeleteKey(registry.CLASSES_ROOT, command.Path)

			key, _, err := registry.CreateKey(registry.CLASSES_ROOT, command.Path, registry.ALL_ACCESS)
			if err != nil {
				fmt.Println("Error creating key", err)
				return err
			}
			key.SetStringValue("", "Create Local.Host Link")
			key.SetStringValue("Icon", "C:\\local.host\\modules\\gui\\Local.Host.exe")
			key.Close()

			key, _, err = registry.CreateKey(registry.CLASSES_ROOT, command.Path+"\\command", registry.ALL_ACCESS)
			if err != nil {
				fmt.Println("Error creating command key", err)
				return err
			}
			key.SetStringValue("", command.Command)
			key.Close()
		}
	} else if action == "remove" {
		fmt.Println("Removing symlink context menu")
		Elevate()

		for _, command := range commands {
			registry.DeleteKey(registry.CLASSES_ROOT, command.Path+"\\command")
			registry.DeleteKey(registry.CLASSES_ROOT, command.Path)
		}
	}

	return nil
}
