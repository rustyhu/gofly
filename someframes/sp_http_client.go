package someframes

import (
	"fmt"
	"net/http"
	"os"
)

// A client to download http response into a local file.

const DEFAULT_BUFSIZE = 8192

// RequestAndWrite requests then write response into a file.
//   - Fill filename and URL;
//   - BufSize is optional if do not want just use 0;
func RequestAndWriteFile(filename string, URL string, BufSize int) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	response, err := http.Get(URL)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	// use default
	if BufSize <= 0 {
		BufSize = DEFAULT_BUFSIZE
	}
	resBytes := make([]byte, BufSize)
	rCnt := 0
	for n, err := response.Body.Read(resBytes); n != 0; n, err = response.Body.Read(resBytes) {
		rCnt++
		if err != nil {
			fmt.Println("Read http err:", err)
			continue
		}

		_, err = file.Write(resBytes[:n])
		if err != nil {
			fmt.Println("Write file err:", err)
		}
	}

	fmt.Printf("Download finish! Use %v bytes buffer, %v times.\n", BufSize, rCnt)
}
