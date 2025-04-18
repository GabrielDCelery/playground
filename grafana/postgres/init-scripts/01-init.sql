CREATE TABLE IF NOT EXISTS metrics (
    id SERIAL PRIMARY KEY,
    timestamp TIMESTAMP NOT NULL,
    value FLOAT NOT NULL,
    metric_name VARCHAR(255) NOT NULL
);

-- Insert some sample data
INSERT INTO metrics (timestamp, value, metric_name) VALUES
    (NOW() - INTERVAL '1 hour', 42.5, 'cpu_usage'),
    (NOW() - INTERVAL '45 minutes', 45.2, 'cpu_usage'),
    (NOW() - INTERVAL '30 minutes', 39.8, 'cpu_usage'),
    (NOW() - INTERVAL '15 minutes', 41.1, 'cpu_usage'),
    (NOW(), 43.7, 'cpu_usage');

