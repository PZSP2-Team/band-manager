CREATE TABLE "groups" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255)
);

CREATE TABLE "subgroups" (
    id SERIAL PRIMARY KEY,
    group_id INTEGER REFERENCES "groups"(id),
    name VARCHAR(255),
    description VARCHAR(255)
);

CREATE TABLE "users" (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(10),
    group_id INTEGER REFERENCES "groups"(id)
);

CREATE TABLE "announcements" (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255),
    sender_id INTEGER REFERENCES "users"(id), 
    description VARCHAR(255),
    priority INTEGER,
    group_id INTEGER REFERENCES "groups"(id)
);

CREATE TABLE "events" (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255),
    description VARCHAR(255),
    date DATE,
    location VARCHAR(255),
    group_id INTEGER REFERENCES "groups"(id)
);

CREATE TABLE "tracks" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255)
);

CREATE TABLE "notesheets" (
    id SERIAL PRIMARY KEY,
    track_id INTEGER REFERENCES "tracks"(id),
    filepath VARCHAR(255)
);

CREATE TABLE "announcement_subgroup" (
    announcement_id INTEGER REFERENCES "announcements"(id),
    subgroup_id INTEGER REFERENCES "subgroups"(id),
    PRIMARY KEY (announcement_id, subgroup_id)
);

CREATE TABLE "subgroup_user" (
    subgroup_id INTEGER REFERENCES "subgroups"(id),
    user_id INTEGER REFERENCES "users"(id),
    PRIMARY KEY (subgroup_id, user_id)
);

CREATE TABLE "notesheet_subgroup" (
    notesheet_id INTEGER REFERENCES "notesheets"(id),
    subgroup_id INTEGER REFERENCES "subgroups"(id),
    PRIMARY KEY (notesheet_id, subgroup_id)
);

CREATE TABLE "event_track" (
    event_id INTEGER REFERENCES "events"(id),
    track_id INTEGER REFERENCES "tracks"(id),
    PRIMARY KEY (event_id, track_id)
);
