CREATE TABLE IF NOT EXISTS events (
    id UUID PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description TEXT,
    start_time TIMESTAMP WITH TIME ZONE NOT NULL,
    end_time TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_events_start_time ON events(start_time);
CREATE INDEX IF NOT EXISTS idx_events_created_at ON events(created_at);
