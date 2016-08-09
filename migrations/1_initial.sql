-- +migrate Up

CREATE TABLE users (
 user_id       SERIAL       PRIMARY KEY
,email         VARCHAR(200) NOT NULL
,password_salt VARCHAR(255) NOT NULL
,password_hash BYTEA        NOT NULL
,created_on    TIMESTAMP    DEFAULT CURRENT_TIMESTAMP

,UNIQUE(email)
);

CREATE TABLE addresses (
 address_id     SERIAL       PRIMARY KEY
,country        VARCHAR(5)
,state_region   VARCHAR(2)
,city           VARCHAR(35)
,postal_area    VARCHAR(10)
,street_address VARCHAR(100)
,latitude       FLOAT
,longitude      FLOAT
);

CREATE TABLE plans (
 plan_id   SERIAL      PRIMARY KEY
,plan_name VARCHAR(35) NOT NULL
,price     FLOAT       NOT NULL
);

CREATE TABLE members (
 member_id  SERIAL       PRIMARY KEY,
,address_id INTEGER      REFERENCES addresses,
,plan_id    INTEGER      REFERENCES plans,
,first_name VARCHAR(35)  NOT NULL,
,last_name  VARCHAR(35)  NOT NULL,
,img_path   VARCHAR(100) NOT NULL
);

CREATE TABLE gyms (
 gym_id   SERIAL      PRIMARY KEY
,gym_name VARCHAR(50) NOT NULL
);

CREATE TABLE locations (
 location_id         SERIAL PRIMARY KEY,
,address_id          INTEGER REFERENCES addresses,
,locationsation_name VARCHAR(50) NOT NULL
);

CREATE TABLE visits (
 visit_id SERIAL     PRIMARY KEY
,member_id INTEGER   REFERENCES members
,location_id INTEGER REFERENCES locations
);

-- +migrate Down

DROP TABLE users;
DROP TABLE addresses;
DROP TABLE plans;
DROP TABLE members;
DROP TABLE gyms;
DROP TABLE locations;
DROP TABLE visits;
