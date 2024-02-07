package cmdmanager

import "fmt"

type CMDManager struct {
}

func New() CMDManager {
	return CMDManager{}
}

func (fm CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter prices. ENTER to confirm")

	var prices []string
	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)

		if price == "0" {
			break
		}

		prices = append(prices, price)
	}

	return prices, nil
}

func (fm CMDManager) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}
