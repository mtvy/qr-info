
CREATE extension IF NOT EXISTS pg_stat_statements;

CREATE TABLE IF NOT EXISTS qrcodes_tb (
  id      serial primary key,
  url     VARCHAR(32),
  code_id VARCHAR(8),
  folder  VARCHAR(32),
  name    VARCHAR(12),
  path    VARCHAR(48),
  initer  VARCHAR(48),
  img_b   bytea
);
