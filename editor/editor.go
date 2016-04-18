package editor

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func Run(editor, args, file string) error {
	cmd := exec.Command("bash", "-c", fmt.Sprintf("%s %s %s", editor, args, file))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return errors.New("Problem opening " + editor)
	}
	return nil
}
