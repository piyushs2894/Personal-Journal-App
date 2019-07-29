package journal

import (
	"fmt"
	"github.com/personal_journal_app/constant"
	"github.com/personal_journal_app/lib"
	"github.com/personal_journal_app/user"
	"os"
	"path/filepath"
	"testing"
)

func TestCreateEntryHandler(t *testing.T) {
	wd := lib.GetParentDirectory()
	parent := filepath.Dir(wd)
	constant.PARENT_DIRECTORY = parent

	fileName := fmt.Sprintf("%s/files/testJournalData.txt", parent)

	tmpfile, err := lib.OpenFile(fileName)
	if err != nil {
		t.Errorf("Failed in TestCreateEntryHandler. Error: %+v", err)
	}

	oldStdin := os.Stdin
	defer func() {
		os.Stdin = oldStdin
	}()
	os.Stdin = tmpfile

	user := user.UserData{
		ID:       1,
		UserName: "Tom",
		Password: "Pass",
		Name:     "Tom Martin",
		Email:    "tom@gmail.com",
		Mobile:   "9876557899",
	}

	if err = CreateEntryHandler(user); err != nil {
		t.Errorf("Test Failed CreateEntryHandler. Error: %+v", err)
	}

	t.Log("Journal Entry test successfully passed")
}

func TestCreateEntry(t *testing.T) {
	wd := lib.GetParentDirectory()
	parent := filepath.Dir(wd)
	constant.PARENT_DIRECTORY = parent

	user := user.UserData{
		ID:       1,
		UserName: "Tom",
		Password: "Pass",
		Name:     "Tom Martin",
		Email:    "tom@gmail.com",
		Mobile:   "9876557899",
	}
	journal := JournalEntry{
		CreatedAt:  "2006/01/02 15:04:05",
		JournalLog: "TEST LOG",
	}

	if err := journal.CreateEntry(user); err != nil {
		t.Errorf("Failed in testing CreateEntry. Error:%+v", err)
	}

	t.Log("Journal Entry test successfully passed")
}

func TestDisplayEntries(t *testing.T) {
	wd := lib.GetParentDirectory()
	parent := filepath.Dir(wd)
	constant.PARENT_DIRECTORY = parent

	user := user.UserData{
		ID:       1,
		UserName: "Tom",
		Password: "Pass",
		Name:     "Tom Martin",
		Email:    "tom@gmail.com",
		Mobile:   "9876557899",
	}

	if err := DisplayEntries(user); err != nil {
		t.Errorf("Test Failed DisplayEntries. Error: %+v", err)
	}

	t.Log("DisplayEntries test successfully passed")
}
