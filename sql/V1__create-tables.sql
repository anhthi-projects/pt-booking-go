CREATE TABLE IF NOT EXISTS trainers (
  id int NOT NULL GENERATED ALWAYS AS IDENTITY,
  first_name text NOT NULL,
  last_name text NOT NULL,
  date_of_birth timestamp NOT NULL,
  level text NULL,
  year_of_exp int NULL,
  unique(id) 
)