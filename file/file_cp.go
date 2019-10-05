package main
import (
	"fmt"
	"os"
	"io"
)

func copyFile(src string, dst string) (err error){
	sf, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sf.Close()

	df, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer df.Close()
	_, err = io.Copy(df, sf)
	if err == nil {
		si, err := os.Stat(src)
		if err == nil {
			err = os.Chmod(dst, si.Mode())
		}
	}
	return
}

func main(){
	fmt.Println(copyFile("file.txt", "testcp"))
}
