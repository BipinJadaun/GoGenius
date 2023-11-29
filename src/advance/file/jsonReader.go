package file

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Blossom struct {
	Id                      string `json:"id"`
	Name                    string `json:"name"`
	Product_owner_firstname string `json:"product_owner_firstname"`
	Product_owner_lastname  string `json:"product_owner_lastname"`
	Portfolio_name          string `json:"portfolio_name"`
	Product_group_name      string `json:"product_group_name"`
	Support_group_name      string `json:"support_group_name"`
	State                   string `json:"state"`
}

func ReadJsonFile() {

	file, err := os.Open("/Users/Z009XLP/Downloads/applications.json")
	if err != nil {
		fmt.Printf("error while reading jason %v", err)
	}
	var data []Blossom
	byteValue, _ := ioutil.ReadAll(file)

	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		fmt.Printf("error while reading jason %v", err)
	}

	for _, app := range data {
		if app.State == "Registered" {
			fmt.Println(app)
		}
	}

}
