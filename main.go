package main

import (
	"github.com/personal_journal_app/constant"
	"github.com/personal_journal_app/journal"
	"github.com/personal_journal_app/lib"
	"github.com/personal_journal_app/user"
)

func main() {
	constant.PARENT_DIRECTORY = lib.GetParentDirectory()

	user.Init()
	journal.Init()
}
