package main

import "github.com/JitenPalaparthi/depscan/cmd"

func main() {
	cmd.Execute()
	// gradle := new(implement.Gradle)
	// gradle.FilePath = "implement/test/build.gradle"
	// gdeps, err := gradle.Scan()
	// fmt.Println(gdeps, err)

	// maven := new(implement.Maven)
	// maven.FilePath = "implement/test/pom.xml"
	// gdeps, err := maven.Scan()
	// fmt.Println(gdeps, err)
}

//path := "/home/jiten/workspace/projects/depscan/test_repos/python/eLearning"
//path := /home/jiten/workspace/projects/depscan/test_repos/gradle/zuul
