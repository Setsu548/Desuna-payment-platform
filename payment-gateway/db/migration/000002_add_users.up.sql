
-- -- CREATE UNIQUE INDEX ON "accounts" ("owner", "currency");
-- ALTER TABLE "accounts" ADD CONSTRAINT "owner_currency_key" UNIQUE ("owner", "currency");

CREATE TABLE "users" (
    "username"            varchar PRIMARY KEY,
    "hashed_password"     varchar NOT NULL,
    "full_name"           varchar NOT NULL,
    "email"               varchar UNIQUE NOT NULL,
    "account_number"      bigint NULL DEFAULT(0),
    "bank_name"           varchar NULL,
    "password_changed_at" timestamptz NOT NULL DEFAULT('01-01-0001 00:00:00Z'),
    "created_at"          timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Payments" (
    "id"          bigserial PRIMARY KEY,
    "card_number" bigint NOT NULL,
    "Card_name"   varchar NOT NULL,
    "Expire_date" varchar NOT NULL,
    "ccv_number"  int NOT NULL,
    "created_at"  timestamptz NOT NULL DEFAULT (now()),
    "Deleted_at"  timestamptz NULL
);

CREATE TABLE "transfers" (
    "id"              bigserial PRIMARY KEY,
    "Amount"          decimal(10,2) NOT NULL,
    "from_id_payment" bigint NOT NULL,
    "to_id_collect"   bigint NOT NULL,
    "date"            timestamptz NOT NULL DEFAULT (now()),
    "created_at"      timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "collect_payments" ADD FOREIGN KEY ("from_id_payment") REFERENCES "Payments" ("id");
ALTER TABLE "collect_payments" ADD FOREIGN KEY ("to_id_collect") REFERENCES "users" ("username");

