CREATE TABLE "role"(
    "id" UUID PRIMARY KEY,
    "role" VARCHAR
);
CREATE TABLE "branch"(
  "id" UUID PRIMARY KEY,
  "name" VARCHAR,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP

);


CREATE TABLE "teacher"(
    "id" UUID PRIMARY KEY,
    "full_name" VARCHAR ,
    "phone" VARCHAR UNIQUE,
    "password" VARCHAR NOT NULL,
    "login" VARCHAR UNIQUE,
    "salary" NUMERIC,
    "ielts_score" VARCHAR,
    "branch_id" UUID REFERENCES "branch"("id"),
    "role_id" UUID REFERENCES "role"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);


CREATE TABLE "support_teacher"(
    "id" UUID PRIMARY KEY,
    "full_name" VARCHAR,
    "phone" VARCHAR UNIQUE,
    "password" VARCHAR NOT NULL,
    "login" VARCHAR UNIQUE,
    "salary" NUMERIC,
    "ielts_score" VARCHAR,
    "branch_id" UUID REFERENCES "branch"("id"),
    "role_id" UUID REFERENCES "role"("id"), 
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);




CREATE TABLE "super_admin"(
    "id" UUID PRIMARY KEY,
    "full_name" VARCHAR,
    "password" VARCHAR,
    "login" VARCHAR UNIQUE,
    "role_id" UUID REFERENCES "role"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE "manager"(
    "id" UUID PRIMARY KEY,
    "full_name" VARCHAR NOT NULL,
    "phone" VARCHAR NOT NULL,
    "salary" NUMERIC,
    "password" VARCHAR NOT NULL,
    "login" VARCHAR UNIQUE,
    "branch_id" UUID REFERENCES "branch"("id"),
    "role_id" UUID REFERENCES "role"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);




CREATE TABLE "administrator"(
    "id" UUID PRIMARY KEY,
    "full_name" VARCHAR,
    "phone" VARCHAR UNIQUE,
    "password" VARCHAR NOT NULL,
    "login" VARCHAR UNIQUE,
    "salary" NUMERIC,
    "ielts_score" VARCHAR,
    "branch_id" UUID REFERENCES "branch"("id"),
    "role_id" UUID REFERENCES "role"("id"), 
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);



CREATE TABLE "student"(
    "id" UUID PRIMARY KEY,
    "full_name" VARCHAR,
    "phone" VARCHAR UNIQUE,
    "password" VARCHAR NOT NULL,
    "login" VARCHAR UNIQUE,
    "group_id" UUID REFERENCES "group"("id"),
    "branch_id" UUID REFERENCES "branch"("id"),
    "role_id" UUID REFERENCES "role"("id"), 
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);





CREATE TABLE "payment"(
    "id" UUID PRIMARY KEY,
    "student_id" UUID REFERENCES "student"("id"),
    "branch_id" UUID REFERENCES "branch"("id"), 
    "paid_sum" NUMERIC,
    "total_sum" NUMERIC,
    "course_count" INTEGER,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);
