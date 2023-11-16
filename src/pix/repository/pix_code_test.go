package repository

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"projeto.com/src/pix/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreatePixCode(t *testing.T) {
	mockUuid := uuid.New()
	pixId := uuid.New()

	cases := []struct {
		name     string
		req      *model.PixCode
		mockFunc func(sqlMock sqlmock.Sqlmock)
		wantErr  error
		want     *model.PixCode
	}{
		{
			name: "Success creating new PixCode",
			req: &model.PixCode{
				ID:    mockUuid,
				PixID: pixId,
				Type:  model.PixTypeEmail,
				Code:  "Code",
			},
			mockFunc: func(sqlMock sqlmock.Sqlmock) {
				sqlMock.ExpectBegin()
				sqlMock.ExpectQuery("INSERT INTO").
					WithArgs(pixId, model.PixTypeEmail, "Code", mockUuid).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(mockUuid))
				sqlMock.ExpectCommit()
			},
			wantErr: nil,
			want: &model.PixCode{
				ID:    mockUuid,
				PixID: pixId,
				Type:  model.PixTypeEmail,
				Code:  "Code",
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

			d := NewPixCode(db)

			tc.mockFunc(mockSql)
			response, err := d.CreatePixCode(tc.req)

			assert.Equal(t, tc.want, response)
			assert.Equal(t, tc.wantErr, err)
			assert.Nil(t, mockSql.ExpectationsWereMet())
		})
	}
}
