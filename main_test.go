package main

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert" // from Practicum
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
	// мой вариант	require.Equal(t, got.ID, clientID)
	// из практикума
	assert.Equal(t, clientID, got.ID)
	assert.NotEmpty(t, got.FIO)
	assert.NotEmpty(t, got.Login)
	assert.NotEmpty(t, got.Birthday)
	assert.NotEmpty(t, got.Email)
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

	//	require.Empty(t, got)
	//
	// from Practicum
	assert.Empty(t, got.ID)
	assert.Empty(t, got.FIO)
	assert.Empty(t, got.Login)
	assert.Empty(t, got.Birthday)
	assert.Empty(t, got.Email)
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
	id, err := insertClient(db, cl)
	cl.ID = int(id)

	require.NoError(t, err)
	require.NotEmpty(t, cl.ID)

	got, err := selectClient(db, cl.ID)
	require.NoError(t, err)

	assert.Equal(t, cl.ID, got.ID)

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

	require.NoError(t, err)
	require.NotEmpty(t, id)

	_, err = selectClient(db, id)
	require.NoError(t, err)

	err = deleteClient(db, id)
	require.NoError(t, err)

	_, err = selectClient(db, id)
	require.Equal(t, sql.ErrNoRows, err)
}
