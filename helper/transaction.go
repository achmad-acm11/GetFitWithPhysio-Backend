package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()

	if err != nil {
		errRollback := tx.Rollback()
		HandleError(errRollback)
		panic(err)
	} else {
		errCommit := tx.Commit()
		HandleError(errCommit)
	}
}
