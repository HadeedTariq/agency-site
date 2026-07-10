CREATE TABLE insights (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    title TEXT NOT NULL,

    slug TEXT NOT NULL UNIQUE,

    excerpt TEXT,

    hero_image TEXT,

    content_markdown TEXT NOT NULL,

    category TEXT NOT NULL CHECK (
        category IN (
            'BLOG',
            'NEWSROOM',
            'CASE_STUDY'
        )
    ),

    status TEXT NOT NULL DEFAULT 'DRAFT' CHECK (
        status IN (
            'DRAFT',
            'PUBLISHED'
        )
    ),

    featured INTEGER NOT NULL DEFAULT 0 CHECK (
        featured IN (0, 1)
    ),

    published_at DATETIME,

    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_insights_category
ON insights(category);

CREATE INDEX idx_insights_status
ON insights(status);

CREATE INDEX idx_insights_published_at
ON insights(published_at);

CREATE INDEX idx_insights_featured
ON insights(featured);