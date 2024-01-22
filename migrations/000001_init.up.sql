CREATE TYPE duration AS ENUM (
  'day'
);

CREATE TABLE IF NOT EXISTS resource_categories (
  id uuid PRIMARY KEY,
  title varchar NOT NULL,
  category_key varchar NOT NULL,
  allow_all_resources boolean NOT NULL DEFAULT true
);

CREATE TABLE IF NOT EXISTS resources (
  id uuid PRIMARY KEY,
  title varchar NOT NULL,
  resource_key varchar NOT NULL,
  category_id uuid NOT NULL REFERENCES resource_categories(id)
);

CREATE TABLE IF NOT EXISTS subscription_categories (
  id uuid PRIMARY KEY,
  title_uz varchar NOT NULL,
  title_ru varchar NOT NULL,
  title_en varchar NOT NULL,
  description_uz text NOT NULL,
  description_ru text NOT NULL,
  description_en text NOT NULL,
  image_uz varchar NOT NULL,
  image_ru varchar NOT NULL,
  image_en varchar NOT NULL,
  active boolean NOT NULL DEFAULT false,
  visible boolean NOT NULL DEFAULT true
);

CREATE TABLE IF NOT EXISTS resource_subsription_categories (
  id uuid PRIMARY KEY,
  resource_category_id uuid NOT NULL REFERENCES resource_categories(id),
  subscription_category_id uuid NOT NULL REFERENCES subscription_categories(id)
);

CREATE TABLE IF NOT EXISTS subscriptions (
  id uuid PRIMARY KEY,
  title_uz varchar NOT NULL,
  title_ru varchar NOT NULL,
  title_en varchar NOT NULL,
  active boolean NOT NULL DEFAULT false,
  price int NOT NULL DEFAULT 0,
  duration_type duration NOT NULL DEFAULT 'day',
  duration int NOT NULL,
  category_id uuid NOT NULL REFERENCES subscription_categories(id)
);

CREATE TABLE IF NOT EXISTS user_subscriptions (
  id uuid PRIMARY KEY,
  user_key uuid,
  subscription_id uuid NOT NULL REFERENCES subscriptions(id),
  start_time timestamp NOT NULL DEFAULT 'now()',
  end_time timestamp NOT NULL,
  active boolean NOT NULL DEFAULT false
);