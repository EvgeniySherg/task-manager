CREATE TABLE "user" (
                        "id" bigserial PRIMARY KEY,
                        "name" varchar NOT NULL,
                        "pass" varchar NOT NULL
);


CREATE TABLE "task" (
                        "id" bigserial PRIMARY KEY,
                        "name" varchar NOT NULL,
                        "description" text,
                        "status" varchar,
                        "created_at" date NOT NULL DEFAULT (now()),
                        "update_at" date NOT NULL DEFAULT (now()),
                        "owner_id" bigint NOT NULL REFERENCES users ("User_id") ON DELETE CASCADE
);
