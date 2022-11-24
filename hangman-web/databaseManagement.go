package hangmanweb

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"net/http"
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
		globaldata.DeadSanta = AtoiWithoutErr(getDataGlobalDB[0][0])
		globaldata.SaveSanta = AtoiWithoutErr(getDataGlobalDB[0][1])
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

func UpdateGlobalValue(save bool, globaldata GlobalInfo) GlobalInfo{

	globalDatabase, err := os.OpenFile("./server/database/global.csv", os.O_RDONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
	}

	csvReaderGlobalDB := csv.NewReader(globalDatabase)
	getDataGlobalDB, err := csvReaderGlobalDB.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	if len(getDataGlobalDB) != 0 {
		globaldata.DeadSanta = AtoiWithoutErr(getDataGlobalDB[0][0])
		globaldata.SaveSanta = AtoiWithoutErr(getDataGlobalDB[0][1])
	} else {
		globaldata.DeadSanta = 0
		globaldata.SaveSanta = 0
	}
	globalDatabase.Close()

	globalDatabase, err = os.OpenFile("./server/database/global.csv", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
	}

	csvWriterGlobalDB := csv.NewWriter(globalDatabase)

	if save {
		globaldata.SaveSanta++
	} else {
		globaldata.DeadSanta++
	}

	newData := []string{strconv.Itoa(globaldata.DeadSanta), strconv.Itoa(globaldata.SaveSanta)}
	err = csvWriterGlobalDB.Write(newData)
	if err != nil {
		fmt.Println(err)
	}
	defer csvWriterGlobalDB.Flush()

	if globaldata.SaveSanta+globaldata.DeadSanta != 0 {
		globaldata.Ratio = globaldata.SaveSanta * 100 / (globaldata.SaveSanta + globaldata.DeadSanta)
	} else {
		globaldata.Ratio = 50
	}

	globaldata.Total = globaldata.DeadSanta + globaldata.SaveSanta

	return globaldata

}

func UpdateUserValue(win bool, w http.ResponseWriter, r *http.Request, sbUsersList ScoreboardData, gameLaunch map[string]Hangman) ScoreboardData {
	userDatabase, err := os.OpenFile("./server/database/users.csv", os.O_RDONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
	}

	csvReaderUsersDB := csv.NewReader(userDatabase)
	getDataUsersDB, err := csvReaderUsersDB.ReadAll()

	sbUsersList.UsersList = []User{}

	if len(getDataUsersDB) != 0 {
		for ligne, userIngetData := range getDataUsersDB {
			if userIngetData[0] == gameLaunch[CookieSession(w, r, gameLaunch)].PlayerName {
				if win {
					sbUsersList.UsersList = append(sbUsersList.UsersList, User{userIngetData[0], AtoiWithoutErr(userIngetData[2]) + 1, AtoiWithoutErr(userIngetData[3]), AtoiWithoutErr(userIngetData[4]) + 1})
					getDataUsersDB[ligne][2] = strconv.Itoa(AtoiWithoutErr(getDataUsersDB[ligne][2]) + 1)
					getDataUsersDB[ligne][4] = strconv.Itoa(AtoiWithoutErr(getDataUsersDB[ligne][4]) + 1)
				} else {
					sbUsersList.UsersList = append(sbUsersList.UsersList, User{userIngetData[0], AtoiWithoutErr(userIngetData[2]), AtoiWithoutErr(userIngetData[3]) + 1, AtoiWithoutErr(userIngetData[4]) + 1})
					getDataUsersDB[ligne][3] = strconv.Itoa(AtoiWithoutErr(getDataUsersDB[ligne][3]) + 1)
					getDataUsersDB[ligne][4] = strconv.Itoa(AtoiWithoutErr(getDataUsersDB[ligne][4]) + 1)
				}
			} else {
				sbUsersList.UsersList = append(sbUsersList.UsersList, User{userIngetData[0], AtoiWithoutErr(userIngetData[2]), AtoiWithoutErr(userIngetData[3]), AtoiWithoutErr(userIngetData[4])})
			}
		}
	}
	userDatabase.Close()

	userDatabase, err = os.OpenFile("./server/database/users.csv", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
	}

	csvWriterUsersDB := csv.NewWriter(userDatabase)

	err = csvWriterUsersDB.WriteAll(getDataUsersDB)
	if err != nil {
		fmt.Println(err)
	}
	defer csvWriterUsersDB.Flush()

	return sbUsersList
}
