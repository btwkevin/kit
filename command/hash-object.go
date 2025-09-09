package command

import (
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func HandleHashObjectCommand() {
	args := os.Args
	var flag, filepath string
	if len(args) > 2 {
		flag = args[2]
	}

	if len(args) > 3 {
		filepath = args[3]
	}

	if filepath == "" {
		filepath = flag
		flag = ""
	}
	fmt.Println(filepath, flag)
	HashObject(flag, filepath)
}

func HashObject(flag string, filepath string) {
	fileHandle, err := os.Open(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
			return
		}
	}
	defer fileHandle.Close()

	fileContent, _ := io.ReadAll(fileHandle)
	fileStat, _ := fileHandle.Stat()
	fileLength := fileStat.Size()

	blobHeader := fmt.Sprintf("blob %d\x00", fileLength)
	blob := append([]byte(blobHeader), fileContent...)
	fmt.Printf("Blob (%d bytes):\n%s\n", len(blob), blob)

	hash := sha1.Sum(blob)
	fmt.Printf("SHA-1: %x\n", hash)
	hashStr := fmt.Sprintf("%x", hash)

	// TODO  := mkdir dir and write file with zlib compression
	if flag != "" {
		dir := hashStr[:2]
		file := hashStr[2:]
		fmt.Println(dir, file)
	}
}
