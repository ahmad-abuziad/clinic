CREATE TABLE IF NOT EXISTS permissions (
    id bigserial PRIMARY KEY,
    code text NOT NULL
);

CREATE TABLE IF NOT EXISTS users_permissions (
    user_id bigserial NOT NULL REFERENCES users ON DELETE CASCADE,
    permission_id bigserial NOT NULL REFERENCES permissions ON DELETE CASCADE,
    PRIMARY KEY (user_id, permission_id)
);

INSERT INTO permissions (code)
VALUES
    ('patients:read'),
    ('patients:write');

INSERT INTO users_permissions (user_id, permission_id)
VALUES
    (1, 1),
    (2, 2);