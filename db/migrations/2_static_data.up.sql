INSERT INTO plans (plan_name, price) VALUES
  ('All Access', 50);

INSERT INTO statuses (status_name) VALUES
  ('Pending'),
  ('Approved'),
  ('Denied - Identity'),
  ('Denied - Banned');

INSERT INTO holidays (holiday_name) VALUES
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

INSERT INTO features (feature_name, feature_description) VALUES
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

INSERT INTO users (email) VALUES
  ('lukas.hambsch@gmail.com');

INSERT INTO roles (role_name) VALUES
  ('admin'),
  ('gym'),
  ('location'),
  ('member');

INSERT INTO user_roles (user_id, role_id) VALUES
  (
    (SELECT user_id FROM users WHERE email = 'lukas.hambsch@gmail.com'),
    (SELECT role_id FROM roles WHERE role_name = 'admin')
  );
