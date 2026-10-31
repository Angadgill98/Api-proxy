package proxy

import "database/sql"


func (Application)InsertApp(db *sql.DB, name string, url string) error {
	query := `INSERT INTO application (name, url) VALUES (?, ?)`
	_, err := db.Exec(query, name, url)
	return err
}

func (Application)UpdateApp(db *sql.DB, name string, newName string, url string) error {
	query := `UPDATE application SET name = ?, url = ? WHERE name = ?`
	_, err := db.Exec(query, newName, url, name)
	return err
}

 

func GetAppByName(db *sql.DB, name string) (*Application, error) {
    query := `SELECT id, name, url FROM application WHERE name = ?`

    var app Application
    err := db.QueryRow(query, name).Scan(&app.ID, &app.Name, &app.URL)
    if err != nil {
        return nil, err
    }

    return &app, nil
}   