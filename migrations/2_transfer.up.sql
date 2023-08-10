CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS Transfer (
    id uuid DEFAULT uuid_generate_v4 (),
    account_origin_id  VARCHAR NOT NULL,
    account_destination_id  VARCHAR NOT NULL,
    amount INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

