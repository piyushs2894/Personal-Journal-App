package user

import (
	"fmt"
	"github.com/personal_journal_app/constant"
	"github.com/personal_journal_app/lib"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadUserDataFile(t *testing.T) {
	wd := lib.GetParentDirectory()
	parent := filepath.Dir(wd)
	constant.PARENT_DIRECTORY = parent

	fileName := fmt.Sprintf("%s/%s%s", parent, constant.FILE_PATH, constant.USER_FILE_NAME)

	LoadUserDataFile(fileName)
	if UserDataMap == nil {
		t.Errorf("Failed to initialize user")
	}
	t.Log("Init test successfully passed")
}

//Testing SignupHandler function will cover all private functions i.e. getUserName(), getUserPassword(),
//getUserAdditionalDetails(). So no separate test cases for them
func TestSignupHandler(t *testing.T) {
	wd := lib.GetParentDirectory()
	parent := filepath.Dir(wd)
	constant.PARENT_DIRECTORY = parent

	fileName := fmt.Sprintf("%s/files/testSignupData.txt", parent)

	tmpfile, err := lib.OpenFile(fileName)
	if err != nil {
		t.Errorf("Error: %+v", err)
	}

	oldStdin := os.Stdin
	defer func() {
		os.Stdin = oldStdin
	}()
	os.Stdin = tmpfile

	if err := SignupHandler(); err != nil {
		t.Errorf("Failed in TestSignupHandler. Error:%+v", err)
	}
	if UserDataMap == nil {
		t.Errorf("Failed to load User Data")
	}
	if err := tmpfile.Close(); err != nil {
		t.Errorf("Error: %+v", err)
	}

	t.Logf("User Data Map successfully loaded: %+v", UserDataMap)
}

func TestLoginHandler(t *testing.T) {
	wd := lib.GetParentDirectory()
	parent := filepath.Dir(wd)
	constant.PARENT_DIRECTORY = parent

	fileName := fmt.Sprintf("%s/files/testLoginData.txt", parent)

	tmpfile, err := lib.OpenFile(fileName)
	if err != nil {
		t.Errorf("Error: %+v", err)
	}

	oldStdin := os.Stdin
	defer func() {
		os.Stdin = oldStdin
	}()
	os.Stdin = tmpfile

	LoginHandler()
	if CurrUser.UserName != "Tom" {
		t.Errorf("Failed to load User Data")
	}
	if err := tmpfile.Close(); err != nil {
		t.Logf("Error: %+v", err)
	}

	t.Logf("Test Login Handler successful: %+v", CurrUser)
}
