package hangmanweb

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func InitGlobalValue(globaldata GlobalInfo) GlobalInfo{

	globalDatabase, err := os.OpenFile("./server/database/global.csv", os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
	}

	defer globalDatabase.Close()

	csvReaderGlobalDB := csv.NewReader(globalDatabase)
	getDataGlobalDB, err := csvReaderGlobalDB.ReadAll()

	if err != nil {
		fmt.Println(err)
	}

	if len(getDataGlobalDB) != 0 {
		globaldata.DeadSanta, _ = strconv.Atoi(getDataGlobalDB[0][0])
		globaldata.SaveSanta, _ = strconv.Atoi(getDataGlobalDB[0][1])
	} else {

		csvWriterGlobalDB := csv.NewWriter(globalDatabase)

		newData := []string{"0", "0"}
		err = csvWriterGlobalDB.Write(newData)
		if err != nil {
			fmt.Println(err)
		}
		defer csvWriterGlobalDB.Flush()

		globaldata.DeadSanta = 0
		globaldata.SaveSanta = 0
	}

	if globaldata.SaveSanta+globaldata.DeadSanta != 0 {
		globaldata.Ratio = globaldata.SaveSanta * 100 / (globaldata.SaveSanta + globaldata.DeadSanta)
	} else {
		globaldata.Ratio = 50
	}

	globaldata.Total = globaldata.DeadSanta + globaldata.SaveSanta
	return globaldata
}
