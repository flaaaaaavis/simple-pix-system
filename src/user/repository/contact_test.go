package repository

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"projeto.com/src/user/model"
	"testing"

	sqlMock "github.com/DATA-DOG/go-sqlmock"
)

func TestCreateContact(t *testing.T) {
	mockUuid := uuid.New().String()

	cases := []struct {
		name     string
		req      *model.Contact
		mockFunc func(sqlMock sqlMock.Sqlmock)
		wantErr  error
		want     *model.Contact
	}{
		{
			name: "Sucess creating new contact",
			req: &model.Contact{
				PhoneNumber: "(87) 98888-8888",
				Email:       "email@email.com",
			},
			mockFunc: func(sqlMock sqlMock.Sqlmock) {
				sqlMock.ExpectQuery("INSERT INTO").WithArgs("phone_number", "email").WillReturnRows(sqlMock.NewRows([]string{"phone_number", "email"}))
			},
			wantErr: nil,
			want: &model.Contact{
				ID:          mockUuid,
				PhoneNumber: "(87) 98888-8888",
				Email:       "email@email.com",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			conn, mockSql, err := sqlMock.New()
			assert.Nil(t, err)
			defer conn.Close()

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

	/*
		mock
		assert
		execute
	*/

}
