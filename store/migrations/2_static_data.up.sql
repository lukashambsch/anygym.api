INSERT INTO plans (plan_name, price)
VALUES
  ('All Access', 50),
  ('All Access', 0);

INSERT INTO statuses (status_name)
VALUES
  ('Pending'),
  ('Approved'),
  ('Denied - Identity'),
  ('Denied - Banned');

INSERT INTO holidays (holiday_name)
VALUES
  ('New Year''s Day'),
  ('Martin Luther King Day'),
  ('Presidents Day'),
  ('Memorial Day'),
  ('Independence Day'),
  ('Labor Day'),
  ('Columbus Day'),
  ('Veterans Day'),
  ('Thanksgiving'),
  ('Christmans Eve'),
  ('Christmas Day'),
  ('New Year''s Eve');

INSERT INTO features (feature_name, feature_description)
VALUES
  ('Towel Service', 'Towels are provided at the location at no additional charge.'),
  ('Sauna', 'A sauna is provided.'),
  ('Personal Viewing Screens', 'Cardio equipment is equipped with a small television.'),
  ('Childcare (Free)', 'Childcare is provided at no additional charge.'),
  ('Childcare (Extra Cost)', 'Childcare is provided with an additional charge. (Please call this location for details.)'),
  ('Wifi (Free)', 'Wifi service is provided at no additional charge.'),
  ('Wifi (Extra Cost)', 'Wifi service is provided with an additional charge. (Please call this location for deatils.)'),
  ('Lap Pool (Indoor)', 'An indoor lap pool is provided.'),
  ('Lap Pool (Outdoor)', 'An outdoor lap pool is provided.'),
  ('Group Exercise', 'Group Exercise classes are provided.'),
  ('Racquetball Court (Indoor)', 'Indoor racquetball courts are provided.'),
  ('Racquetball Court (Outdoor)', 'Outdoor racquetball courts are provided.'),
  ('Personal Training (Free)', 'Free personal training is provided.'),
  ('Personal Training (Extra Cost)', 'Personal traingin is provided at an additional cost. (Please call this location for details.)'),
  ('Basketball Court (Full Size)', 'Full size basketball courts are provided.'),
  ('Basektball Court (Partial)', 'A small full court or half court basketball courts are provided.'),
  ('Strength Machines', 'Strength training machines are provided.'),
  ('Free Weights (Full)', 'A full range of free weights are provided. Including dumbbells, barbells, and squat racks.'),
  ('Free Weights (Partial)', 'Some free weights are provided.'),
  ('Whirlpool', 'A whirlpool hot tub is provided.'),
  ('Cardio Equipment', 'Cardio Equipment is provided.'),
  ('Steam Room', 'A steam room is provided.'),
  ('Locker Rooms', 'Locker rooms are provided.'),
  ('Group Cycling (Spin)', 'Group cycling or spin classes are provided.');

INSERT INTO users (email, password_hash)
VALUES
  ('lukas.hambsch@gmail.com', crypt('testpass', gen_salt('bf', 10))),
  ('bugentry@hotmail.com', crypt('testpass', gen_salt('bf', 10)));

INSERT INTO roles (role_name)
VALUES
  ('admin'),
  ('employee'),
  ('gym'),
  ('location'),
  ('member');

INSERT INTO user_roles (user_id, role_id)
VALUES
  (
    (SELECT user_id FROM users WHERE email = 'lukas.hambsch@gmail.com'),
    (SELECT role_id FROM roles WHERE role_name = 'admin')
  ),
  (
    (SELECT user_id FROM users WHERE email = 'bugentry@hotmail.com'),
    (SELECT role_id FROM roles WHERE role_name = 'employee')
  );

INSERT INTO images (image_path)
VALUES
  ('lukas-hambsch-profile.jpg'),
  ('mckenzie-hambsch-profile.jpg');

INSERT INTO support_sources (support_source_name)
VALUES
  ('website'),
  ('mobile app - logged in'),
  ('mobile app - logged out'),
  ('web app - members'),
  ('web app - gyms'),
  ('email');

INSERT INTO members (user_id, first_name, last_name, image_id)
VALUES
  (
    (SELECT user_id FROM users WHERE email = 'lukas.hambsch@gmail.com'),
    'Lukas',
    'Hambsch',
    (SELECT image_id FROM images WHERE image_path = 'lukas-hambsch-profile.jpg')
  ),
  (
    (SELECT user_id FROM users WHERE email = 'bugentry@hotmail.com'),
    'McKenzie',
    'Hambsch',
    (SELECT image_id FROM images WHERE image_path = 'mckenzie-hambsch-profile.jpg')
  );

INSERT INTO memberships (plan_id, member_id, active)
VALUES
  (
    (SELECT plan_id FROM plans WHERE plan_name = 'All Access' AND price = 0),
    (SELECT member_id FROM members WHERE first_name = 'Lukas'),
    true
  ),
  (
    (SELECT plan_id FROM plans WHERE plan_name = 'All Access' AND price = 0),
    (SELECT member_id FROM members WHERE first_name = 'McKenzie'),
    true
  );

