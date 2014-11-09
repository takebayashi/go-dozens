package main

// simply show domains and records listed in dozens.jp

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
		return
	}
	domains, err := client.ListDomains()
	if err != nil {
		fmt.Println(err)
		return
	}
	for di, d := range domains {
		fmt.Printf("Domain #%02d: %s\n", di, d.Name)
		records, err := client.ListRecords(d)
		if err != nil {
			fmt.Println(err)
			return
		}
		for ri, r := range records {
			fmt.Printf("\tRecord #%02d: %s %s %s\n", ri, r.Type, r.Name, r.Content)
		}
	}
}
