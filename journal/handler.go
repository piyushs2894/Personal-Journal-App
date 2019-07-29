package journal

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/personal_journal_app/constant"
	"github.com/personal_journal_app/lib"
	"github.com/personal_journal_app/user"
	"io"
	"os"
	"strings"
	"time"
)

func CreateEntryHandler(currUser user.UserData) error {
	var err error
	entry := JournalEntry{}
	fmt.Println("Enter text to be logged")

	r := bufio.NewReader(os.Stdin)
	entry.JournalLog, err = r.ReadString('\n')
	if err != nil {
		return err
	}

	entry.JournalLog = strings.TrimRight(entry.JournalLog, "\n")
	entry.CreatedAt = time.Now().Format("2006/01/02 15:04:05")

	if err = entry.CreateEntry(currUser); err != nil {
		return err
	}

	return nil
}

func (entry JournalEntry) CreateEntry(currUser user.UserData) error {
	fileName := fmt.Sprintf("%s/%s%d.csv", constant.PARENT_DIRECTORY, constant.FILE_PATH, currUser.ID)

	allEntries, csvFile, err := getAllEntries(fileName)
	if err != nil {
		return err
	}

	defer csvFile.Close()

	//Removing oldest journal entry if MAX_JOURNAL_ENTRY_LIMIT is reached
	if len(allEntries) >= constant.MAX_JOURNAL_ENTRY_LIMIT {
		allEntries = allEntries[1:]
	}
	allEntries = append(allEntries, entry)

	//Remove all content of a file and rewrite again
	if err = os.Truncate(fileName, 0); err != nil {
		return fmt.Errorf("Error in removing content of a file. Error: %+v", err)
	}

	writer := csv.NewWriter(csvFile)
	for _, entry := range allEntries {
		var encryptedRecord []string
		date, err := lib.Encrypt([]byte(constant.PASSKEY), []byte(entry.CreatedAt))
		if err != nil {
			return err
		}
		encryptedRecord = append(encryptedRecord, date)

		logEntry, err := lib.Encrypt([]byte(constant.PASSKEY), []byte(entry.JournalLog))
		if err != nil {
			return err
		}
		encryptedRecord = append(encryptedRecord, logEntry)
		if err = lib.WriteFile(writer, encryptedRecord); err != nil {
			return err
		}
	}

	fmt.Println("Journal Entry successfully logged")

	return nil
}

func DisplayEntries(currUser user.UserData) error {
	fileName := fmt.Sprintf("%s/%s%d.csv", constant.PARENT_DIRECTORY, constant.FILE_PATH, currUser.ID)

	journalEntries, csvFile, err := getAllEntries(fileName)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	for _, entry := range journalEntries {
		fmt.Printf("%s - %s\n", entry.CreatedAt, entry.JournalLog)
	}

	return nil
}

func getAllEntries(fileName string) ([]JournalEntry, *os.File, error) {
	var allEntries []JournalEntry

	csvFile, err := lib.OpenFile(fileName)
	if err != nil {
		return allEntries, csvFile, err
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return allEntries, csvFile, fmt.Errorf("Error in reading csv records: %+v", err)
		}

		for i, _ := range record {
			record[i], err = lib.Decrypt(constant.PASSKEY, record[i])
			if err != nil {
				return allEntries, csvFile, err
			}
		}

		var row JournalEntry
		row = JournalEntry{
			CreatedAt:  record[0],
			JournalLog: record[1],
		}
		allEntries = append(allEntries, row)
	}

	return allEntries, csvFile, nil
}
