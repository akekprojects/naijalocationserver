-- +goose Up
UPDATE states SET name = 'FCT' WHERE name = 'fct';