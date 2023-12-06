package repository

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	model "mentoria/src/pix/model/postgres"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateBankAccount(t *testing.T) {
	mockUuid := uuid.New()

	cases := []struct {
		name     string
		req      *model.BankAccount
		mockFunc func(sqlMock sqlmock.Sqlmock)
		wantErr  error
		want     *model.BankAccount
	}{
		{
			name: "Success creating new BankAccount",
			req: &model.BankAccount{
				ID:            mockUuid,
				BankCode:      "BankCode",
				BankName:      "BankName",
				BankBranch:    "BankBranch",
				AccountNumber: "AccountNumber",
			},
			mockFunc: func(sqlMock sqlmock.Sqlmock) {
				sqlMock.ExpectBegin()
				sqlMock.ExpectQuery("INSERT INTO").
					WithArgs("BankCode", "BankName", "BankBranch", "AccountNumber", mockUuid).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(mockUuid))
				sqlMock.ExpectCommit()
			},
			wantErr: nil,
			want: &model.BankAccount{
				ID:            mockUuid,
				BankCode:      "BankCode",
				BankName:      "BankName",
				BankBranch:    "BankBranch",
				AccountNumber: "AccountNumber",
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
					log.Fatalf("Error closing connection")
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

			d := NewBankAccount(db)

			tc.mockFunc(mockSql)
			response, err := d.CreateBankAccount(tc.req)

			assert.Equal(t, tc.want, response)
			assert.Equal(t, tc.wantErr, err)
			assert.Nil(t, mockSql.ExpectationsWereMet())
		})
	}
}
