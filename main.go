package main

import (
	"fmt"
	"github.com/tebeka/selenium"
)

func main() {
	caps := selenium.Capabilities{"browserName": "chrome", "version": "74.0"}
	driver, err := selenium.NewRemote(caps, "http://localhost:4444/wd/hub")
	if err != nil {
		fmt.Printf("\ncreate selenium session error: %v", err)
		return
	}

	fmt.Println("\nSuccessful!")
	defer driver.Quit()
}
