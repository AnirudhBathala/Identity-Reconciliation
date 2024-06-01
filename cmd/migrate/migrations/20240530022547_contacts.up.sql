CREATE TYPE precedence AS ENUM ('secondary','primary');
CREATE TABLE Contacts (
    contact_id SERIAL PRIMARY KEY,
    phone_number VARCHAR(255),
    email   VARCHAR(255),
    linked_id INT,
    link_precedence precedence,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
)