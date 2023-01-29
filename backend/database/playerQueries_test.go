package database

import (
	"reflect"
	"testing"
)

func TestCreatePlayer(t *testing.T) {
	type args struct {
		player Player
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Positive",
			args: args{
				player: Player{Id: "4"},
			},
			wantErr: false,
		},
		{
			name: "Negative: Duplicate id",
			args: args{
				player: Player{Id: "1"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreatePlayer(tt.args.player); (err != nil) != tt.wantErr {
				t.Errorf("CreatePlayer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeletePlayerById(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Positive",
			args:    args{id: "2"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeletePlayerById(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeletePlayerById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetPlayerById(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    Player
		wantErr bool
	}{
		{
			name: "Positive",
			args: args{
				id: "1",
			},
			want: Player{
				Id: "1",
			},
			wantErr: false,
		},
		{
			name: "Negative: player not found",
			args: args{
				id: "5",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPlayerById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPlayerById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPlayerById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdatePlayer(t *testing.T) {
	type args struct {
		player Player
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdatePlayer(tt.args.player); (err != nil) != tt.wantErr {
				t.Errorf("UpdatePlayer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
