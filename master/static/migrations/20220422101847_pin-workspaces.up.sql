CREATE TABLE workspace_pins (
  id SERIAL PRIMARY KEY,
  workspace_id INT REFERENCES workspaces(id),
  user_id INT REFERENCES users(id),
  created_at TIMESTAMP with time zone NOT NULL DEFAULT NOW(),
  UNIQUE (user_id, workspace_id)
);

CREATE INDEX ix_workspace_pins ON public.workspace_pins USING btree (workspace_id);
