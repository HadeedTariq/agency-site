-- name: GetHomePageInsights :many
WITH
blogs AS (
    SELECT
        id,
        title,
        hero_image,
        slug,
        category,
        published_at
    FROM insights
    WHERE
        category = 'BLOG'
        AND status = 'PUBLISHED'
    ORDER BY published_at DESC
    LIMIT 3
),

case_studies AS (
    SELECT
        id,
        title,
        hero_image,
        slug,
        category,
        published_at
    FROM insights
    WHERE
        category = 'CASE_STUDY'
        AND status = 'PUBLISHED'
    ORDER BY published_at DESC
    LIMIT 3
),

newsroom AS (
    SELECT
        id,
        title,
        hero_image,
        slug,
        category,
        published_at
    FROM insights
    WHERE
        category = 'NEWSROOM'
        AND status = 'PUBLISHED'
    ORDER BY published_at DESC
    LIMIT 3
)

SELECT
    id,
    title,
    hero_image,
    slug,
    category
FROM blogs

UNION ALL

SELECT
    id,
    title,
    hero_image,
    slug,
    category
FROM case_studies

UNION ALL

SELECT
    id,
    title,
    hero_image,
    slug,
    category
FROM newsroom;

-- name: CreateInsight :one
INSERT INTO insights (
    title,
    slug,
    excerpt,
    hero_image,
    content_markdown,
    category,
    status,
    featured,
    published_at
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?
)
RETURNING *;
-- name: GetNewsRoomDetails :one
SELECT *
FROM insights
WHERE slug = sqlc.arg(slug)
  AND category = 'NEWSROOM';
-- name: GetBlogDetails :one
SELECT *
FROM insights
WHERE slug = sqlc.arg(slug)
  AND category = 'BLOG';
-- name: GetCaseStudyDetails :one
SELECT *
FROM insights
WHERE slug = sqlc.arg(slug)
  AND category = 'CASE_STUDY';
