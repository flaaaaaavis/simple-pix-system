package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"projeto.com/src/user/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateContact(t *testing.T) {
	mockUuid := uuid.New()

	cases := []struct {
		name     string
		req      *model.Contact
		mockFunc func(sqlMock sqlmock.Sqlmock)
		wantErr  error
		want     *model.Contact
	}{
		{
			name: "Success creating new contact",
			req: &model.Contact{
				ID:          mockUuid,
				PhoneNumber: "(87) 98888-8888",
				Email:       "email@email.com",
			},
			mockFunc: func(sqlMock sqlmock.Sqlmock) {
				sqlMock.ExpectBegin()
				sqlMock.ExpectQuery("INSERT INTO").
					WithArgs("(87) 98888-8888", "email@email.com", mockUuid).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(mockUuid))
				sqlMock.ExpectCommit()
			},
			wantErr: nil,
			want: &model.Contact{
				ID:          mockUuid,
				PhoneNumber: "(87) 98888-8888",
				Email:       "email@email.com",
			},
		},
		{
			name: "Error creating new contact",
			req: &model.Contact{
				ID:          mockUuid,
				PhoneNumber: "(87) 98888-8888",
				Email:       "email@email.com",
			},
			mockFunc: func(sqlMock sqlmock.Sqlmock) {
				sqlMock.ExpectBegin()
				sqlMock.ExpectQuery("INSERT INTO").
					WithArgs("(87) 98888-8888").
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(mockUuid))
				sqlMock.ExpectCommit()
			},
			wantErr: errors.Unwrap(&fmt.wrapError{msg: "Query 'INSERT INTO \"contacts\" (\"phone_number\",\"email\",\"id\") VALUES ($1,$2,$3) RETURNING \"id\"', arguments do not match: expected 1, but got 3 arguments; call to Rollback transaction, was not expected, next expectation is: ExpectedQuery => expecting Query, QueryContext or QueryRow which:\n  - matches sql: 'INSERT INTO'\n  - is with arguments:\n    0 - (87) 98888-8888\n  - should return rows:\n    row 0 - [89612574-08df-4954-87cb-3f754504be5f]", err: (*errors.errorString)(0xc000242480)}),
			want:    nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			conn, mockSql, err := sqlmock.New()
			assert.Nil(t, err)
			defer func(conn *sql.DB) {
				err := conn.Close()
				if err != nil {

				}
			}(conn)

			postgresConfig := postgres.New(postgres.Config{
				DriverName:           "postgres",
				DSN:                  "sqlMock_db",
				PreferSimpleProtocol: true,
				Conn:                 conn,
			})

			db, err := gorm.Open(postgresConfig, &gorm.Config{})
			assert.NoError(t, err)

			d := NewContact(db)

			tc.mockFunc(mockSql)
			response, err := d.CreateContact(tc.req)

			assert.Equal(t, tc.want, response)
			assert.Equal(t, tc.wantErr, err)
			assert.Nil(t, mockSql.ExpectationsWereMet())
		})
	}
}
