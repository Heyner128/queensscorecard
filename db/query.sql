-- name: GetScores :many
SELECT
    *
FROM
    scores;

-- name: GetNames :many
SELECT DISTINCT
    name
FROM
    scores;

-- name: GetScoresByNames :many
SELECT
    *
FROM
    scores
WHERE
    name = ?;

-- name: GetFastestTimeByMonth :many
SELECT
    name,
    gameNumber,
    min(secondsToSolve) as minSecondsToSolve,
    strftime('%Y-%m', datetime(timestamp, 'unixepoch')) as date
FROM
    scores
GROUP BY name, date
ORDER BY date, minSecondsToSolve;

-- name: GetFastestPlayersByMonth :many
SELECT
    name,
    group_concat(strftime('%Y-%m', datetime(timestamp, 'unixepoch')),', ') as week,
    COUNT(*) as fastest_count
FROM
    scores
WHERE (name, timestamp) in (
    SELECT
        name,
    timestamp
FROM
    scores
GROUP BY strftime('%Y-%m', datetime(timestamp, 'unixepoch'))
HAVING secondsToSolve = min(secondsToSolve)
    )
GROUP BY name
ORDER BY fastest_count DESC, name;

-- name: GetFastestPlayersByWeek :many
SELECT
    name,
    group_concat(strftime('%Y-%m', datetime(timestamp, 'unixepoch')),', ') as week,
    COUNT(*) as fastest_count
FROM
    scores
WHERE (name, timestamp) in (
    SELECT
        name,
    timestamp
FROM
    scores
GROUP BY strftime('%Y-%m', datetime(timestamp, 'unixepoch'))
HAVING secondsToSolve = min(secondsToSolve)
    )
GROUP BY name
ORDER BY fastest_count DESC, name;

-- name: CreateScore :exec
INSERT OR IGNORE INTO scores (
    name,
    gameNumber,
    secondsToSolve,
    timestamp
)
VALUES (
 ?,
 ?,
 ?,
 ?
);