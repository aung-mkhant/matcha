package main
import (
	"os"
	"path/filepath"
)

// Logger
func logToFile(path string,data string) {
dir, err := os.Getwd()
		check(err)
	p := filepath.Join(dir, path)
	f, err := os.Create(p)
		check(err)
 _, err = f.WriteString(data)
	check(err)
	f.Sync()
}
	
func check(e error) {
    if e != nil {
        panic(e)
    }
}