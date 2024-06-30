
CREATE TABLE "group"(
    "id" UUID PRIMARY KEY,
    "unique_id" VARCHAR NOT NULL UNIQUE,
    "branch_id" UUID ,
    "type" varchar(255) NOT NULL CHECK (type IN ('intermediate', 'ielts', 'beginner', 'elementary')),
    "teacher_id" UUID,
    "support_teacher_id" UUID ,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP

);


CREATE TABLE "event"(
    "id" UUID PRIMARY KEY,
    "topic" VARCHAR NOT NULL,
    "date" VARCHAR,
    "start_time" TIME NOT NULL,
    "branch_id" UUID ,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE "assign_student"(
    "id" UUID PRIMARY KEY,
    "event_id" UUID REFERENCES "event"("id"),
    "student_id" UUID UNIQUE,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE "journal"(
    "id" UUID PRIMARY KEY,
    "group_id" UUID REFERENCES "group"("id"),
    "from_date" VARCHAR,
    "to_date" VARCHAR,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);


CREATE TABLE "schedule"(
    "id" UUID PRIMARY KEY,
    "journal_id" UUID REFERENCES "journal"("id"),
    "start_time" VARCHAR,
    "end_time" VARCHAR,
    "date" VARCHAR,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP

);

CREATE TABLE "lesson"(
    "id" UUID PRIMARY KEY,
    "schedule_id" UUID REFERENCES "schedule"("id") UNIQUE,
    "lesson" VARCHAR,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE "task"(
    "id" UUID PRIMARY KEY,
    "lesson_id" UUID REFERENCES "lesson"("id"),
    "label" VARCHAR,
    "deadline" VARCHAR,
    "score" NUMERIC,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);


CREATE TABLE "score"(
    "id" UUID PRIMARY KEY,
    "task_id" UUID REFERENCES "task"("id"),
    "student_id" UUID ,
    "score" INTEGER,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP

);
CREATE TABLE "do_task"(
    "id" UUID PRIMARY KEY,
    "task_id" UUID REFERENCES "task"("id"),
    "student_id" UUID ,
    "score" INTEGER,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP

);







