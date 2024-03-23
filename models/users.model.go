package models

import (
	"net/http"
	"test-echo/db"

	"github.com/go-playground/validator/v10"
)

type Users struct {
	Id      int    `json:"id"`
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required" `
	Telp    string `json:"telp" validate:"required"`
}

func FecthAllUsers() (Response, error) {
	var data Users
	var arrObj []Users
	var res Response

	conn := db.CreateConn()

	sqlQuery := "SELECT id,name,address,telp FROM users"

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

	v := validator.New()

	usr := Users{
		Name:    name,
		Address: address,
		Telp:    telp,
	}

	err := v.Struct(usr)
	if err != nil {
		return res, err
	}

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

func UpdateUsers(id int, name string, address string, telp string) (Response, error) {
	var res Response

	conn := db.CreateConn()

	sqlQuery := "UPDATE users SET name = ?, address = ?, telp = ? WHERE id = ?"

	stmt, err := conn.Prepare(sqlQuery)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, address, telp, id)

	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"row_affected": rowsAffected,
	}

	return res, nil
}

func DeleteUsers(id int) (Response, error) {
	var res Response

	conn := db.CreateConn()

	sqlQuery := "DELETE FROM users WHERE id = ?"

	stmt, err := conn.Prepare(sqlQuery)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = rowAffected

	return res, nil
}

func DetailUsers(id int) (Response, error) {
	var res Response
	var obj Users

	conn := db.CreateConn()

	sqlQuery := "SELECT id,name,address,telp FROM users WHERE id = ?"

	err := conn.QueryRow(sqlQuery, id).Scan(&obj.Id, &obj.Name, &obj.Address, &obj.Telp)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = obj

	return res, nil
}
