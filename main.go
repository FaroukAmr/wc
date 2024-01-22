package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	cFlag := flag.Bool("c", false, "c")
	flag.Parse()
	var args []string =flag.Args()
	if(*cFlag){
		handleCountBytes(args)
	}
}

func handleCountBytes(args []string){
	if len(args)>0{
		fileName:= args[0]
		absPath, err := filepath.Abs(fileName)
        if err != nil {
            fmt.Println("Error getting absolute path:", err)
            return
        }
		file, err := os.Open(absPath)

		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		contents,err := io.ReadAll(file)
		if err!=nil{
			fmt.Println("Error reading file:", err)
			return
		}
		fmt.Println(len(contents))
		defer file.Close()
	}
}