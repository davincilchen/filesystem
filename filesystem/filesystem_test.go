package filesystem

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func getAfero() *afero.Afero {
	fs := afero.NewMemMapFs()
	afs := &afero.Afero{Fs: fs}
	return afs
}

func TestFileManager_Get(t *testing.T) {
	fileName := "AAA.txt"
	data := []byte("AAA")
	afs := getAfero()
	afs.WriteFile(fullName(fileName), data, 0644)

	f := FileManager{ReadFile: afs.ReadFile}

	gData1, err1 := f.Get(fileName)
	gData2, err2 := f.Get(fullName(".") + "/" + fileName)

	assert.Nil(t, err1)
	assert.Nil(t, err2)
	assert.Equal(t, data, gData1)
	assert.Equal(t, data, gData2)
	assert.False(t, &gData1[0] == &gData2[0])

	failData, err := f.Get("a.txt")
	assert.NotNil(t, err)
	assert.Nil(t, failData)
}

func TestFileManager_TurnOnCache(t *testing.T) {
	fileName := "AAA.txt"
	data := []byte("AAA")
	afs := getAfero()
	afs.WriteFile(fullName(fileName), data, 0644)

	f := FileManager{ReadFile: afs.ReadFile}
	f.TurnOnCache()

	gData1, err1 := f.Get(fileName)
	assert.Nil(t, err1)
	assert.Equal(t, data, gData1)

	afs.Remove(fullName(fileName))
	gData2, err2 := f.Get(fullName(".") + "/" + fileName)
	assert.Nil(t, err2)
	assert.Equal(t, data, gData2)
}

func TestFileManager_TurnOffCache(t *testing.T) {
	fileName := "AAA.txt"
	data := []byte("AAA")
	afs := getAfero()
	afs.WriteFile(fullName(fileName), data, 0644)

	f := FileManager{ReadFile: afs.ReadFile}
	f.TurnOnCache()

	gData1, err1 := f.Get(fileName)
	assert.Nil(t, err1)
	assert.Equal(t, data, gData1)

	f.TurnOffCache()
	afs.Remove(fullName(fileName))
	gData2, err2 := f.Get(fileName)
	assert.NotNil(t, err2)
	assert.Nil(t, gData2)
}

func TestFileManager_ClearCache(t *testing.T) {
	fileName := "AAA.txt"
	data := []byte("AAA")
	afs := getAfero()
	afs.WriteFile(fullName(fileName), data, 0644)

	f := FileManager{ReadFile: afs.ReadFile}
	f.TurnOnCache()

	gData1, err1 := f.Get(fileName)
	assert.Nil(t, err1)
	assert.Equal(t, data, gData1)

	f.ClearCache()
	afs.Remove(fullName(fileName))
	gData2, err2 := f.Get(fileName)
	assert.NotNil(t, err2)
	assert.Nil(t, gData2)
}
