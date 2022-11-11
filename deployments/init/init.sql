
CREATE extension IF NOT EXISTS pg_stat_statements;

CREATE TABLE IF NOT EXISTS qrcodes_tb (
  id serial primary key,
  code_id VARCHAR(8)
);
