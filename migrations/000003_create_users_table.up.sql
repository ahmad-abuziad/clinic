CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    name text NOT NULL,
    email citext UNIQUE NOT NULL,
    password_hash bytea NOT NULL,
    activated bool NOT NULL,
    version integer NOT NULL DEFAULT 1
);

INSERT INTO users (name, email, password_hash, activated) 
VALUES 
    ('The Doctor', 'doctor@example.com', '$2a$12$zIo/MOC.otJXueYWJqapbudTnY2jHewCjSDKKmrpVgtMCTx91gpiG', true),
    ('The Receptionist', 'receptionist@example.com', '$2a$12$zIo/MOC.otJXueYWJqapbudTnY2jHewCjSDKKmrpVgtMCTx91gpiG', true);
