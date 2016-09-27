create table projects (
  id text not null,
  user_id text not null,
  data jsonb not null,
  created_at timestamptz not null default now()
);
