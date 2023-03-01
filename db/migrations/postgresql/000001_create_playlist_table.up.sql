CREATE TABLE IF NOT EXISTS playlist(
                                    "id" serial PRIMARY KEY,
                                    "song_name" VARCHAR(255) NOT NULL,
                                    "status" VARCHAR(255) NOT NULL
    );