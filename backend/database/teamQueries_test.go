package database

import (
	"reflect"
	"testing"
)

func TestDeleteTeam(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Positive",
			args: args{
				id: "2",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteTeam(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteTeam() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetAllTeamIds(t *testing.T) {
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
			got, err := GetAllTeamIds()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllTeamIds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !((len(got) == 0) == tt.isEmpty) {
				t.Errorf("GetAllTeamIds() returned %v elements, isEmpty = %v", len(got), tt.isEmpty)
			}
		})
	}
}

func TestGetTeamById(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    Team
		wantErr bool
	}{
		{
			name: "Positive",
			args: args{
				id: "1",
			},
			want: Team{
				Id:        "1",
				FullName:  "Golden State Warriors",
				Acronym:   "GSW",
				Wins:      72,
				Losses:    10,
				PlayerIds: nil,
			},
			wantErr: false,
		},
		{
			name: "Negative",
			args: args{
				id: "5",
			},
			want:    Team{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTeamById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTeamById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTeamById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateOrUpdateTeam(t *testing.T) {
	type args struct {
		team Team
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Positive: Create",
			args: args{
				team: Team{
					Id:        "4",
					FullName:  "Oklahoma City Thunder",
					Acronym:   "OKC",
					Wins:      5,
					Losses:    6,
					PlayerIds: nil,
				},
			},
			wantErr: false,
		},
		{
			name: "Positive: Update",
			args: args{
				team: Team{
					Id:        "3",
					FullName:  "Boston Celtics",
					Acronym:   "BOS",
					Wins:      56,
					Losses:    27,
					PlayerIds: nil,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateOrUpdateTeam(tt.args.team); (err != nil) != tt.wantErr {
				t.Errorf("CreateOrUpdateTeam() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
