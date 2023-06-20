package main

import (
	"github.com/JitenPalaparthi/depscan/cmd"
)

func main() {
	cmd.Execute()
	// fmt.Println("Hello Muruaga!")
	// fcount, dcount := 0, 0
	// filepath.WalkDir("/home/jiten/workspace/personal/aka-labs.io", func(path string, d fs.DirEntry, err error) error {
	// 	if d.IsDir() {
	// 		dcount++
	// 	} else {
	// 		fcount++
	// 	}
	// 	fmt.Println(path, d.Info, err)

	// 	return err
	// })

	// fmt.Println("Total files scanned", fcount)
	// fmt.Println("Total directories scaned", dcount)
}
