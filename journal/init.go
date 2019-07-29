package journal

import (
	"fmt"
	"github.com/personal_journal_app/user"
)

func Init() {
	var err error
	var option int

	if user.CurrUser.ID == 0 {
		fmt.Println("Invalid User for journal entry")
		return
	}

INIT:
	fmt.Println("\nFor creating new journal entry, enter 1. For viewing previous entries, enter 2. For Exit, enter 3")
	fmt.Scanln(&option)

	switch option {
	case 1:
		if err = CreateEntryHandler(user.CurrUser); err != nil {
			fmt.Printf("Failed to create new entry. Error: %+v\n", err)
			return
		}
		goto INIT
	case 2:
		if err = DisplayEntries(user.CurrUser); err != nil {
			fmt.Printf("Failed to display all entries. Error: %+v\n", err)
			return
		}
		goto INIT
	case 3:
		fmt.Println("Exiting....... ")
		return
	default:
		fmt.Println("Invalid option entered, option: ", option)
		goto INIT
	}
}
