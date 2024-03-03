package app

import (
	"fmt"

	"golang.org/x/sys/windows/registry"
)

func GitGUI(action string) error {
	commands := []registrySpec{
		{
			Path:    "Directory\\Background\\shell\\git_gui",
			Command: "\"C:\\local.host\\modules\\git\\cmd\\git-gui.exe\" \"--working-dir\" \"%v.\"",
		},
		{
			Path:    "Directory\\shell\\git_gui",
			Command: "\"C:\\local.host\\modules\\git\\cmd\\git-gui.exe\" \"--working-dir\" \"%1\"",
		},
	}

	if action == "add" {
		fmt.Println("Adding Git GUI context menu")
		Elevate()

		for _, command := range commands {
			registry.DeleteKey(registry.CLASSES_ROOT, command.Path+"\\command")
			registry.DeleteKey(registry.CLASSES_ROOT, command.Path)

			key, _, err := registry.CreateKey(registry.CLASSES_ROOT, command.Path, registry.ALL_ACCESS)
			if err != nil {
				fmt.Println("Error creating key", err)
				return err
			}
			key.SetStringValue("", "Open Git GUI here")
			key.SetStringValue("Icon", "C:\\local.host\\modules\\git\\cmd\\git-gui.exe")
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
		fmt.Println("Removing Git GUI context menu")
		Elevate()

		for _, command := range commands {
			registry.DeleteKey(registry.CLASSES_ROOT, command.Path+"\\command")
			registry.DeleteKey(registry.CLASSES_ROOT, command.Path)
		}
	}

	return nil
}
