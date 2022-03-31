package filesystem

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDB(path string) *sql.DB {
	db, _ := sql.Open("sqlite3", path)
	return db
}

func GetDb() *sql.DB {
	return OpenDB("./database/data.db")
}

func InsertMemo(db *sql.DB, memo Memo) (generatedMemo Memo, err error) {
	InsertMemoSQL := `INSERT INTO memo (LastEditedBy, Content) VALUES (?, ?)`
	statement, err := db.Prepare(InsertMemoSQL)

	if err != nil {
		return
	}

	_, err = statement.Exec(memo.LastEditedBy, memo.Content)

	if err != nil {
		return
	}

	generatedMemo, err = GetLastMemoInserted(db)

	return
}

func GetLastMemoInserted(db *sql.DB) (memo Memo, err error) {

	GetLastSQL := `SELECT * FROM memo WHERE id = last_insert_rowid()`

	row, err := db.Query(GetLastSQL)

	if err != nil {
		return
	}

	for row.Next() {
		var Id int
		var LastEditedBy string
		var Content string
		row.Scan(&Id, &LastEditedBy, &Content)

		memo = Memo{
			Id:           Id,
			LastEditedBy: LastEditedBy,
			Content:      Content,
		}
	}

	return

}

func GetAllMemo(db *sql.DB) (memos []Memo, err error) {
	row, err := db.Query("SELECT * FROM memo")
	if err != nil {
		return memos, err
	}

	defer row.Close()

	for row.Next() {
		var Id int
		var LastEditedBy string
		var Content string
		row.Scan(&Id, &LastEditedBy, &Content)

		memo := Memo{
			Id:           Id,
			LastEditedBy: LastEditedBy,
			Content:      Content,
		}

		memos = append(memos, memo)
	}

	return memos, nil
}

func DeleteMemo(db *sql.DB, id int) error {
	deleteSQL := `DELETE FROM memo WHERE id = ?`
	statement, err := db.Prepare(deleteSQL)

	if err != nil {
		return err
	}

	_, err = statement.Exec(id)

	return err
}

func UpdateMemo(db *sql.DB, memo Memo) error {
	updateSQL := `UPDATE memo SET LastEditedBy = ?, Content = ? WHERE Id = ?`
	statement, err := db.Prepare(updateSQL)

	if err != nil {
		return err
	}

	res, err := statement.Exec(memo.LastEditedBy, memo.Content, memo.Id)

	affect, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if affect == 0 {
		return NoRowAffectedErr{}
	}

	return err
}

type NoRowAffectedErr struct{}

func (err NoRowAffectedErr) Error() string {
	return "No memo was updated!"
}
