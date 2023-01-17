package buildpacks

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func BuildImage(Workdir string, Name string) error {

	//Check os type
	os_name := runtime.GOOS

	//Get dockerfile path
	dockerfile := Workdir + "/" + Name

	//Execute command
	if os_name == "windows" {
		command := "docker build -t " + Name + " " + dockerfile
		cmd := exec.Command("cmd", "/C", command)
		//Get html output
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}

	} else {
		command := "DOCKER_BUILDKIT=1 docker build -t " + Name + " " + dockerfile
		cmd := exec.Command("sh", "-c", command)
		//Get html output
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	}

	return nil
}
