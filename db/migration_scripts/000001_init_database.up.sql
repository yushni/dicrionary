CREATE TABLE "words" (
	"id" serial NOT NULL,
	"created_by" integer NOT NULL,
	"word" VARCHAR(255) NOT NULL,
	"transcription" VARCHAR(255) NOT NULL,
	"origin" VARCHAR(255) NOT NULL,
	"created_at" timestamp with time zone NOT NULL,
	"updated_at" timestamp with time zone NOT NULL,
	CONSTRAINT "words_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "translations" (
	"id" serial NOT NULL,
	"word_id" integer NOT NULL,
	"created_by" integer NOT NULL,
	"translation" VARCHAR(255) NOT NULL,
	"transcription" VARCHAR(255) NOT NULL,
	"frequency" integer NOT NULL,
	"created_at" timestamp with time zone NOT NULL,
	"updated_at" timestamp with time zone NOT NULL,
	CONSTRAINT "translations_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "samples" (
	"id" serial NOT NULL,
	"translation_id" integer NOT NULL,
	"created_by" integer NOT NULL,
	"sample" VARCHAR(255) NOT NULL,
	"created_at" timestamp with time zone NOT NULL,
	"updated_at" timestamp with time zone NOT NULL,
	CONSTRAINT "samples_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "users" (
	"id" serial NOT NULL,
	"first_name" VARCHAR(255) NOT NULL,
	"last_name" VARCHAR(255) NOT NULL,
	"token" VARCHAR(255) NOT NULL,
	"created_at" timestamp with time zone NOT NULL,
	"updated_at" timestamp with time zone NOT NULL,
	"deleted_at" timestamp with time zone,
	CONSTRAINT "users_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



ALTER TABLE "words" ADD CONSTRAINT "words_user_id_fk" FOREIGN KEY ("created_by") REFERENCES "users"("id");

ALTER TABLE "translations" ADD CONSTRAINT "translations_word_id_fk" FOREIGN KEY ("word_id") REFERENCES "words"("id");
ALTER TABLE "translations" ADD CONSTRAINT "translations_user_id_fk" FOREIGN KEY ("created_by") REFERENCES "users"("id");

ALTER TABLE "samples" ADD CONSTRAINT "samples_translation_id_fk" FOREIGN KEY ("translation_id") REFERENCES "translations"("id");
ALTER TABLE "samples" ADD CONSTRAINT "samples_user_id_fk" FOREIGN KEY ("created_by") REFERENCES "users"("id");