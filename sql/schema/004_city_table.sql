-- +goose Up
-- Enable the uuid-ossp extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE city (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL UNIQUE,
    population TEXT NOT NULL,
    state_id UUID REFERENCES states(id) NOT NULL
    
);

-- +goose Down
DROP TABLE city;