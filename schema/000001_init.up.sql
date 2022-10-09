CREATE TABLE "user" (
                        "User_id" bigserial PRIMARY KEY,
                        "user_name" varchar NOT NULL,
                        "user_pass" varchar NOT NULL
);


CREATE TABLE "task" (
                        "id" bigserial PRIMARY KEY,
                        "task_name" varchar NOT NULL,
                        "task_description" text,
                        "status" varchar,
                        "created_at" date NOT NULL DEFAULT (now()),
                        "update_at" date NOT NULL DEFAULT (now()),
                        "owner_id" bigint NOT NULL REFERENCES "user" ("User_id") ON DELETE CASCADE
);