INSERT INTO days (day_id, day_name)
VALUES
  (1, 'Sunday'),
  (2, 'Monday'),
  (3, 'Tuesday'),
  (4, 'Wednesday'),
  (5, 'Thursday'),
  (6, 'Friday'),
  (7, 'Saturday');

INSERT INTO gyms (gym_name)
VALUES
  ('24 Hour Fitness'),
  ('LA Fitness'),
  ('Crunch Fitness'),
  ('YMCA');

INSERT INTO addresses (country, state_region, city, postal_area, street_address)
VALUES
  ('USA', 'CA', 'San Diego', '92122', '4425 La Jolla Village Dr'),
  ('USA', 'CA', 'San Diego', '92111', '7715 Balboa Ave');

INSERT INTO gym_locations (gym_id, address_id, location_name, in_network)
VALUES
  (
    (SELECT gym_id FROM gyms WHERE gym_name = '24 Hour Fitness'),
    (SELECT address_id from addresses WHERE postal_area = '92122'),
    'Westfield UTC',
    false
  ),
  (
    (SELECT gym_id FROM gyms WHERE gym_name = '24 Hour Fitness'),
    (SELECT address_id from addresses WHERE postal_area = '92111'),
    'Balboa',
    false
  );

INSERT INTO business_hours (gym_location_id, day_id, open_time, close_time)
VALUES
  (
    (SELECT gym_location_id FROM gym_locations WHERE location_name = 'Westfield UTC'),
    1,
    '08:00 AM',
    '08:00 PM'
  ),
  (
    (SELECT gym_location_id FROM gym_locations WHERE location_name = 'Westfield UTC'),
    2,
    '06:00 AM',
    '10:00 PM'
  ),
  (
    (SELECT gym_location_id FROM gym_locations WHERE location_name = 'Westfield UTC'),
    3,
    '06:00 AM',
    '10:00 PM'
  ),
  (
    (SELECT gym_location_id FROM gym_locations WHERE location_name = 'Westfield UTC'),
    4,
    '06:00 AM',
    '10:00 PM'
  ),
  (
    (SELECT gym_location_id FROM gym_locations WHERE location_name = 'Westfield UTC'),
    5,
    '06:00 AM',
    '10:00 PM'
  ),
  (
    (SELECT gym_location_id FROM gym_locations WHERE location_name = 'Westfield UTC'),
    6,
    '06:00 AM',
    '10:00 PM'
  ),
  (
    (SELECT gym_location_id FROM gym_locations WHERE location_name = 'Westfield UTC'),
    7,
    '06:00 AM',
    '08:00 PM'
  ),
  (
    (SELECT gym_location_id FROM gym_locations WHERE location_name = 'Balboa'),
    1,
    '08:00 AM',
    '08:00 PM'
  ),
  (
    (SELECT gym_location_id FROM gym_locations WHERE location_name = 'Balboa'),
    2,
    '06:00 AM',
    '10:00 PM'
  ),
  (
    (SELECT gym_location_id FROM gym_locations WHERE location_name = 'Balboa'),
    3,
    '06:00 AM',
    '10:00 PM'
  ),
  (
    (SELECT gym_location_id FROM gym_locations WHERE location_name = 'Balboa'),
    4,
    '06:00 AM',
    '10:00 PM'
  ),
  (
    (SELECT gym_location_id FROM gym_locations WHERE location_name = 'Balboa'),
    5,
    '06:00 AM',
    '10:00 PM'
  ),
  (
    (SELECT gym_location_id FROM gym_locations WHERE location_name = 'Balboa'),
    6,
    '06:00 AM',
    '10:00 PM'
  ),
  (
    (SELECT gym_location_id FROM gym_locations WHERE location_name = 'Balboa'),
    7,
    '06:00 AM',
    '08:00 PM'
  );

INSERT INTO visits (member_id, gym_location_id, status_id)
VALUES
  (
    (SELECT member_id FROM members WHERE first_name = 'Lukas'),
    (SELECT gym_location_id FROM gym_locations WHERE location_name = 'Westfield UTC'),
    (SELECT status_id FROM statuses WHERE status_name = 'Pending')
  ),
  (
    (SELECT member_id FROM members WHERE first_name = 'McKenzie'),
    (SELECT gym_location_id FROM gym_locations WHERE location_name = 'Westfield UTC'),
    (SELECT status_id FROM statuses WHERE status_name = 'Pending')
  ),
  (
    (SELECT member_id FROM members WHERE first_name = 'Lukas'),
    (SELECT gym_location_id FROM gym_locations WHERE location_name = 'Balboa'),
    (SELECT status_id FROM statuses WHERE status_name = 'Pending')
  ),
  (
    (SELECT member_id FROM members WHERE first_name = 'McKenzie'),
    (SELECT gym_location_id FROM gym_locations WHERE location_name = 'Balboa'),
    (SELECT status_id FROM statuses WHERE status_name = 'Pending')
  ),
  (
    (SELECT member_id FROM members WHERE first_name = 'Lukas'),
    (SELECT gym_location_id FROM gym_locations WHERE location_name = 'Westfield UTC'),
    (SELECT status_id FROM statuses WHERE status_name = 'Pending')
  );
