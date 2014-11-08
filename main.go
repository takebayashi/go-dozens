package main

// simply show domains listed in dozens.jp

import (
	"fmt"
	"github.com/takebayashi/go-dozens/dozens"
	"net/http"
	"os"
)

func main() {
	user := os.Args[1]
	key := os.Args[2]
	client, err := dozens.NewClient(&http.Client{}, user, key)
	if err != nil {
		fmt.Println(err)
	}
	domains, err := client.ListDomains()
	if err != nil {
		fmt.Println(err)
	}
	for _, d := range domains {
		fmt.Println(d.Name)
	}
}
