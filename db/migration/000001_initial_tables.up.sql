CREATE TABLE "emails" (
    id SERIAL PRIMARY KEY NOT NULL,
    dateSendEmail TIMESTAMPTZ NOT NULL DEFAULT  (now()),
    subjectEmail VARCHAR(100) NOT NULL,
    messageEmail VARCHAR(250) NOT NULL
);