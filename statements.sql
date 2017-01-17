-- noinspection SqlResolveForFile
--name: create-tables
CREATE TABLE IF NOT EXISTS projects (
  id          SERIAL PRIMARY KEY,
  title       VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  date        DATE NOT NULL
);
CREATE TABLE IF NOT EXISTS project_images (
  id          SERIAL PRIMARY KEY,
  project_id  SERIAL NOT NULL,
  text        TEXT NOT NULL,
  url         VARCHAR(255) NOT NULL,
  index       SMALLINT NOT NULL,
  FOREIGN KEY (project_id) REFERENCES projects(id)
);
CREATE TABLE IF NOT EXISTS home_images (
  id    SERIAL PRIMARY KEY,
  url   VARCHAR(255) NOT NULL,
  index SMALLINT NOT NULL
);
CREATE TABLE IF NOT EXISTS portfolio_images (
  id            SERIAL PRIMARY KEY,
  thumbnail_url VARCHAR(255) NOT NULL,
  image_url     VARCHAR(255) NOT NULL,
  title         VARCHAR(255) NOT NULL,
  description   TEXT NOT NULL,
  index         SMALLINT DEFAULT NULL,
  project_id    INT DEFAULT NULL
);
CREATE TABLE IF NOT EXISTS social_accounts (
  id    VARCHAR(255) PRIMARY KEY NOT NULL,
  link  VARCHAR(255) NOT NULL
);
CREATE TABLE IF NOT EXISTS session (
  key     CHAR(16) NOT NULL PRIMARY KEY,
  data    BYTEA,
  expiry  INTEGER NOT NULL
);

--name: insert-project
INSERT INTO projects (title, description, date) VALUES ($1, $2, $3);

--name: update-project
UPDATE projects SET title=$1, description=$2, date=$3 WHERE id=$4;

--name: delete-project
DELETE FROM projects WHERE id=$1;

--name: insert-project-image
INSERT INTO project_images (project_id, text, url, index) VALUES ($1, $2, $3, $4);

--name: update-project-image
UPDATE project_images SET project_id=$1, text=$2, url=$3, index=$4 WHERE id=$5;

--name: delete-project-image
DELETE FROM project_images WHERE id=$1;



--name: select-home-images
SELECT * FROM home_images;

--name: select-home-images-max-index
SELECT MAX(index) FROM home_images;

--name: insert-home-image
INSERT INTO home_images (url, index) VALUES ($1, $2);

--name: update-home-image
UPDATE home_images SET url=$1 WHERE id=$2;

--name: update-home-image-order
UPDATE home_images SET index=$1 WHERE id=$2;

--name: delete-home-image
DELETE FROM home_images WHERE id=$1;



--name: select-portfolio-image
SELECT * FROM portfolio_images WHERE id=$1;

--name: select-portfolio-images
SELECT * FROM portfolio_images;

--name: select-portfolio-images-max-index
SELECT MAX(index) FROM portfolio_images;

--name: insert-portfolio-image
INSERT INTO portfolio_images (thumbnail_url, image_url, title, description, index, project_id) VALUES ($1, $2, $3, $4, $5, $6);

--name: update-portfolio-image
UPDATE portfolio_images SET thumbnail_url=$1, image_url=$2, title=$3, description=$4, project_id=$5 WHERE id=$6;

--name: update-portfolio-image-thumbnail
UPDATE portfolio_images SET thumbnail_url=$1, title=$2, description=$3, project_id=$4 WHERE id=$5;

--name: update-portfolio-image-main
UPDATE portfolio_images SET image_url=$1, title=$2, description=$3, project_id=$4 WHERE id=$5;

--name: update-portfolio-image-info
UPDATE portfolio_images SET title=$1, description=$2, project_id=$3 WHERE id=$4;

--name: update-portfolio-image-order
UPDATE portfolio_images SET index=$1 WHERE id=$2;

--name: delete-portfolio-image
DELETE FROM portfolio_images WHERE id=$1;



--name: select-social-accounts
SELECT * FROM social_accounts;

--name: select-social-account
SELECT * FROM social_accounts WHERE id=$1;

--name: insert-social-account
INSERT INTO social_accounts (id, link) VALUES ($1, $2);

--name: update-social-account
UPDATE social_accounts SET link=$1 WHERE id=$2;

--name: delete-social-account
DELETE FROM social_accounts WHERE id=$1;