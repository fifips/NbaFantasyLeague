USE test_nba_fantasy_league;
INSERT INTO Team (id, full_name, acronym, wins, losses)
VALUES ('1', 'Golden State Warriors', 'GSW', 72, 10),
       ('2', 'Houston Rockets', 'HOU', 65, 17),
       ('3', 'Boston Celtics', 'BOS', 55, 27);

INSERT INTO Schedule (game_id, game_date, home_team_id, home_score, away_team_id, away_score)
VALUES ('1', '2022-01-01 12:00:00', '1', 120, '2', 110),
       ('2', '2022-01-02 12:00:00', '2', 115, '3', 110),
       ('3', '2022-01-03 12:00:00', '1', 110, '3', 105);

INSERT INTO Player (id)
VALUES ('1'), ('2'), ('3');

INSERT INTO player_to_team (player_id, team_id)
VALUES ('1', '1'), ('2', '1'), ('3', '2');

INSERT INTO player_performance (game_id, player_id, minutes, points, assists, rebounds, turnovers, steals, blocks, fouls)
VALUES ('1', '1', '01:20:00', 30, 10, 5, 2, 2, 2, 2),
       ('1', '2', '01:30:00', 25, 5, 8, 3, 1, 0, 3),
       ('2', '3', '01:40:00', 20, 8, 6, 2, 1, 2, 2);

INSERT INTO User (email, password, is_active)
VALUES ('user1@example.com', 'password1', 0),
       ('user2@example.com', 'password2', 1),
       ('user3@example.com', 'password3', 1),
       ('user4@example.com', 'password4', 1);

INSERT INTO Session (id, user_id)
VALUES ('1ba7b810-9dad-11d1-80b4-00c04fd430c8', '1'),
       ('2ba7b810-9dad-11d1-80b4-00c04fd430c8', '2'),
       ('3ba7b810-9dad-11d1-80b4-00c04fd430c8', '3');

INSERT INTO Activation_Code (code, user_id, expires)
VALUES ('1ba7b810-9dad-11d1-80b4-00c04fd430c8', '1', '2022-01-01 12:00:00'),
       ('2ba7b810-9dad-11d1-80b4-00c04fd430c8', '2', '2022-01-02 12:00:00'),
       ('3ba7b810-9dad-11d1-80b4-00c04fd430c8', '3', '2022-01-03 12:00:00');

INSERT INTO league (id, owner_id, name, pts_ratio, reb_ratio, ast_ratio)
VALUES ('1', '1', 'first_league', 1.1, 1.2, 1.3),
       ('2', '2', 'second_league', 1.1, 1.2, 1.3),
       ('3', '3', 'third_league', 1.1, 1.2, 1.3),
       ('4', '4', 'fourth_league', 1.1, 1.2, 1.3);

INSERT INTO user_to_league (user_id, league_id)
VALUES ('1', '1'),
       ('2', '2'),
       ('3', '3'),
       ('4', '4'),
       ('3', '2'),
       ('2', '1');
