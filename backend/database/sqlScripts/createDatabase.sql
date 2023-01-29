CREATE DATABASE IF NOT EXISTS nba_fantasy_league;
USE nba_fantasy_league;

CREATE TABLE Team (
                      id CHAR(36),
                      full_name CHAR(36),
                      acronym CHAR(3),
                      wins INT,
                      losses INT,

                      PRIMARY KEY (id)
);

CREATE TABLE Schedule (
    game_id CHAR(36),
    game_date DATETIME,
    home_team_id CHAR(36),
    home_score INT NULL,
    away_team_id CHAR(36),
    away_score INT NULL,
    PRIMARY KEY (game_id),
    FOREIGN KEY (away_team_id) REFERENCES Team(id),
    FOREIGN KEY (home_team_id) REFERENCES Team(id)
);

CREATE TABLE Player (
                      id CHAR(36),
                      PRIMARY KEY (id)
);

CREATE TABLE player_to_team (
    player_id CHAR(36),
    team_id CHAR(36),
    FOREIGN KEY (player_id) REFERENCES Player(id),
    FOREIGN KEY (team_id) REFERENCES Team(id)
);

CREATE TABLE player_performance (
    game_id CHAR(36),
    player_id CHAR(36),
    minutes TIME,
    points int,
    assists int,
    rebounds int,
    turnovers int,
    steals int,
    blocks int,
    fouls int,
    PRIMARY KEY (game_id, player_id),
    FOREIGN KEY (game_id) REFERENCES Schedule(game_id),
    FOREIGN KEY (player_id) REFERENCES Player(id)
);

CREATE TABLE User (
                      id INT AUTO_INCREMENT,
                      email VARCHAR(255) NOT NULL UNIQUE,
                      password LONGTEXT NOT NULL,

                      PRIMARY KEY (id)
);

CREATE TABLE Session (
                         id VARCHAR(36) NOT NULL,
                         user_id INT NOT NULL UNIQUE,
                         FOREIGN KEY (user_id) REFERENCES User(id) ON DELETE CASCADE,
                         PRIMARY KEY (id)
);