// package main

// import "demo/server"

// func main() {
// 	server.StartServer()
// }

package main

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
)

func main() {
	url := "https://api.clarvion.io/graphql"
	query := `{"query":"mutation { login(email: \"curiousbeing477@gmail.com\", password: \"Curiousbeing477@gmail.com\") { token user { userID email } } }"}`
	client := &http.Client{} // reuse TCP connection

	// client := &http.Client{}
	var wg sync.WaitGroup

	for i := 1; i <= 20; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			resp, err := client.Post(url, "application/json", bytes.NewBufferString(query))
			if err != nil {
				fmt.Printf("#%d error: %v\n", n, err)
				return
			}
			defer resp.Body.Close()
			// body, _ := io.ReadAll(resp.Body)
			fmt.Printf("#%d → %s\n", n, resp.Status)
		}(i)
	}

	wg.Wait()
}
