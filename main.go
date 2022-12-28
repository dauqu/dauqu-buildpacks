package main

import (
	bp "dauqu.com/buildpacks/buildpacks"
	"fmt"
)

func main() {

	//Get current directory
	// dir, err := os.Getwd()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	project_dir := "/Users/harshaweb/Documents/projects"
	name := "wordpress"
	expose_port := "80"

	//Dectect language
	language, err := bp.DetectLanguage(project_dir + "/" + name)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(language + " project detected")

	//Create dockerfile
	err = bp.CreateDockerfile(project_dir, expose_port, language, name)
	if err != nil {
		fmt.Println(err)
	}

	// //Build image
	err = bp.BuildImage(project_dir, name)
	if err != nil {
		fmt.Println(err)
	}
}
