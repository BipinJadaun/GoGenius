package file

import (
	"fmt"
	"log"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func ReadExcelFile() {
	vmaas, err := excelize.OpenFile("/Users/Z009XLP/Downloads/vmaasApps.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	apps := make(map[string]string)

	for i := 2; i <= 7307; i++ {
		cid, _ := vmaas.GetCellValue("Workingsheet", "A"+strconv.Itoa(i))
		app, _ := vmaas.GetCellValue("Workingsheet", "AE"+strconv.Itoa(i))
		fmt.Println(cid + ": " + app)
		apps[cid] = app
	}

	fmt.Println(len(apps))
}
