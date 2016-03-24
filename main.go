package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aalpern/luminosity"
)

func main() {
	err := aggregate()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func aggregate() error {
	merged := luminosity.NewCatalog()
	for _, path := range os.Args[1:] {
		c, err := luminosity.OpenCatalog(path)
		if err != nil {
			return err
		}
		err = c.Load()
		if err != nil {
			return err
		}
		merged.Merge(c)
	}
	print(merged)
	return nil
}

func print(data interface{}) {
	js, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(js))
}
