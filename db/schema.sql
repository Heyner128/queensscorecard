CREATE TABLE IF NOT EXISTS scores (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    gameNumber INTEGER NOT NULL,
    secondsToSolve INTEGER NOT NULL,
    timestamp INTEGER NOT NULL,
    UNIQUE(name, gameNumber, timestamp)
);

