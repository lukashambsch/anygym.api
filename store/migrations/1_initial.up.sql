CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE users (
 user_id       SERIAL       PRIMARY KEY
,email         VARCHAR(200) NOT NULL UNIQUE
,token         VARCHAR(128) NOT NULL DEFAULT ''
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

CREATE TABLE gyms (
 gym_id             SERIAL       PRIMARY KEY
,user_id            INTEGER      REFERENCES users
,gym_name           VARCHAR(50)  NOT NULL
,monthly_member_fee FLOAT
);

CREATE TABLE gym_locations (
 gym_location_id    SERIAL       PRIMARY KEY
,gym_id             INTEGER      NOT NULL REFERENCES gyms
,address_id         INTEGER      NOT NULL UNIQUE REFERENCES addresses
,location_name      VARCHAR(50)  NOT NULL
,phone_number       VARCHAR(15)  NOT NULL DEFAULT ''
,website_url        VARCHAR(255) NOT NULL DEFAULT ''
,in_network         BOOLEAN      NOT NULL DEFAULT False
,monthly_member_fee FLOAT
);

CREATE TABLE images (
 image_id        SERIAL       PRIMARY KEY
,gym_id          INTEGER      REFERENCES gyms ON DELETE CASCADE
,gym_location_id INTEGER      REFERENCES gym_locations ON DELETE CASCADE
,user_id         INTEGER      UNIQUE REFERENCES users ON DELETE CASCADE
,image_path      VARCHAR(255)
,UNIQUE(gym_location_id, image_path)
,UNIQUE(gym_id, image_path)
);

CREATE TABLE members (
 member_id  SERIAL       PRIMARY KEY
,user_id    INTEGER      NOT NULL UNIQUE REFERENCES users ON DELETE CASCADE
,image_id   INTEGER      UNIQUE REFERENCES images
,address_id INTEGER      UNIQUE REFERENCES addresses
,first_name VARCHAR(35)  NOT NULL
,last_name  VARCHAR(35)  NOT NULL
,CONSTRAINT valid_name CHECK(first_name <> '' OR last_name <> '')
);

CREATE TABLE memberships (
 membership_id SERIAL    PRIMARY KEY
,plan_id       INTEGER   NOT NULL REFERENCES plans ON DELETE CASCADE
,member_id     INTEGER   NOT NULL REFERENCES members ON DELETE CASCADE
,start_date    TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
,renew_date    TIMESTAMP
,end_date      TIMESTAMP
,active        BOOLEAN
);

CREATE TABLE devices (
 device_id    SERIAL      PRIMARY KEY
,user_id      INTEGER     NOT NULL REFERENCES users ON DELETE CASCADE
,device_token VARCHAR(64) NOT NULL
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
,gym_id         INTEGER NOT NULL REFERENCES gyms ON DELETE CASCADE
,feature_id     INTEGER NOT NULL REFERENCES features ON DELETE CASCADE
,UNIQUE(gym_id, feature_id)
);

CREATE TABLE days (
 day_id   SERIAL     PRIMARY KEY
,day_name VARCHAR(9) NOT NULL UNIQUE
);

CREATE TABLE business_hours (
 business_hour_id    SERIAL  PRIMARY KEY
,gym_location_id     INTEGER NOT NULL REFERENCES gym_locations ON DELETE CASCADE
,holiday_id          INTEGER REFERENCES holidays ON DELETE CASCADE
,day_id              INTEGER REFERENCES days ON DELETE CASCADE
,open_time           TIME    WITH TIME ZONE NOT NULL
,close_time          TIME    WITH TIME ZONE NOT NULL
,UNIQUE(gym_location_id, day_id)
,UNIQUE(gym_location_id, holiday_id)
,CONSTRAINT holiday_or_day CHECK((holiday_id IS NOT NULL or day_id IS NOT NULL) AND (holiday_id IS NULL or day_id IS NULL))
);

CREATE TABLE statuses (
 status_id   SERIAL      PRIMARY KEY
,status_name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE visits (
 visit_id        SERIAL    PRIMARY KEY
,member_id       INTEGER   NOT NULL REFERENCES members
,gym_location_id INTEGER   NOT NULL REFERENCES gym_locations
,status_id       INTEGER   NOT NULL REFERENCES statuses
,created_on      TIMESTAMP DEFAULT CURRENT_TIMESTAMP
,modified_on     TIMESTAMP
);

-- Table to hold required info for outside memberships.
-- Additional necessary fields will be added later.
CREATE TABLE outside_memberships (
 outside_membership_id SERIAL  PRIMARY KEY
,member_id             INTEGER NOT NULL REFERENCES members ON DELETE CASCADE
,gym_location_id       INTEGER REFERENCES gym_locations ON DELETE CASCADE
,gym_id                INTEGER REFERENCES gyms ON DELETE CASCADE
,CONSTRAINT gym_location_or_gym CHECK(
  (gym_location_id IS NOT NULL OR gym_id IS NOT NULL) AND (gym_location_id IS NULL OR gym_id IS NULL)
)
);

CREATE TABLE support_sources (
 support_source_id   SERIAL      PRIMARY KEY
,support_source_name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE support_requests (
 support_request_id SERIAL    PRIMARY KEY
,user_id            INTEGER   REFERENCES users ON DELETE CASCADE
,support_source_id  INTEGER   REFERENCES support_sources
,content            TEXT      NOT NULL
,notes              TEXT
,created_on         TIMESTAMP DEFAULT CURRENT_TIMESTAMP
,resolved_on        TIMESTAMP
,CONSTRAINT has_content CHECK(content <> '')
);
