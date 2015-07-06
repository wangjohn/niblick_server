-- -*- sql-product: postgresql -*-
CREATE TABLE rounds (
  id SERIAL,

  PRIMARY KEY (id)
);

CREATE TABLE round_holes (
  round_id INTEGER REFERENCES rounds (id)
  hole_id INTEGER REFERENCES holes (id)
);

CREATE TABLE holes (
  id SERIAL,
  score INTEGER NOT NULL,
  par INTEGER,
  yardage INTEGER,
  handicap INTEGER,
  green_in_reg BOOLEAN,
  fairway BOOLEAN,
  putts INTEGER,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),

  PRIMARY KEY (id)
);
