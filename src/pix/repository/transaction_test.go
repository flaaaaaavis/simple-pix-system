package repository

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"projeto.com/src/pix/model"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateTransaction(t *testing.T) {
	mockUuid := uuid.New()
	date := time.Now()
	senderId := uuid.New()
	receiverId := uuid.New()

	cases := []struct {
		name     string
		req      *model.Transaction
		mockFunc func(sqlMock sqlmock.Sqlmock)
		wantErr  error
		want     *model.Transaction
	}{
		{
			name: "Success creating new Transaction",
			req: &model.Transaction{
				ID:         mockUuid,
				Type:       model.TransactionTypePayment,
				Date:       date,
				Amount:     decimal.NewFromFloat(500.00),
				SenderID:   senderId,
				ReceiverID: receiverId,
				Status:     model.TransactionStatusDone,
			},
			mockFunc: func(sqlMock sqlmock.Sqlmock) {
				sqlMock.ExpectBegin()
				sqlMock.ExpectQuery("INSERT INTO").
					WithArgs(model.TransactionTypePayment, date, decimal.NewFromFloat(500.00), senderId, receiverId, model.TransactionStatusDone, mockUuid).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(mockUuid))
				sqlMock.ExpectCommit()
			},
			wantErr: nil,
			want: &model.Transaction{
				ID:         mockUuid,
				Type:       model.TransactionTypePayment,
				Date:       date,
				Amount:     decimal.NewFromFloat(500.00),
				SenderID:   senderId,
				ReceiverID: receiverId,
				Status:     model.TransactionStatusDone,
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
					fmt.Sprintf("Error closing connection")
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

			d := NewTransaction(db)

			tc.mockFunc(mockSql)
			response, err := d.CreateTransaction(tc.req)

			assert.Equal(t, tc.want, response)
			assert.Equal(t, tc.wantErr, err)
			assert.Nil(t, mockSql.ExpectationsWereMet())
		})
	}
}
