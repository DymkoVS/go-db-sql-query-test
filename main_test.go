package main

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
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
	// Получить объект клиента функцией selectClient(); если функция вернула ошибку — завершить тест.
	got, err := selectClient(db, clientID)
	require.NoError(t, err)

	// мой вариант
	// require.Equal(t, got.ID, clientID)
	//	require.NotEmpty(t, got.FIO)
	//	require.NotEmpty(t, got.Login)
	//	require.NotEmpty(t, got.Birthday)
	//	require.NotEmpty(t, got.Email)
	// из практикума

	// Проверить, что поле ID объекта Client, который вернёт функция selectClient(), совпадает с идентификатором в переменной clientID.
	assert.Equal(t, clientID, got.ID)

	// Проверить, что остальные поля не пустые. Если какие-то из полей не прошли проверку, тест не должен завершиться.
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
	// Вызвать функцию selectClient().
	got, err := selectClient(db, clientID)

	// Проверить, что функция вернула ошибку и ошибка равна sql.ErrNoRows. Иначе завершить тест.
	require.Equal(t, sql.ErrNoRows, err)

	// these are my lines
	//	require.Empty(t, got.ID)
	//	require.Empty(t, got.FIO)
	//	require.Empty(t, got.Login)
	//	require.Empty(t, got.Birthday)
	//	require.Empty(t, got.Email)
	//
	// from Practicum
	//Проверить, что все поля объекта Client пустые. Если какие-то из полей не пустые, тест не должен завершиться.
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
	//	Добавить запись в таблицу функцией insertClient(). Сохранить полученный идентификатор в поле ID в переменной cl.
	id, err := insertClient(db, cl)
	cl.ID = int(id)

	// Проверить, что функция вернула не пустой идентификатор и пустую ошибку. Иначе завершить тест.
	require.NoError(t, err)
	require.NotEmpty(t, cl.ID)

	// Функцией selectClient() получить объект Client по идентификатору. Проверить, что функция вернула пустую ошибку. Иначе завершить тест.
	got, err := selectClient(db, cl.ID)
	require.NoError(t, err)

	//	Проверить, что значения всех полей полученного объекта совпадают со значениями полей объекта в переменной cl.
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
	// Добавить запись в таблицу функцией insertClient().
	id, err := insertClient(db, cl)

	// Проверить, что функция вернула не пустой идентификатор и пустую ошибку. Иначе завершить тест.
	require.NoError(t, err)
	require.NotEmpty(t, id)

	// Получить объект клиента функцией selectClient(). Если функция вернула ошибку, завершить тест.
	_, err = selectClient(db, id)
	require.NoError(t, err)

	// Удалить запись функцией deleteClient(). Если функция вернула ошибку, завершить тест.
	err = deleteClient(db, id)
	require.NoError(t, err)

	// Получить объект клиента функцией selectClient(). Проверить, что функция вернула ошибку и ошибка равна sql.ErrNoRows. Иначе завершить тест.
	_, err = selectClient(db, id)
	require.Equal(t, sql.ErrNoRows, err)
}

/* func Test_deleteClient_LastRow(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()

	res, err := dbstmt.Exec(db)

	lid, err := res.LastInsertId()

	fmt.Println(lid)

} */
