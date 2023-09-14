--CREATE DATABASE email_sender;

--\c email_sender

CREATE TABLE "emails" (
    idEmail SERIAL PRIMARY KEY NOT NULL,
    dateSendEmail TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    subjectEmail VARCHAR(100),
    messageEmail VARCHAR(250) NOT NULL
);