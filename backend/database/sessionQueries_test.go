package database

import (
	"github.com/google/uuid"
	"reflect"
	"testing"
)

func TestCreateOrUpdateSession(t *testing.T) {
	type args struct {
		s Session
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
				s: Session{
					Id:     newUuid,
					UserId: 4,
				},
			},
			wantErr: false,
		},
		{
			name: "Positive: Update",
			args: args{
				s: Session{
					Id:     newUuid,
					UserId: 1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateOrUpdateSession(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("CreateOrUpdateSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteSessionByUserId(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Positive",
			args: args{
				id: 2,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteSessionByUserId(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteSessionByUserId() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetSessionById(t *testing.T) {
	type args struct {
		id uuid.UUID
	}

	newUuid, err := uuid.Parse("3ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err != nil {
		t.Error(err)
	}

	tests := []struct {
		name    string
		args    args
		want    *Session
		wantErr bool
	}{
		{
			name: "Positive",
			args: args{
				id: newUuid,
			},
			want: &Session{
				Id:     newUuid,
				UserId: 3,
			},
			wantErr: false,
		},
		{
			name: "Negative",
			args: args{
				id: uuid.UUID{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSessionById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSessionById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSessionById() got = %v, want %v", got, tt.want)
			}
		})
	}
}
