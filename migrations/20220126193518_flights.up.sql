-- type Flight struct {
-- 	Number      string
-- 	Origin      string
-- 	Destination string
-- 	Date        time.Time
-- 	Departure   time.Time
-- 	Arrival     time.Time
-- }

CREATE TABLE flights
(
    flight_id   INT GENERATED ALWAYS AS IDENTITY,
    number      varchar(10) NOT NULL,
    origin      VARCHAR(50) NOT NULL,
    destination varchar(50) NOT NULL,
    departure   timestamptz NOT NULL,
    arrival     timestamptz NOT NULL,
    UNIQUE (number, departure)
);