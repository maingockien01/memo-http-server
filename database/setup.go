package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"webserver/filesystem"
)

func main() {
	db, _ := sql.Open("sqlite3", "./data.db")
	defer db.Close()

	createTable(db)

	memo1 := filesystem.Memo{
		LastEditedBy: "Kien",
		Content:      "Hi this is a sample 1",
	}
	memo2 := filesystem.Memo{
		LastEditedBy: "Kien",
		Content:      "Hi this is a sample 2",
	}
	memo3 := filesystem.Memo{
		LastEditedBy: "Kien",
		Content:      "Hi this is a sample 3",
	}

	filesystem.InsertMemo(db, memo1)
	filesystem.InsertMemo(db, memo2)
	filesystem.InsertMemo(db, memo3)

	memos, err := filesystem.GetAllMemo(db)

	if err != nil {
		panic(err)
	}

	text, _ := json.Marshal(memos)
	fmt.Println(string(text))
	filesystem.UpdateMemo(db, filesystem.Memo{Id: 1, LastEditedBy: "KienAgain", Content: "New content"})

	memos, err = filesystem.GetAllMemo(db)

	if err != nil {
		panic(err)
	}
	text, _ = json.Marshal(memos)
	fmt.Println(string(text))

	filesystem.DeleteMemo(db, 1)
	memos, err = filesystem.GetAllMemo(db)

	if err != nil {
		panic(err)
	}

	text, _ = json.Marshal(memos)
	fmt.Println(string(text))
}

func createTable(db *sql.DB) {
	createMemoTable := `CREATE TABLE memo (
		"Id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"LastEditedBy"	varchar(256),
		"Content" TEXT
	)`

	statement, err := db.Prepare(createMemoTable)
	if err != nil {
		panic(err)
	}

	statement.Exec()
}
