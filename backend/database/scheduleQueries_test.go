package database

import (
	"testing"
	"time"
)

func TestGetSchedule(t *testing.T) {
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
			got, err := GetSchedule()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSchedule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !((len(got) == 0) == tt.isEmpty) {
				t.Errorf("GetSchedule() returned %v elements, isEmpty: %v", len(got), tt.isEmpty)
			}
		})
	}
}

func TestDeleteMatch(t *testing.T) {
	type args struct {
		game Match
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Positive",
			args: args{
				game: Match{
					GameId:     "3",
					GameDate:   time.Time{},
					HomeTeamId: "",
					HomeScore:  nil,
					AwayTeamId: "",
					AwayScore:  nil,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteMatch(tt.args.game); (err != nil) != tt.wantErr {
				t.Errorf("DeleteMatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateOrUpdateScheduledGame(t *testing.T) {
	type args struct {
		game Match
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Positive: Create",
			args: args{game: Match{
				GameId:     "4",
				GameDate:   time.Now(),
				HomeTeamId: "1",
				HomeScore:  nil,
				AwayTeamId: "2",
				AwayScore:  nil,
			}},
			wantErr: false,
		},
		{
			name: "Positive: Update",
			args: args{
				game: Match{
					GameId:     "1",
					GameDate:   time.Now(),
					HomeTeamId: "2",
					HomeScore:  nil,
					AwayTeamId: "1",
					AwayScore:  nil,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateOrUpdateScheduledGame(tt.args.game); (err != nil) != tt.wantErr {
				t.Errorf("CreateOrUpdateScheduledGame() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
