# Personal Journal Application

Personal Journal App is all about storing personal journal logs. It allows maximum of 10 user accounts creation and maximum of 50 journal entries.

## Functionalities

Functionalities it serves - 
```
1) It allows user account creation upto limit MAX_USER_LIMIT in constant.go file. Modification of user data is not allowed.
2) It allows creation of new journal entries upto MAX_JOURNAL_ENTRY_LIMIT in constant.go file.
3) User can see his recent journal entries upto MAX_JOURNAL_ENTRY_LIMIT.
4) All the data in files is encrypted using aes encryption. So User won't be able to read data directly from files.
```

## Assumptions

Assumptions taken in developing this application - 
```
1) Modification of user data is not allowed.
2) Assumed user is executing all commands in project main directory.
3) Allows maximum 10 user accounts creation. For modifying the limit, change MAX_USER_LIMIT in constant.go
4) Allows maximum 50 journal entries for a single user. For modifying the limit, change MAX_JOURNAL_ENTRY_LIMIT in constant.go
5) Validations on user data is not done.
```

## Start

Start application by running below command
```
1) Install go and set GOPATH, for which you can use https://github.com/golang/go/wiki/SettingGOPATH
2) This project should be in this path `$GOPATH/src/github.com`
3) Go to project directoy and run command `go run main.go`
```

## Test

To run execution of test functions. As our application involved user input, so for this mock data is prepared in testJournalData.txt, testLoginData.txt, testSignupData.txt files.

```
1) Go to project main directory.
2) Run this command `go test -v -race ./...`
3) It will perform testing of all test cases written in <file_name>_test.go.
4) On successful completion of all test cases, new files user.csv and <userId>.csv will be generated.
```