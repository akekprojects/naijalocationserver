-- +goose Up 
-- Enable the uuid-ossp extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE lgas(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT  NOT NULL ,
    state_id UUID  NOT NULL REFERENCES states(id) ON DELETE CASCADE
    );



-- +goose Down
DROP TABLE lgas;
