CREATE TABLE IF NOT EXISTS playlist(
                                    "id" integer PRIMARY KEY,
                                    "songs" integer[],
                                    "status" VARCHAR(255) NOT NULL
    );