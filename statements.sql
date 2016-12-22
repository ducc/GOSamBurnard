--name: create-tables
CREATE TABLE projects (
  id SERIAL8 NOT NULL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  date DATE NOT NULL
);
CREATE TABLE project_images (
  id SERIAL8 NOT NULL PRIMARY KEY,
  project_id SERIAL8 NOT NULL,
  text TEXT NOT NULL,
  image BYTEA NOT NULL,
  index SMALLINT NOT NULL
);
CREATE TABLE home_images (
  id SERIAL8 NOT NULL PRIMARY KEY,
  image BYTEA NOT NULL,
  index SMALLINT NOT NULL
);
CREATE TABLE portfolio_images (
  id SERIAL8 NOT NULL PRIMARY KEY,
  image BYTEA NOT NULL,
  title VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  index SMALLINT NOT NULL,
  project_id SERIAL8 NOT NULL
);
CREATE TABLE information (
  about TEXT NOT NULL,
  contact TEXT NOT NULL
);
CREATE TABLE social_accounts (
  id VARCHAR(255) NOT NULL,
  icon VARCHAR(63) NOT NULL,
  link VARCHAR(255) NOT NULL
);