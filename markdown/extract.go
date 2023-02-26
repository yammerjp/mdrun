package markdown

import (
	"fmt"
	"os"
	"io"
)

func Extract() (error) {
	file, err := os.Open("README.md")
	if err != nil {
		return err
	}
	defer file.Close()
	buf := make([]byte, 1024)

	for {
		count, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Printf("%s", buf[:count])
	}
	return nil
}


