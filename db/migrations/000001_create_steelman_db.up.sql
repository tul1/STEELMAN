-- https://stackoverflow.com/questions/22446478/extension-exists-but-uuid-generate-v4-fails
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "user"(
    "id" UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    "username" VARCHAR(50) UNIQUE NOT NULL,
    "password" VARCHAR(50) NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS "sport"(
    "id" UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    "sport_name" VARCHAR(50) UNIQUE NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS "challenge"(
    "id" UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    "challenge_name" VARCHAR(50) UNIQUE NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "sport_id" UUID  NOT NULL REFERENCES sport(id),
    "distance_goal" VARCHAR(50),
    "time_goal" VARCHAR(50)
);

CREATE TABLE IF NOT EXISTS "activity"(
    "id" UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "date" VARCHAR(50),
    "time" VARCHAR(50),
    "distance" VARCHAR(50),
    -- https://stackoverflow.com/questions/69696581/postgresql-migration-syntax-error-with-user-reference-column
    "user_id" UUID NOT NULL REFERENCES "user"(id),
    "sport_id" UUID NOT NULL REFERENCES sport(id),
    "challenge_id" UUID REFERENCES challenge(id)
);
CREATE INDEX ON activity(user_id);
