CREATE TABLE IF NOT EXISTS scores (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    gameNumber INT NOT NULL,
    secondsToSolve INT NOT NULL,
    timestamp INT NOT NULL,
    UNIQUE KEY unique_name_game_timestamp (name, gameNumber, timestamp)
);