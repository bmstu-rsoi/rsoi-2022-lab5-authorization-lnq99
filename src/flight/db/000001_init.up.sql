CREATE TABLE airport
(
    id      SERIAL PRIMARY KEY,
    name    VARCHAR(255) NOT NULL,
    city    VARCHAR(255) NOT NULL,
    country VARCHAR(255) NOT NULL
);

CREATE TABLE flight
(
    id              SERIAL PRIMARY KEY,
    flight_number   VARCHAR(20)              NOT NULL,
    datetime        TIMESTAMP WITH TIME ZONE NOT NULL,
    from_airport_id INT REFERENCES airport (id) NOT NULL,
    to_airport_id   INT REFERENCES airport (id) NOT NULL,
    price           INT                      NOT NULL
);