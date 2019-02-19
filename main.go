package main

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/afero"
	tfs "github.com/tronfs_2/filesystem"
)

func fullName(fileName string) string {
	fullName, _ := filepath.Abs(fileName)
	return fullName
}

func getAfero() *afero.Afero {
	fs := afero.NewMemMapFs()
	afs := &afero.Afero{Fs: fs}
	return afs
}

func main() {

	fmt.Println("Hello FileSystem")

	fileName := "AAA.txt"
	res := "AAA"
	data := []byte(res)
	afs := getAfero()
	afs.WriteFile(fullName(fileName), data, 0644)

	f := tfs.FileManager{ReadFile: afs.ReadFile}
	f.TurnOnCache()

	gData1, err1 := f.Get(fileName)
	if err1 != nil {
		fmt.Println("Find Error(gData1):", err1)
	}
	res1 := string(gData1[:])
	if res1 != res {
		fmt.Println("Get Data Error(gData1 != data):", res1, " != ", res)
	}

	afs.Remove(fullName(fileName))
	gData2, err2 := f.Get(fullName(".") + "/" + fileName)
	if err2 != nil {
		fmt.Println("Find Error(gData1):", err1)
	}
	res2 := string(gData2[:])
	if res2 != res {
		fmt.Println("Get Data Error(gData2 != data):", res2, " != ", res)
	}

}
