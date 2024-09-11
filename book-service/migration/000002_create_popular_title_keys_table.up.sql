CREATE TABLE IF NOT EXISTS popular_title_keys (
    title_key VARCHAR(100) NOT NULL PRIMARY KEY,
    total_search BIGINT NOT NULL DEFAULT 1
);

CREATE INDEX idx_total_search_desc ON popular_title_keys (total_search DESC);