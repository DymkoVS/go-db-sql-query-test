package main

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"

	_ "modernc.org/sqlite"
)

func Test_SelectClient_WhenOk(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()

	clientID := 1

	// напиши тест здесь
	got, err := selectClient(db, clientID)
	require.NoError(t, err)
	require.Equal(t, got.ID, clientID)

}

func Test_SelectClient_WhenNoClient(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()

	clientID := -1
	// напиши тест здесь
	got, err := selectClient(db, clientID)
	require.Equal(t, sql.ErrNoRows, err)
	require.Empty(t, got)
}

func Test_InsertClient_ThenSelectAndCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}
	// напиши тест здесь
	id, _ := insertClient(db, cl)
	got, err := selectClient(db, id)
	require.Equal(t, sql.ErrNoRows, err)
	require.Empty(t, got)

}

func Test_InsertClient_DeleteClient_ThenCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	// напиши тест здесь
	id, err := insertClient(db, cl)
	err = deleteClient(db, id)
	require.NoError(t, err)

	got, err := selectClient(db, id)
	require.Equal(t, sql.ErrNoRows, err)
	require.Empty(t, got)
}
