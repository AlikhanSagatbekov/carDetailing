CREATE TABLE IF NOT EXISTS appointments (
    id BIGSERIAL PRIMARY KEY,
    customer_id BIGINT NOT NULL,
    service_id BIGINT NOT NULL,
    date VARCHAR(255) NOT NULL,
    status VARCHAR(50) NOT NULL
);
