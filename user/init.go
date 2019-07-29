package user

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/personal_journal_app/constant"
	"github.com/personal_journal_app/lib"
	"io"
	"strconv"
)

//Global Variables to be initialized on app start
var UserDataMap map[string]UserData
var CurrUser UserData

func Init() {
	var option int
	var err error

	fileName := fmt.Sprintf("%s/%s%s", constant.PARENT_DIRECTORY, constant.FILE_PATH, constant.USER_FILE_NAME)

	LoadUserDataFile(fileName)

	fmt.Println("Login or SignUp. For Login, enter 1 and for SignUp, enter 2. For EXIT, enter 3")
	fmt.Scanln(&option)

	switch option {
	case 1:
		if err = LoginHandler(); err != nil {
			fmt.Println("Failed to Login. Error: ", err)
			return
		}
	case 2:
		if err = SignupHandler(); err != nil {
			fmt.Println("Failed to Login. Error ", err)
			return
		}
	case 3:
		fmt.Println("Exiting....... ")
		return
	default:
		fmt.Println("Invalid options entered ", option)
		return
	}

}

func LoadUserDataFile(fileName string) {
	//Initialize global UserDataMap
	UserDataMap = make(map[string]UserData)

	csvFile, err := lib.OpenFile(fileName)
	if err != nil {
		fmt.Printf("[LoadUserDataFile] Error: %+v\n", err)
	}

	defer csvFile.Close()
	//Checking if userName exists or not
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error in reading csv records: ", err)
			return
		}

		for i, _ := range record {
			record[i], err = lib.Decrypt(constant.PASSKEY, record[i])
			if err != nil {
				fmt.Println("Error in decrypting csv record: ", err)
			}
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			fmt.Println("Invalid User ID ", err)
		}

		var row UserData
		row = UserData{
			ID:       id,
			UserName: record[1],
			Password: record[2],
			Name:     record[3],
			Email:    record[4],
			Mobile:   record[5],
		}

		UserDataMap[row.UserName] = row
	}

	defer csvFile.Close()

	return
}
