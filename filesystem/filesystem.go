package filesystem

//TODO: lock for rontine if necessarry
import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	cache "github.com/patrickmn/go-cache"
)

const (
	//DefaultCacheExpiration is a
	DefaultCacheExpiration = cache.NoExpiration
)

//FileSystem is a
type FileSystem struct {
	FileManager
}

//FileManager is a
type FileManager struct {
	cache    *cache.Cache
	ReadFile func(string) ([]byte, error)
}

//Initialize is a
func (t *FileSystem) Initialize() error {
	//TODO:
	return nil
}

//Reinitialize is a
func (t *FileSystem) Reinitialize() error {
	//TODO:
	return nil
}

//Uninitialize is a
func (t *FileSystem) Uninitialize() error {
	//TODO:
	return nil
}

//fullName is a
func fullName(fileName string) string {
	fullName, _ := filepath.Abs(fileName)
	return fullName
}

//readFile is a
func (t *FileManager) readFile(fileName string) ([]byte, error) {
	if t.ReadFile == nil {
		return ioutil.ReadFile(fileName)
	}

	return t.ReadFile(fileName)
}

//updateCache is a
func (t *FileManager) updateCache(fileName string, data []byte) {

	if t.cache == nil {
		return
	}

	t.cache.Set(fileName, data, cache.DefaultExpiration)
}

//readCache is a
func (t *FileManager) readCache(fileName string) ([]byte, error) {

	if t.cache == nil {
		return nil, nil
	}

	raw, found := t.cache.Get(fileName)
	if !found {
		return nil, nil
	}

	data, ok := raw.([]byte)
	if ok {
		return data, nil
	}
	return nil, fmt.Errorf("data type is not []byte")

}

//Get is a
func (t *FileManager) Get(fileName string) ([]byte, error) {

	pathName := fileName
	if filepath.IsAbs(fileName) == false {
		pathName = fullName(fileName)
	}
	pathName = strings.Replace(pathName, "/", "\\", -1)

	data, err := t.readCache(pathName)
	if err != nil {
		return nil, err
	}
	if data != nil {
		return data, nil
	}

	data, err = t.readFile(pathName)
	if err == nil && data != nil {
		t.updateCache(pathName, data)
	}

	return data, err

}

//TurnOnCache is a
func (t *FileManager) TurnOnCache() {
	if t.cache != nil {
		return
	}
	t.cache = cache.New(DefaultCacheExpiration, 0)
}

//TurnOffCache is a
func (t *FileManager) TurnOffCache() {
	t.cache = nil
}

//ClearCache is a
func (t *FileManager) ClearCache() {
	if t.cache == nil {
		return
	}

	t.cache.Flush()
}
