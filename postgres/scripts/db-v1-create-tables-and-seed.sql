\c tickets;

CREATE TABLE ticket
(
    id            SERIAL PRIMARY KEY,
    ticket_uid    uuid UNIQUE NOT NULL,
    username      VARCHAR(80) NOT NULL,
    flight_number VARCHAR(20) NOT NULL,
    price         INT         NOT NULL,
    status        VARCHAR(20) NOT NULL
        CHECK (status IN ('PAID', 'CANCELED'))
);


\c flights;

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

INSERT INTO airport(name, city, country)
VALUES
    ('Шереметьево', 'Москва', 'Россия'),
    ('Пулково', 'Санкт-Петербург', 'Россия')
    ON CONFLICT
DO NOTHING;

INSERT INTO flight(flight_number, datetime, from_airport_id, to_airport_id, price)
VALUES
    ('AFL031', '2021-10-08 20:00', 2, 1, 1500)
    ON CONFLICT
DO NOTHING;


\c privileges;

CREATE TABLE privilege
(
    id       SERIAL PRIMARY KEY,
    username VARCHAR(80) NOT NULL UNIQUE,
    status   VARCHAR(80) NOT NULL DEFAULT 'BRONZE'
        CHECK (status IN ('BRONZE', 'SILVER', 'GOLD')),
    balance  INT NOT NULL
);

CREATE TABLE privilege_history
(
    id             SERIAL PRIMARY KEY,
    privilege_id   INT REFERENCES privilege (id) NOT NULL,
    ticket_uid     uuid        NOT NULL,
    datetime       TIMESTAMP   NOT NULL,
    balance_diff   INT         NOT NULL,
    operation_type VARCHAR(20) NOT NULL
        CHECK (operation_type IN ('FILL_IN_BALANCE', 'DEBIT_THE_ACCOUNT'))
);
