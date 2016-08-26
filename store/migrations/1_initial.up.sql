CREATE TABLE users (
 user_id       SERIAL       PRIMARY KEY
,email         VARCHAR(200) NOT NULL UNIQUE
,token         VARCHAR(128) NOT NULL DEFAULT ''
,secret        VARCHAR(128) NOT NULL DEFAULT ''
,password_salt VARCHAR(128) NOT NULL DEFAULT ''
,password_hash VARCHAR(128) NOT NULL DEFAULT ''
,created_on    TIMESTAMP    DEFAULT CURRENT_TIMESTAMP
,CONSTRAINT valid_name CHECK(email <> '')
);

CREATE TABLE roles (
 role_id   SERIAL PRIMARY KEY
,role_name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE user_roles (
 user_role_id SERIAL  PRIMARY KEY
,user_id      INTEGER NOT NULL REFERENCES users
,role_id      INTEGER NOT NULL REFERENCES roles
,UNIQUE(user_id, role_id)
);

CREATE TABLE addresses (
 address_id     SERIAL       PRIMARY KEY
,country        VARCHAR(10)
,state_region   VARCHAR(10)
,city           VARCHAR(35)
,postal_area    VARCHAR(10)
,street_address VARCHAR(100)
,latitude       FLOAT
,longitude      FLOAT

,UNIQUE(country, state_region, city, postal_area, street_address)
);

CREATE TABLE plans (
 plan_id   SERIAL      PRIMARY KEY
,plan_name VARCHAR(35) NOT NULL
,price     FLOAT       NOT NULL

,UNIQUE(plan_name, price)
);

CREATE TABLE members (
 member_id  SERIAL       PRIMARY KEY
,user_id    INTEGER      NOT NULL UNIQUE REFERENCES users
,address_id INTEGER      UNIQUE REFERENCES addresses
,first_name VARCHAR(35)  NOT NULL
,last_name  VARCHAR(35)  NOT NULL
,CONSTRAINT valid_name CHECK(first_name <> '' OR last_name <> '')
);

CREATE TABLE memberships (
 membership_id SERIAL    PRIMARY KEY
,plan_id       INTEGER   NOT NULL REFERENCES plans
,member_id     INTEGER   NOT NULL REFERENCES members
,start_date    TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
,renew_date    TIMESTAMP
,end_date      TIMESTAMP
,active        BOOLEAN
);

CREATE TABLE devices (
 device_id    SERIAL      PRIMARY KEY
,user_id      INTEGER     NOT NULL REFERENCES users
,device_token VARCHAR(64) NOT NULL
);

CREATE TABLE gyms (
 gym_id             SERIAL       PRIMARY KEY
,user_id            INTEGER      REFERENCES users
,gym_name           VARCHAR(50)  NOT NULL
,monthly_member_fee FLOAT
);

CREATE TABLE images (
 image_id    SERIAL       PRIMARY KEY
,gym_id      INTEGER      REFERENCES gyms
,user_id     INTEGER      UNIQUE REFERENCES users
,image_path  VARCHAR(255)
,UNIQUE(gym_id, image_path)
);

CREATE TABLE holidays (
 holiday_id   SERIAL      PRIMARY KEY
,holiday_name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE features (
 feature_id          SERIAL       PRIMARY KEY
,feature_name        VARCHAR(100) NOT NULL UNIQUE
,feature_description TEXT         NOT NULL UNIQUE
,CONSTRAINT name_or_description CHECK(feature_name <> '' OR feature_description <> '')
);

CREATE TABLE gym_features (
 gym_feature_id SERIAL  PRIMARY KEY
,gym_id         INTEGER NOT NULL REFERENCES gyms
,feature_id     INTEGER NOT NULL REFERENCES features
,UNIQUE(gym_id, feature_id)
);

CREATE TABLE locations (
 location_id        SERIAL       PRIMARY KEY
,address_id         INTEGER      REFERENCES addresses
,location_name      VARCHAR(50)  NOT NULL
,phone_number       VARCHAR(15)
,website_url        VARCHAR(255)
,in_network         BOOLEAN
,monthly_member_fee FLOAT
);

CREATE TABLE days (
 day_id   SERIAL PRIMARY KEY
,day_name VARCHAR(9) NOT NULL
);

CREATE TABLE business_hours (
 business_hour_id    SERIAL  PRIMARY KEY
,location_id         INTEGER NOT NULL REFERENCES locations
,holiday_id          INTEGER REFERENCES holidays
,day_id              INTEGER REFERENCES days
,open_time           TIME    NOT NULL
,close_time          TIME    NOT NULL
,UNIQUE(location_id, day_id)
,UNIQUE(location_id, holiday_id)
);

CREATE TABLE statuses (
 status_id   SERIAL      PRIMARY KEY
,status_name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE visits (
 visit_id    SERIAL    PRIMARY KEY
,member_id   INTEGER   NOT NULL REFERENCES members
,location_id INTEGER   NOT NULL REFERENCES locations
,status_id   INTEGER   NOT NULL REFERENCES statuses
,created_on  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
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
 support_request_id SERIAL    PRIMARY KEY
,user_id            INTEGER   REFERENCES users
,support_source_id  INTEGER   REFERENCES support_sources
,content            TEXT      NOT NULL
,created_on         TIMESTAMP DEFAULT CURRENT_TIMESTAMP
,resolved_on        TIMESTAMP
);
