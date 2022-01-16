-- Autogenerated by migration tool on 2022-01-16 11:12:09

-- Resource: demo_resource
CREATE TABLE IF NOT EXISTS "demo_table" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"account_id" text,
	"region" text,
	"creation_date" timestamp without time zone,
	"name" text,
	CONSTRAINT demo_table_pk PRIMARY KEY(cq_fetch_date,cq_id),
	UNIQUE(cq_fetch_date,cq_id)
);