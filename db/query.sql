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
WITH monthly_fastest AS (
    SELECT
        name,
        DATE_FORMAT(FROM_UNIXTIME(timestamp), '%Y-%m') AS month,
    MIN(secondsToSolve) AS min_seconds
FROM
    scores
GROUP BY DATE_FORMAT(FROM_UNIXTIME(timestamp), '%Y-%m')
    )
SELECT
    name,
    GROUP_CONCAT(month ORDER BY month) AS months,
    COUNT(*) AS fastest_count
FROM
    monthly_fastest
GROUP BY name
ORDER BY fastest_count, name;


-- name: GetFastestPlayersByWeek :many
WITH weekly_fastest AS (
    SELECT
        name,
        DATE_FORMAT(FROM_UNIXTIME(timestamp), '%Y-%u') AS week,
    MIN(secondsToSolve) AS min_seconds
FROM
    scores
GROUP BY DATE_FORMAT(FROM_UNIXTIME(timestamp), '%Y-%u')
    )
SELECT
    name,
    GROUP_CONCAT(week ORDER BY week) AS week,
    COUNT(*) AS fastest_count
FROM
    weekly_fastest
GROUP BY name
ORDER BY fastest_count, name;



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