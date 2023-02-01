package database

import (
	"reflect"
	"testing"
)

func TestCreateLeague(t *testing.T) {
	type args struct {
		league League
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Positive",
			args: args{
				league: League{
					OwnerId:      1,
					Participants: nil,
					Name:         "new_first_league",
					PtsRatio:     1.0,
					RebRatio:     2.0,
					AstRatio:     3.0,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateLeague(tt.args.league); (err != nil) != tt.wantErr {
				t.Errorf("CreateLeague() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteLeagueById(t *testing.T) {
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
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteLeagueById(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteLeagueById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetLeagueById(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    League
		wantErr bool
	}{
		{
			name: "Positive",
			args: args{
				id: 1,
			},
			want: League{
				Id:           1,
				OwnerId:      1,
				Participants: nil,
				Name:         "first_league",
				PtsRatio:     1.1,
				RebRatio:     1.2,
				AstRatio:     1.3,
			},
			wantErr: false,
		},
		{
			name: "Negative",
			args: args{
				id: 7,
			},
			want:    League{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetLeagueById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLeagueById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLeagueById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllLeagueIds(t *testing.T) {
	tests := []struct {
		name    string
		isEmpty bool
		wantErr bool
	}{
		{
			name:    "Positive",
			isEmpty: false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAllLeagueIds()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllLeagueIds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !((len(got) == 0) == tt.isEmpty) {
				t.Errorf("GetAllLeagueIds() got %v elements, isEmpty = %v", len(got), tt.isEmpty)
			}
		})
	}
}
