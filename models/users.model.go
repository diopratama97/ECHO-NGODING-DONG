package models

import (
	"net/http"
	"test-echo/db"
)

type Users struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Telp    string `json:"telp"`
}

func FecthAllUsers() (Response, error) {
	var data Users
	var arrObj []Users
	var res Response

	conn := db.CreateConn()

	sqlQuery := "SELECT * FROM users"

	rows, err := conn.Query(sqlQuery)

	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&data.Id, &data.Name, &data.Address, &data.Telp)
		if err != nil {
			return res, err
		}

		arrObj = append(arrObj, data)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrObj

	return res, nil
}

func CreateUsers(name, address, telp string) (Response, error) {
	var res Response

	conn := db.CreateConn()

	sqlQuery := "INSERT users (name,address,telp) VALUES (?,?,?)"

	stmt, err := conn.Prepare(sqlQuery)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, address, telp)

	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_id": lastInsertedId,
	}

	return res, nil
}
