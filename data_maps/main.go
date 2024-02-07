package main

import "fmt"

func main() {
	website := map[string]string{
		"Google":   "https://google.com",
		"Facebook": "https://facebook.com",
		"Youtube":  "https://youtube.com",
	}
	fmt.Println(website)
	fmt.Println(website["Google"])

	website["Instagram"] = "https://instagram.com"
	fmt.Println(website)

	delete(website, "Facebook")
	fmt.Println(website)
}
