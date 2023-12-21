package repository

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mentoria/src/user/postgres/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateUser(t *testing.T) {
	mockUuid := uuid.New()
	contactId := uuid.New()

	cases := []struct {
		name     string
		req      *model.User
		mockFunc func(sqlMock sqlmock.Sqlmock)
		wantErr  error
		want     *model.User
	}{
		{
			name: "Success creating new user",
			req: &model.User{
				ID:         mockUuid,
				FullName:   "FullName",
				SocialName: "SocialName",
				CPF:        "CPF",
				ContactID:  contactId,
			},
			mockFunc: func(sqlMock sqlmock.Sqlmock) {
				sqlMock.ExpectBegin()
				sqlMock.ExpectQuery("INSERT INTO").
					WithArgs("FullName", "SocialName", "CPF", contactId, mockUuid).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(mockUuid))
				sqlMock.ExpectCommit()
			},
			wantErr: nil,
			want: &model.User{
				ID:         mockUuid,
				FullName:   "FullName",
				SocialName: "SocialName",
				CPF:        "CPF",
				ContactID:  contactId,
			},
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

			d := NewUser(db)

			tc.mockFunc(mockSql)
			response, err := d.CreateUser(tc.req)

			assert.Equal(t, tc.want, response)
			assert.Equal(t, tc.wantErr, err)
			assert.Nil(t, mockSql.ExpectationsWereMet())
		})
	}
}
