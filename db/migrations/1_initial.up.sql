CREATE TABLE users (
 user_id       SERIAL       PRIMARY KEY
,email         VARCHAR(200) NOT NULL UNIQUE
,token         VARCHAR(128)
,secret        VARCHAR(128)
,expiration    TIMESTAMP
,password_salt VARCHAR(128)
,password_hash VARCHAR(128)
,created_on    TIMESTAMP    DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE roles (
 role_id   SERIAL PRIMARY KEY
,role_name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE user_roles (
 user_role_id SERIAL  PRIMARY KEY
,user_id      INTEGER REFERENCES users
,role_id      INTEGER REFERENCES roles
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
,plan_name VARCHAR(35) NOT NULL UNIQUE
,price     FLOAT       NOT NULL
);

CREATE TABLE members (
 member_id  SERIAL       PRIMARY KEY
,user_id    INTEGER      REFERENCES users
,address_id INTEGER      REFERENCES addresses
,plan_id    INTEGER      NOT NULL REFERENCES plans
,active     BOOLEAN
,first_name VARCHAR(35)  NOT NULL
,last_name  VARCHAR(35)  NOT NULL
);

CREATE TABLE devices (
 device_id    SERIAL      PRIMARY KEY
,user_id      INTEGER     NOT NULL REFERENCES users
,device_token VARCHAR(64) NOT NULL
);

CREATE TABLE gyms (
 gym_id   SERIAL       PRIMARY KEY
,user_id  INTEGER      REFERENCES users
,gym_name VARCHAR(50)  NOT NULL
);

CREATE TABLE images (
 image_id    SERIAL       PRIMARY KEY
,gym_id      INTEGER      REFERENCES gyms
,user_id     INTEGER      REFERENCES users
,image_path  VARCHAR(100)
);

CREATE TABLE hours (
 hour_id    SERIAL  PRIMARY KEY
,gym_id     INTEGER REFERENCES gyms
,day        INTEGER NOT NULL
,open_time  TIME    NOT NULL
,close_time TIME    NOT NULL
);

CREATE TABLE holidays (
 holiday_id   SERIAL      PRIMARY KEY
,holiday_name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE gym_holidays (
 gym_holiday_id SERIAL  PRIMARY KEY
,gym_id         INTEGER NOT NULL REFERENCES gyms
,holiday_id     INTEGER NOT NULL REFERENCES holidays
,hour_id        INTEGER NOT NULL REFERENCES hours
);

CREATE TABLE features (
 feature_id          SERIAL       PRIMARY KEY
,feature_name        VARCHAR(100) NOT NULL UNIQUE
,feature_description TEXT         NOT NULL UNIQUE
);

CREATE TABLE gym_features (
 gym_feature_id SERIAL  PRIMARY KEY
,gym_id         INTEGER REFERENCES gyms
,feature_id     INTEGER REFERENCES features
);

CREATE TABLE locations (
 location_id   SERIAL       PRIMARY KEY
,address_id    INTEGER      REFERENCES addresses
,location_name VARCHAR(50)  NOT NULL
,phone_number  VARCHAR(15)
,website_url   VARCHAR(100)
,in_network    BOOLEAN
);

CREATE TABLE statuses (
 status_id   SERIAL      PRIMARY KEY
,status_name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE visits (
 visit_id    SERIAL    PRIMARY KEY
,member_id   INTEGER   REFERENCES members
,location_id INTEGER   REFERENCES locations
,created_on  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
,status_id   INTEGER   REFERENCES statuses
);

-- Table to hold required info for outside memberships.
-- Additional necessary fields will be added later.
CREATE TABLE outside_memberships (
 outside_membership_id SERIAL  PRIMARY KEY
,member_id             INTEGER NOT NULL REFERENCES members
,location_id           INTEGER NOT NULL REFERENCES locations
,gym_id                INTEGER NOT NULL REFERENCES gyms
);

CREATE TABLE support_sources (
 support_source_id   SERIAL      PRIMARY KEY
,support_source_name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE support_requests (
 support_request_id SERIAL  PRIMARY KEY
,user_id            INTEGER REFERENCES users
,support_source_id  INTEGER REFERENCES support_sources
,content            TEXT    NOT NULL
);
