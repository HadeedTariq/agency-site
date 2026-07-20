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
	"Zero-Trust Security Overhaul for Global FinTech",
	"SaaS Platform Scaling to Millions of Active Users",
	"Legacy System Modernization for Telecom Provider",
	"Cloud-Native Core Banking System Migration",
	"Kubernetes Orchestration for Multi-Region DevOps",
	"Disaster Recovery Architecture for Legal Enterprise",
	"Migrating Monolithic Architectures to Serverless Microservices",
	"Hybrid Cloud Implementation for High-Frequency Trading",
	"API-First Replatforming for National Insurance Carrier",

	// Supply Chain, Logistics & Manufacturing
	"Supply Chain Optimization for Logistics Giant",
	"Predictive Maintenance Implementation in Aviation",
	"IoT Smart Warehouse Deployment Success Story",
	"Fleet Telematics Integration for Global Shipping",
	"Automating Procurement Workflows for Heavy Industry",
	"Digital Freight Matching Engine Implementation",
	"Cold Chain Visibility Enhancement for Food Logistics",
	"Smart Factory Automation via Edge Computing",
	"Just-In-Time Inventory Management System Rollout",

	// Data Engineering, Analytics & AI
	"Data Lakehouse Architecture for Media Enterprise",
	"Predictive Churn Analytics for Subscription Service",
	"Computer Vision Implementation for Quality Control",
	"Natural Language Processing for Document Automation",
	"Real-Time Fraud Detection Engine for Digital Wallet",
	"Enterprise Data Governance Framework Deployment",
	"AI-Driven Dynamic Pricing Engine for Hospitality",
	"Personalization Recommendation Graph for Streaming Giant",
	"Automating Credit Scoring with Machine Learning Models",

	// Healthcare, Pharma & Compliance
	"Automating Regulatory Compliance for Pharma",
	"Telehealth Platform Scale-Up During Demand Spikes",
	"Interoperable Electronic Health Record (EHR) Integration",
	"HIPAA-Compliant Patient Portal Re-Engineering",
	"Clinical Trial Data Analytics Modernization",
	"Medical Device Remote Monitoring IoT Network",
	"Decentralized Research Database for Genomic Science",
	"Optimizing Hospital Resource Allocation via Analytics",

	// Retail, Marketing & Customer Experience
	"Omnichannel Customer Experience Replatforming",
	"Headless Commerce Migration for Luxury Retailer",
	"Loyalty Program Gamification Driving Higher Retention",
	"Automated Customer Journey Mapping Engine Success",
	"Point of Sale (POS) Hardware Ecosystem Modernization",
	"Hyper-Local Digital Marketing Attribute Platform",

	// Energy, Government & Special Fields
	"Grid Modernization Analytics for Energy Provider",
	"Smart City Traffic Management System Implementation",
	"Renewable Energy Asset Portfolio Tracking System",
	"Public Sector Digital Identity Verification Infrastructure",
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

	gofakeit.Seed(38)

	tx, err := database.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	q := db.New(tx)

	// if err := seedCategory(ctx, q, "BLOG", blogTitles, 10); err != nil {
	// 	return err
	// }

	// if err := seedCategory(ctx, q, "NEWSROOM", newsroomTitles, 10); err != nil {
	// 	return err
	// }

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
