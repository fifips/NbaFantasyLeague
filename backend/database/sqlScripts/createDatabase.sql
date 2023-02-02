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
                          FOREIGN KEY (away_team_id) REFERENCES Team(id) ON DELETE CASCADE ,
                          FOREIGN KEY (home_team_id) REFERENCES Team(id) ON DELETE CASCADE
);

CREATE TABLE Player (
                        id CHAR(36),
                        PRIMARY KEY (id)
);

CREATE TABLE player_to_team (
                                player_id CHAR(36) UNIQUE,
                                team_id CHAR(36),

                                FOREIGN KEY (player_id) REFERENCES Player(id) ON DELETE CASCADE,
                                FOREIGN KEY (team_id) REFERENCES Team(id) ON DELETE CASCADE
);

CREATE TABLE player_performance (
                                    id INT AUTO_INCREMENT,
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
                                    PRIMARY KEY (id),
                                    FOREIGN KEY (game_id) REFERENCES Schedule(game_id) ON DELETE SET NULL,
                                    FOREIGN KEY (player_id) REFERENCES Player(id) ON DELETE SET NULL
);

CREATE TABLE User (
                      id INT AUTO_INCREMENT,
                      email VARCHAR(255) NOT NULL UNIQUE,
                      password LONGTEXT NOT NULL,
                      is_active TINYINT(1) DEFAULT 0,

                      PRIMARY KEY (id)
);

CREATE TABLE Session (
                         id VARCHAR(36) NOT NULL,
                         user_id INT,

                         PRIMARY KEY (user_id),
                         FOREIGN KEY (user_id) REFERENCES User(id) ON DELETE CASCADE
);

CREATE TABLE Activation_Code (
                                 code VARCHAR(36) NOT NULL,
                                 user_id INT NOT NULL UNIQUE,
                                 expires datetime NOT NULL,

                                 FOREIGN KEY (user_id) REFERENCES User(id) ON DELETE CASCADE,
                                 PRIMARY KEY (code)
);

CREATE TABLE League (
                        id INT AUTO_INCREMENT,
                        owner_id INT NOT NULL,
                        name VARCHAR(30) NOT NULL,
                        pts_ratio FLOAT DEFAULT 1.0,
                        reb_ratio FLOAT DEFAULT 1.0,
                        ast_ratio FLOAT DEFAULT 1.0,

                        FOREIGN KEY (owner_id) REFERENCES User(id) ON DELETE CASCADE,
                        PRIMARY KEY (id)
);

CREATE TABLE user_to_league (
                                user_id INT NOT NULL,
                                league_id INT NOT NULL,

                                FOREIGN KEY (user_id) REFERENCES User(id) ON DELETE CASCADE,
                                FOREIGN KEY (league_id) REFERENCES League(id) ON DELETE CASCADE
);
