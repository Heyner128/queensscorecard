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
    MIN(secondsToSolve) AS minSecondsToSolve,
    DATE_FORMAT(FROM_UNIXTIME(timestamp), '%Y-%m') AS date
FROM
    scores
GROUP BY name, date
ORDER BY date, minSecondsToSolve;


-- name: GetFastestPlayersByMonth :many
SELECT
    name,
    GROUP_CONCAT(DATE_FORMAT(FROM_UNIXTIME(timestamp), '%Y-%m'), ', ') AS month,
    COUNT(*) AS fastest_count
FROM
    scores
WHERE (name, timestamp) IN (
    SELECT
    name,
    timestamp
    FROM
    scores
    GROUP BY DATE_FORMAT(FROM_UNIXTIME(timestamp), '%Y-%m')
    HAVING secondsToSolve = MIN(secondsToSolve)
    )
GROUP BY name
ORDER BY fastest_count DESC, name;

-- name: GetFastestPlayersByWeek :many
SELECT
    name,
    GROUP_CONCAT(DATE_FORMAT(FROM_UNIXTIME(timestamp), '%Y-%u'), ', ') AS week,
    COUNT(*) AS fastest_count
FROM
    scores
WHERE (name, timestamp) IN (
    SELECT
        name,
    timestamp
FROM
    scores
GROUP BY DATE_FORMAT(FROM_UNIXTIME(timestamp), '%Y-%u')
HAVING secondsToSolve = MIN(secondsToSolve)
    )
GROUP BY name
ORDER BY fastest_count DESC, name;


-- name: CreateScore :exec
INSERT IGNORE INTO scores (
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