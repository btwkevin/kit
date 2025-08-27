package command
import (
	"os"
	"fmt"
	"path/filepath"
	"compress/zlib"
	"bytes"
)

func HandleCatFileCommand() {
	flag := os.Args[2]
	// add switch case for -p flag 
	commitSHA := os.Args[3]
	fmt.Println(flag,commitSHA)
	CatFileCmd(flag,commitSHA)
}

func CatFileCmd (flag string,commitSHA string ) {
	twoChar := commitSHA[:2]
	afterChar := commitSHA[2:]
	dir , _ := os.Getwd()
	path := filepath.Join(dir,".git/objects/",twoChar,afterChar)

	
	data , err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
			return
		}
		fmt.Println("Error reading file")
		return
	}

	reader , _ := zlib.NewReader(bytes.NewReader(data))
	decompressed := new(bytes.Buffer)

	_, err = decompressed.ReadFrom(reader)
	if err != nil {
		return
	}
	reader.Close()
	fmt.Println(decompressed.String())
	// fmt.Println(twoChar)
	// fmt.Println(afterChar)
}