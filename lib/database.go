package lib

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

func PrepareDatabase() *sql.DB {
	if _, folder_check_err := os.Stat("database"); os.IsNotExist(folder_check_err) {
		folder_create_err := os.Mkdir("database", 0755)
		if folder_create_err != nil {
			log.Fatal("Folder Create Error:", folder_create_err)
		}
	}
	if _, file_check_err := os.Stat(DB_PATH); os.IsNotExist(file_check_err) {
		_, file_create_err := os.Create(DB_PATH)
		if file_create_err != nil {
			log.Fatal("DB File Create Error:", file_create_err)
		}
	}
	DB, db_open_err := sql.Open("sqlite", DB_PATH)
	if db_open_err != nil {
		log.Fatal("Database Open Error:", db_open_err)
	}
	return DB
}

func HandleMigrations(DB *sql.DB) {
	if _, table_create_err := DB.Exec(`CREATE TABLE IF NOT EXISTS Todos(
		id INTEGER,
		text TEXT,
		description TEXT,
		done BOOLEAN
	);`); table_create_err != nil {
		log.Fatal(table_create_err)
	}
}