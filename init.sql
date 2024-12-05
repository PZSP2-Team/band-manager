CREATE TABLE "groups" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255)
);

CREATE TABLE "subgroups" (
    id SERIAL PRIMARY KEY,
    group_id INTEGER REFERENCES "groups"(id) NOT NULL,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255)
);

CREATE TABLE "users" (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(10) NOT NULL CHECK (role IN ('Default', 'Moderator', 'Manager')),
    group_id INTEGER REFERENCES "groups"(id)
);

CREATE TABLE "announcements" (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    sender_id INTEGER REFERENCES "users"(id) NOT NULL, 
    description VARCHAR(255) NOT NULL,
    priority INTEGER NOT NULL,
    group_id INTEGER REFERENCES "groups"(id) NOT NULL
);

CREATE TABLE "events" (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    date DATE NOT NULL,
    location VARCHAR(255) NOT NULL,
    group_id INTEGER REFERENCES "groups"(id) NOT NULL
);

CREATE TABLE "tracks" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE "notesheets" (
    id SERIAL PRIMARY KEY,
    track_id INTEGER REFERENCES "tracks"(id) NOT NULL,
    filepath VARCHAR(255) NOT NULL
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

CREATE TABLE "performances" (
    event_id INTEGER REFERENCES "events"(id),
    track_id INTEGER REFERENCES "tracks"(id),
    PRIMARY KEY (event_id, track_id)
);
