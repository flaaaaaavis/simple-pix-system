package repository

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mentoria/src/pix/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreatePix(t *testing.T) {
	mockUuid := uuid.New()
	userId := uuid.New()
	bankAccountId := uuid.New()

	cases := []struct {
		name     string
		req      *model.Pix
		mockFunc func(sqlMock sqlmock.Sqlmock)
		wantErr  error
		want     *model.Pix
	}{
		{
			name: "Success creating new Pix",
			req: &model.Pix{
				ID:            mockUuid,
				UserID:        userId,
				BankAccountID: bankAccountId,
				Balance:       decimal.NewFromFloat(500.00),
			},
			mockFunc: func(sqlMock sqlmock.Sqlmock) {
				sqlMock.ExpectBegin()
				sqlMock.ExpectQuery("INSERT INTO").
					WithArgs(userId, bankAccountId, decimal.NewFromFloat(500.00), mockUuid).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(mockUuid))
				sqlMock.ExpectCommit()
			},
			wantErr: nil,
			want: &model.Pix{
				ID:            mockUuid,
				UserID:        userId,
				BankAccountID: bankAccountId,
				Balance:       decimal.NewFromFloat(500.00),
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

			d := NewPix(db)

			tc.mockFunc(mockSql)
			response, err := d.CreatePix(tc.req)

			assert.Equal(t, tc.want, response)
			assert.Equal(t, tc.wantErr, err)
			assert.Nil(t, mockSql.ExpectationsWereMet())
		})
	}
}
