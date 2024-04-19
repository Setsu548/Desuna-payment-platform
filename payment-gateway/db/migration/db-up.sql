
CREATE TABLE public."users" (
    "username"            varchar PRIMARY KEY,
    "hashed_password"     varchar NOT NULL,
    "full_name"           varchar NOT NULL,
    "email"               varchar UNIQUE NOT NULL,
    "account_number"      bigint NULL DEFAULT(0),
    "bank_name"           varchar NULL,
    "password_changed_at" timestamptz NOT NULL DEFAULT('01-01-0001 00:00:00Z'),
    "created_at"          timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE public."Payments" (
    "id"           bigserial PRIMARY KEY,
    "to_id_user"   varchar NOT NULL, -- foreign key
    "amount"       decimal(10,2) NOT NULL,
    "type"         varchar NOT NULL,
    "email"        varchar NOT NULL,
    "card_number"  bigint NOT NULL,
    "card_name"    varchar NOT NULL,
    "expire_date"  varchar NOT NULL,
    "refunded" BOOLEAN NOT NULL,
    "created_at"   timestamptz NOT NULL DEFAULT (now()),
    "updated_at"   timestamptz NULL,
    "deleted_at"   timestamptz NULL
);

ALTER TABLE "Payments" ADD FOREIGN KEY ("to_id_user") REFERENCES "users" ("username");
