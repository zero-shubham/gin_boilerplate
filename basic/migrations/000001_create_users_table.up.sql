CREATE TABLE IF NOT EXISTS users(
  id UUID PRIMARY KEY,
  username text unique not null,
  password text not null,
  created_at timestamp with time zone,
  updated_at timestamp with time zone
)