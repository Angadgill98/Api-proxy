package proxy

import (
	"database/sql"
	"encoding/json"
)

func InsertProxyInfo(db *sql.DB, info ProxyInfo) (int64, error) {
	requestJSON, err := json.Marshal(info.Request)
	if err != nil {
		return 0, err
	}

	responseJSON, err := json.Marshal(info.Response)
	if err != nil {
		return 0, err
	}

	query := `
		INSERT INTO proxy_info (
			route,
			request,
			response,
			started_at,
			completed_at,
			duration,
			error
		) VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	result, err := db.Exec(
		query,
		info.Route,
		string(requestJSON),
		string(responseJSON),
		info.StartedAt,
		info.CompletedAt,
		int64(info.Duration),
		info.Error,
	)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}


func UpsertProxyInfo(db *sql.DB, info ProxyInfo) (int64, error) {
	requestJSON, err := json.Marshal(info.Request)
	if err != nil {
		return 0, err
	}

	responseJSON, err := json.Marshal(info.Response)
	if err != nil {
		return 0, err
	}

	query := `
		INSERT INTO proxy_info (
			id, route, request, response, started_at, completed_at, duration, error
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT(id) DO UPDATE SET
			route = excluded.route,
			request = excluded.request,
			response = excluded.response,
			started_at = excluded.started_at,
			completed_at = excluded.completed_at,
			duration = excluded.duration,
			error = excluded.error;
	`

	_, err = db.Exec(
		query,
		info.ID,
		info.Route,
		string(requestJSON),
		string(responseJSON),
		info.StartedAt,
		info.CompletedAt,
		int64(info.Duration),
		info.Error,
	)
	if err != nil {
		return 0, err
	}

	return info.ID, nil
}

func ProxyInfoExists(db *sql.DB, id int64) (bool, error) {
	var exists bool

	err := db.QueryRow(
		`SELECT EXISTS(SELECT 1 FROM proxy_info WHERE id = ?)`,
		id,
	).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}