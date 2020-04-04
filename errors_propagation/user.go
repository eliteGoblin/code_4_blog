package main

import (
	"encoding/json"
	se "error_propagation/errors"
	"net/http"

	"github.com/juju/errors"
)

const (
	badUserID    = "baduserid"
	forbidden    = "forbidden"
	notFountInDB = "notfountindb"
	dbTimeout    = "dbtimeout"
)

func handleUser(w http.ResponseWriter, r *http.Request) error {
	userID := r.URL.Query().Get("userID")
	err := checkUserID(userID)

	if err != nil {
		return errors.Trace(err)
	}
	err = authorize(userID)
	if err != nil {
		return errors.Trace(err)
	}
	user, err := getUserInfo(userID)
	if err != nil {
		return errors.Trace(err)
	}
	bytes, err := json.Marshal(user)
	if err != nil {
		return errors.Trace(err)
	}
	_, err = w.Write(bytes)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

func checkUserID(userID string) error {
	if userID == badUserID {
		return errors.Trace(
			se.NewBadRequest("INVALID_USERID", "malformed userID provided: %s", badUserID))
	}
	return nil
}

func authorize(userID string) error {
	if userID == forbidden {
		return errors.Trace(
			se.NewForbidden("USER_NOT_ACTIVATED", "you should activate your account: %s first", userID))
	}
	return nil
}

type UserInfo struct {
	Name    string
	Address string
}

func getUserInfo(userID string) (UserInfo, error) {
	if userID == notFountInDB {
		return UserInfo{}, errors.Trace(
			se.NewNotFound("userID: %s not found", userID))
	}
	user, err := getFromDB(userID)
	if err != nil {
		return UserInfo{}, errors.Trace(err)
	}
	return user, nil
}

func getFromDB(userID string) (UserInfo, error) {
	if userID == dbTimeout {
		return UserInfo{}, errors.Trace(
			errors.New("db timeout"))
	}
	return UserInfo{
		Name:    "Frank.Sun",
		Address: "Gold Coast",
	}, nil
}
