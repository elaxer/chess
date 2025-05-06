CREATE TABLE "public"."games" (
    "id" uuid NOT NULL,
    "created_at" timestamp NOT NULL,
    "player_white_id" uuid NOT NULL,
    "player_black_id" uuid NOT NULL,
    "moves" text NOT NULL,
    "result" integer NOT NULL,
    CONSTRAINT "games_id" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE TABLE "public"."users" (
    "id" uuid NOT NULL,
    "created_at" timestamp NOT NULL,
    "login" character varying NOT NULL,
    "password_hash" character varying NOT NULL,
    "is_bot" boolean DEFAULT false NOT NULL,
    "rating" integer NOT NULL,
    CONSTRAINT "users_id" PRIMARY KEY ("id")
) WITH (oids = false);


ALTER TABLE ONLY "public"."games" ADD CONSTRAINT "games_player_black_id_fkey" FOREIGN KEY (player_black_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE NOT DEFERRABLE;
ALTER TABLE ONLY "public"."games" ADD CONSTRAINT "games_player_white_id_fkey" FOREIGN KEY (player_white_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE NOT DEFERRABLE;