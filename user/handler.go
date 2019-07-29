package user

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/personal_journal_app/constant"
	"github.com/personal_journal_app/lib"
	"os"
	"strings"
	"time"
)

func LoginHandler() error {
	CurrUser = UserData{}

	fmt.Println("Please enter your username")
	fmt.Scanln(&CurrUser.UserName)

	fmt.Println("Please enter your password")
	fmt.Scanln(&CurrUser.Password)

	if _, ok := UserDataMap[CurrUser.UserName]; ok {
		if UserDataMap[CurrUser.UserName].Password == CurrUser.Password {
			CurrUser = UserDataMap[CurrUser.UserName]
			fmt.Println("\nLogin Successful")
			return nil
		}
	}
	return fmt.Errorf("Invalid Login Credentials")
}

func SignupHandler() error {
	lastId := len(UserDataMap)
	if lastId < constant.MAX_USER_LIMIT {
		CurrUser.signup()
		CurrUser.ID = lastId + 1
		if err := CurrUser.registerUser(); err != nil {
			return fmt.Errorf("Failed to Register User. Please try again. Error %+v", err)
		}
	} else {
		return fmt.Errorf("Max User Limit of %d Reached ", constant.MAX_USER_LIMIT)
	}
	return nil
}

func (newUser *UserData) signup() {
	// newUserData := &UserData{}
	var ok bool = true
	newUser.getUserName()
	count := 0
	//Check if username exists or not
	for ok && count < 3 {
		_, ok = UserDataMap[newUser.UserName]
		if !ok {
			break
		}
		fmt.Println("\nUser name already exists. Please choose different username")
		// newUserData = &UserData{}
		newUser.getUserName()
		count++
	}

	newUser.getUserPassword()
	newUser.getAdditionalUserDetails()

	UserDataMap[newUser.UserName] = *newUser
}

func (newUser *UserData) registerUser() error {
	fileName := fmt.Sprintf("%s/%s%s", constant.PARENT_DIRECTORY, constant.FILE_PATH, constant.USER_FILE_NAME)
	csvFile, err := lib.OpenFile(fileName)
	if err != nil {
		return err
	}

	defer csvFile.Close()

	var data []string
	data = append(data, fmt.Sprintf("%d", newUser.ID))
	data = append(data, newUser.UserName)
	data = append(data, newUser.Password)
	data = append(data, newUser.Name)
	data = append(data, newUser.Email)
	data = append(data, newUser.Mobile)
	data = append(data, time.Now().Format("2006/01/02 15:04:05"))
	data = append(data, time.Now().Format("2006/01/02 15:04:05"))

	//Using complete filepath as key in encryption
	var encryptedData []string
	w := csv.NewWriter(csvFile)
	for _, v := range data {
		d, err := lib.Encrypt([]byte(constant.PASSKEY), []byte(v))
		if err != nil {
			return err
		}
		encryptedData = append(encryptedData, d)
	}
	lib.WriteFile(w, encryptedData)

	return nil
}

func (newUser *UserData) getUserName() {
	var userName string
	fmt.Println("Please enter new username")
	fmt.Scanln(&userName)

	newUser.UserName = userName
}

func (newUser *UserData) getUserPassword() {
	var userPassword1 string
	var userPassword2 string
	count := 0

	fmt.Println("Please enter password")
	fmt.Scanln(&userPassword1)

	for userPassword1 != userPassword2 && count < 3 {
		fmt.Println("Please confirm your password")
		fmt.Scanln(&userPassword2)
		count++
		if userPassword1 != userPassword2 {
			fmt.Printf("Passwords mismatch. Please try again, Password1: %s, Password2: %s", userPassword1, userPassword2)
			continue
		}
	}

	newUser.Password = userPassword1
	return
}

func (newUser *UserData) getAdditionalUserDetails() {
	var name, email, mobile string

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter your Name")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid Name Input")
	}
	name = strings.TrimRight(name, "\n")

	fmt.Println("Please enter your Email")
	email, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid Email Input")
	}
	email = strings.TrimRight(email, "\n")

	fmt.Println("Please enter your Mobile number")
	mobile, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid mobile number Input")
	}
	mobile = strings.TrimRight(mobile, "\n")

	newUser.Name = name
	newUser.Email = email
	newUser.Mobile = mobile
}
