CREATE TABLE prize_placement (
  id CHAR(36) PRIMARY KEY,
  playerID CHAR(36),
  tournamentID CHAR(36),
  placement INT NOT NULL,
  FOREIGN KEY (playerID) REFERENCES player(id) ON DELETE CASCADE,
  FOREIGN KEY (tournamentID) REFERENCES tournament(id) ON DELETE CASCADE
);