-- INSERT INTO airport(name, city, country)
-- VALUES
--     ('Airport 1', 'City 1', 'Country 1'),
--     ('Airport 2', 'City 2', 'Country 2'),
--     ('Airport 3', 'City 3', 'Country 3'),
--     ('Airport 4', 'City 4', 'Country 4')
-- ON CONFLICT
-- DO NOTHING;
--
-- INSERT INTO flight(flight_number, datetime, from_airport_id, to_airport_id, price)
-- VALUES
--     ('Flight 1', '2021-10-08T19:59:19Z', 1, 2, 100),
--     ('Flight 2', '2021-10-09T19:59:19Z', 1, 3, 200),
--     ('Flight 3', '2021-10-10T19:59:19Z', 2, 1, 300),
--     ('Flight 4', '2021-10-11T19:59:19Z', 2, 4, 400),
--     ('Flight 5', '2021-10-12T19:59:19Z', 3, 2, 500),
--     ('Flight 6', '2021-10-13T19:59:19Z', 3, 4, 600)
-- ON CONFLICT
-- DO NOTHING;

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