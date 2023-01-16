package main

import (
	bp "dauqu.com/buildpacks/buildpacks"
	"fmt"
)

func main() {

	project_dir := "/var/www"
	name := "laravel"
	// expose_port := "3000"

	//Dectect language
	language, err := bp.DetectLanguage(project_dir + "/" + name)
	if err != nil {
		fmt.Println(err)
	}

	//Create dockerfile
	// err = bp.CreateDockerfile(project_dir, expose_port, language, name)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	fmt.Printf("\x1b[34m%-10s %-10s\x1b[0m\n", language, " Project detected")

	// //Build image
	// err = bp.BuildImage(project_dir, name)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
