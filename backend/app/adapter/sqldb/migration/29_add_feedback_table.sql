-- +migrate Up
CREATE TABLE "feedback"
(
    "app_id" VARCHAR(10) NOT NULL REFERENCES "app"("id"),
    "feedback_id" VARCHAR(5) NOT NULL,
    "customer_rating" SMALLINT NOT NULL,
    "comment" TEXT,
    "customer_email" VARCHAR(254),
    "received_at" TIMESTAMP WITH TIME ZONE NOT NULL,
    CONSTRAINT pk_feedback PRIMARY KEY ("app_id", "feedback_id")
);

-- +migrate Down
DROP TABLE "feedback";
