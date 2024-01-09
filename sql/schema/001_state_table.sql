-- +goose Up 
-- Enable the uuid-ossp extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE states(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT UNIQUE NOT NULL ,
    capital TEXT UNIQUE NOT NULL 
);


-- Inserting a custom row
INSERT INTO states (name, capital ) VALUES
( 'Abia ', 'Umuahia');
INSERT INTO states (name, capital ) VALUES
( 'Adamawa', 'Yola');
INSERT INTO states (name, capital ) VALUES
( 'Akwa Ibom', 'Uyo');
INSERT INTO states (name, capital ) VALUES
( 'Anambra', 'Awka');
INSERT INTO states (name, capital ) VALUES
( 'Bauchi', 'Bauchi');
INSERT INTO states (name, capital ) VALUES
( 'Bayelsa', 'Yenagoa');
INSERT INTO states (name, capital ) VALUES
( 'Benue', 'Makurdi');
INSERT INTO states (name, capital ) VALUES
( 'Borno', 'Maiduguri');
INSERT INTO states (name, capital ) VALUES
( 'Cross River', 'Calabar');
INSERT INTO states (name, capital ) VALUES
( 'Delta', 'Asaba');
INSERT INTO states (name, capital ) VALUES
( 'Ebonyi', 'Abakaliki');
INSERT INTO states (name, capital ) VALUES
( 'Edo', 'Benin City');
INSERT INTO states (name, capital ) VALUES
( 'Ekiti', 'Ado Ekiti');
INSERT INTO states (name, capital ) VALUES
( 'Enugu', 'Enugu');
INSERT INTO states (name, capital ) VALUES
( 'Gombe', 'Gombe');
INSERT INTO states (name, capital ) VALUES
( 'Imo', 'Owerri');
INSERT INTO states (name, capital ) VALUES
( 'Jigawa', 'Dutse');
INSERT INTO states (name, capital ) VALUES
( 'Kaduna', 'Kaduna');
INSERT INTO states (name, capital ) VALUES
( 'Kano', 'Kano');
INSERT INTO states (name, capital ) VALUES
( 'Kastina', 'Kastina');
INSERT INTO states (name, capital ) VALUES
( 'Kebbi', 'Birnin Kebbi');
INSERT INTO states (name, capital ) VALUES
( 'Kogi', 'Lokoja');
INSERT INTO states (name, capital ) VALUES
( 'Kwara', 'Ilorin');
INSERT INTO states (name, capital ) VALUES
( 'Lagos', 'Ikeja');
INSERT INTO states (name, capital ) VALUES
( 'Nassarawa', 'Lafia');
INSERT INTO states (name, capital ) VALUES
( 'Niger', 'Minna');
INSERT INTO states (name, capital ) VALUES
( 'Ogun', 'Abeokuta');
INSERT INTO states (name, capital ) VALUES
( 'Ondo', 'Akure');
INSERT INTO states (name, capital ) VALUES
( 'Osun', 'Oshogbo');
INSERT INTO states (name, capital ) VALUES
( 'Oyo', 'Ibadan');
INSERT INTO states (name, capital ) VALUES
( 'Plateau', 'Jos');
INSERT INTO states (name, capital ) VALUES
( 'Rivers', 'Port Harcourt');
INSERT INTO states (name, capital ) VALUES
( 'sokoto', 'sokoto');
INSERT INTO states (name, capital ) VALUES
( 'Taraba', 'Jalingo');
INSERT INTO states (name, capital ) VALUES
( 'Yobe', 'Damaturu');
INSERT INTO states (name, capital ) VALUES
( 'Zamfara', 'Gusau');
INSERT INTO states (name, capital ) VALUES
( 'FCT', 'Abuja');
-- +goose Down
DROP TABLE states;
