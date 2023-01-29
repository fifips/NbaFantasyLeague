package database

import (
	"github.com/google/uuid"
	"reflect"
	"testing"
	"time"
)

func TestCreateOrUpdateActivationCode(t *testing.T) {
	type args struct {
		aC ActivationCode
	}

	newUuid, err := uuid.NewRandom()
	if err != nil {
		t.Error(err)
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Positive: Create",
			args: args{
				aC: ActivationCode{
					Code:    newUuid,
					UserId:  4,
					Expires: time.Now().Add(time.Minute * 15),
				},
			},
			wantErr: false,
		},
		{
			name: "Positive: Update",
			args: args{
				aC: ActivationCode{
					Code:    newUuid,
					UserId:  1,
					Expires: time.Now().Add(time.Minute * 15),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateOrUpdateActivationCode(tt.args.aC); (err != nil) != tt.wantErr {
				t.Errorf("CreateOrUpdateActivationCode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteActivationCodeByUserId(t *testing.T) {
	type args struct {
		userId int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Positive",
			args: args{
				userId: 2,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteActivationCodeByUserId(tt.args.userId); (err != nil) != tt.wantErr {
				t.Errorf("DeleteActivationCodeByUserId() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetActivationCodeByUserId(t *testing.T) {
	type args struct {
		userId int
	}

	dbUuid, err := uuid.Parse("3ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err != nil {
		t.Error(err)
	}
	dbExpires, err := time.Parse("2006-01-02 15:04:05", "2022-01-03 12:00:00")
	if err != nil {
		t.Error(err)
	}
	tests := []struct {
		name    string
		args    args
		want    ActivationCode
		wantErr bool
	}{
		{
			name: "Positive",
			args: args{
				userId: 3,
			},
			want: ActivationCode{
				Code:    dbUuid,
				UserId:  3,
				Expires: dbExpires,
			},
			wantErr: false,
		},
		{
			name: "Negative",
			args: args{
				userId: 5,
			},
			want:    ActivationCode{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetActivationCodeByUserId(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetActivationCodeByUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetActivationCodeByUserId() got = %v, want %v", got, tt.want)
			}
		})
	}
}
