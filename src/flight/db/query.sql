-- name: GetAirport :one
SELECT * FROM airport
WHERE id=$1 LIMIT 1;

-- name: ListAirports :many
SELECT * FROM airport
ORDER BY name;

-- name: GetFlight :one
SELECT * FROM flight
WHERE flight_number=$1 LIMIT 1;

-- name: ListFlights :many
SELECT * FROM flight;

-- name: ListFlightsWithOffsetLimit :many
SELECT * FROM flight
LIMIT $1
OFFSET $2;
