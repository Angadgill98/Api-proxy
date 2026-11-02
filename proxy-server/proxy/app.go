package proxy

import "database/sql"


func (Application) InsertApp(db *sql.DB, name string, url string) (int64, error) {
    query := `INSERT INTO application (name, url) VALUES (?, ?)`

    result, err := db.Exec(query, name, url)
    if err != nil {
        return 0, err
    }

    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    return id, nil
}

func (Application)UpdateApp(db *sql.DB, name string, newName string, url string) error {
	query := `UPDATE application SET name = ?, url = ? WHERE name = ?`
	_, err := db.Exec(query, newName, url, name)
	return err
}

func (Application) DeleteApp(db *sql.DB, id int64) error {
    query := `DELETE FROM application WHERE id = ?`
    _, err := db.Exec(query, id)
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

func GetAllApps(db *sql.DB) ([]Application, error) {
    apps := make([]Application, 0)

    query := `SELECT id, name, url FROM application`

    rows, err := db.Query(query)
    if err != nil {
        return apps, err
    }
    defer rows.Close()

    for rows.Next() {
        var app Application

        if err := rows.Scan(&app.ID, &app.Name, &app.URL); err != nil {
            return apps, err
        }

        apps = append(apps, app)
    }

    if err := rows.Err(); err != nil {
        return apps, err
    }

    return apps, nil
}