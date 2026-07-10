package seed

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v7"

	db "agency-site/internal/db/queries"
)

const heroImage = "https://www.systemsltd.com/sites/default/files/2026-06/Image-01.png"

var blogTitles = []string{
	"Building Scalable Cloud Applications",
	"Modernizing Enterprise Infrastructure",
	"Why Every Business Needs Digital Transformation",
	"Improving Customer Experience Through AI",
	"Cybersecurity Best Practices for Enterprises",
	"Cloud Migration Strategies",
	"Building High Performance Teams",
	"Leveraging Data Analytics for Growth",
	"How DevOps Accelerates Software Delivery",
	"Future Trends in Enterprise Technology",
}

var newsroomTitles = []string{
	"Systems Ltd Announces New Digital Partnership",
	"Company Expands Global Operations",
	"Launch of New Enterprise Platform",
	"Awarded Technology Excellence Recognition",
	"Opening a New Regional Office",
	"Successful Completion of Government Project",
	"Quarterly Business Update",
	"New Strategic Collaboration Announced",
}

var caseStudyTitles = []string{
	"Digital Transformation for a Leading Bank",
	"Cloud Migration for Healthcare Provider",
	"ERP Implementation Success Story",
	"Retail Analytics Platform Case Study",
	"Modernizing Manufacturing Operations",
	"Scaling Infrastructure for E-commerce",
	"AI Powered Customer Support Solution",
	"Enterprise Automation Success Story",
}

func markdown(title string) string {
	return fmt.Sprintf(`# %s

%s

## Overview

%s

## Challenges

%s

## Solution

%s

## Results

- %s
- %s
- %s

## Conclusion

%s
`,
		title,
		gofakeit.Paragraph(2, 4, 15, " "),
		gofakeit.Paragraph(3, 4, 15, " "),
		gofakeit.Paragraph(3, 4, 15, " "),
		gofakeit.Paragraph(3, 4, 15, " "),
		gofakeit.Sentence(8),
		gofakeit.Sentence(8),
		gofakeit.Sentence(8),
		gofakeit.Paragraph(2, 4, 15, " "),
	)
}

func SeedInsights(
	ctx context.Context,
	database *sql.DB,
) error {

	gofakeit.Seed(20)

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	q := db.New(tx)

	if err := seedCategory(ctx, q, "BLOG", blogTitles, 10); err != nil {
		return err
	}

	if err := seedCategory(ctx, q, "NEWSROOM", newsroomTitles, 10); err != nil {
		return err
	}

	if err := seedCategory(ctx, q, "CASE_STUDY", caseStudyTitles, 10); err != nil {
		return err
	}

	return tx.Commit()
}

func seedCategory(
	ctx context.Context,
	q *db.Queries,
	category string,
	titles []string,
	count int,
) error {

	for i := 0; i < count; i++ {

		title := gofakeit.RandomString(titles)

		_, err := q.CreateInsight(ctx, db.CreateInsightParams{
			Title: title,

			Slug: gofakeit.UrlSlug(len(title)) + fmt.Sprintf("-%d", i),

			Excerpt: sql.NullString{
				String: gofakeit.Paragraph(1, 2, 18, " "),
				Valid:  true,
			},

			HeroImage: sql.NullString{
				String: heroImage,
				Valid:  true,
			},

			ContentMarkdown: markdown(title),

			Category: category,

			Status: "PUBLISHED",

			Featured: 1,

			PublishedAt: sql.NullTime{
				Time: gofakeit.DateRange(
					time.Now().AddDate(-1, 0, 0),
					time.Now(),
				),
				Valid: true,
			},
		})

		if err != nil {
			return err
		}
	}

	return nil
}
