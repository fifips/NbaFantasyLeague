package database

import (
	"reflect"
	"testing"
)

func TestCreateUser(t *testing.T) {
	type args struct {
		u User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Positive",
			args: args{
				u: User{
					Email:    "user5@example.com",
					Password: []byte("password5"),
				},
			},
			wantErr: false,
		},
		{
			name: "Negative",
			args: args{
				u: User{
					Email:    "user1@example.com",
					Password: []byte("password1"),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateUser(tt.args.u); (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteUserById(t *testing.T) {
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
			if err := DeleteUserById(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUserById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetUserByEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "Positive",
			args: args{
				email: "user1@example.com",
			},
			want: &User{
				Id:       1,
				Email:    "user1@example.com",
				Password: []byte("password1"),
				IsActive: false,
			},
			wantErr: false,
		},
		{
			name: "Negative",
			args: args{
				email: "user6@example.com",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUserByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUserById(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "Positive",
			args: args{
				id: 3,
			},
			want: &User{
				Id:       3,
				Email:    "user3@example.com",
				Password: []byte("password3"),
				IsActive: true,
			},
			wantErr: false,
		},
		{
			name: "Negative",
			args: args{
				id: 6,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUserById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserById() got = %v, want %v", got, tt.want)
			}
		})
	}
}
