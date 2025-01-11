--	PostgreSQL	database	dump
--

--	Dumped	from	database	version	16.6
--	Dumped	by	pg_dump	version	16.6

SET	statement_timeout	=	0;
SET	lock_timeout	=	0;
SET	idle_in_transaction_session_timeout	=	0;
SET	client_encoding	=	'UTF8';
SET	standard_conforming_strings	=	on;
SELECT	pg_catalog.set_config('search_path',	'',	false);
SET	check_function_bodies	=	false;
SET	xmloption	=	content;
SET	client_min_messages	=	warning;
SET	row_security	=	off;

SET	default_tablespace	=	'';

SET	default_table_access_method	=	heap;

--
--	Name:	announcement_subgroup;	Type:	TABLE;	Schema:	public;	Owner:	band_manager_user
--

CREATE	TABLE	public.announcement_subgroup	(
	subgroup_id	bigint	NOT	NULL,
	announcement_id	bigint	NOT	NULL
);


ALTER	TABLE	public.announcement_subgroup	OWNER	TO	band_manager_user;

--
--	Name:	announcements;	Type:	TABLE;	Schema:	public;	Owner:	band_manager_user
--

CREATE	TABLE	public.announcements	(
	id	bigint	NOT	NULL,
	title	text	NOT	NULL,
	description	text	NOT	NULL,
	priority	bigint	NOT	NULL,
	group_id	bigint	NOT	NULL,
	sender_id	bigint	NOT	NULL
);


ALTER	TABLE	public.announcements	OWNER	TO	band_manager_user;

--
--	Name:	announcements_id_seq;	Type:	SEQUENCE;	Schema:	public;	Owner:	band_manager_user
--

CREATE	SEQUENCE	public.announcements_id_seq
	START	WITH	1
	INCREMENT	BY	1
	NO	MINVALUE
	NO	MAXVALUE
	CACHE	1;


ALTER	SEQUENCE	public.announcements_id_seq	OWNER	TO	band_manager_user;

--
--	Name:	announcements_id_seq;	Type:	SEQUENCE	OWNED	BY;	Schema:	public;	Owner:	band_manager_user
--

ALTER	SEQUENCE	public.announcements_id_seq	OWNED	BY	public.announcements.id;


--
--	Name:	event_tracks;	Type:	TABLE;	Schema:	public;	Owner:	band_manager_user
--

CREATE	TABLE	public.event_tracks	(
	event_id	bigint	NOT	NULL,
	track_id	bigint	NOT	NULL
);


ALTER	TABLE	public.event_tracks	OWNER	TO	band_manager_user;

--
--	Name:	event_users;	Type:	TABLE;	Schema:	public;	Owner:	band_manager_user
--

CREATE	TABLE	public.event_users	(
	event_id	bigint	NOT	NULL,
	user_id	bigint	NOT	NULL
);


ALTER	TABLE	public.event_users	OWNER	TO	band_manager_user;

--
--	Name:	events;	Type:	TABLE;	Schema:	public;	Owner:	band_manager_user
--

CREATE	TABLE	public.events	(
	id	bigint	NOT	NULL,
	title	text	NOT	NULL,
	location	text	NOT	NULL,
	description	text,
	date	timestamp	with	time	zone,
	group_id	bigint	NOT	NULL
);


ALTER	TABLE	public.events	OWNER	TO	band_manager_user;

--
--	Name:	events_id_seq;	Type:	SEQUENCE;	Schema:	public;	Owner:	band_manager_user
--

CREATE	SEQUENCE	public.events_id_seq
	START	WITH	1
	INCREMENT	BY	1
	NO	MINVALUE
	NO	MAXVALUE
	CACHE	1;


ALTER	SEQUENCE	public.events_id_seq	OWNER	TO	band_manager_user;

--
--	Name:	events_id_seq;	Type:	SEQUENCE	OWNED	BY;	Schema:	public;	Owner:	band_manager_user
--

ALTER	SEQUENCE	public.events_id_seq	OWNED	BY	public.events.id;


--
--	Name:	groups;	Type:	TABLE;	Schema:	public;	Owner:	band_manager_user
--

CREATE	TABLE	public.groups	(
	id	bigint	NOT	NULL,
	name	text	NOT	NULL,
	access_token	text	NOT	NULL,
	description	text
);


ALTER	TABLE	public.groups	OWNER	TO	band_manager_user;

--
--	Name:	groups_id_seq;	Type:	SEQUENCE;	Schema:	public;	Owner:	band_manager_user
--

CREATE	SEQUENCE	public.groups_id_seq
	START	WITH	1
	INCREMENT	BY	1
	NO	MINVALUE
	NO	MAXVALUE
	CACHE	1;


ALTER	SEQUENCE	public.groups_id_seq	OWNER	TO	band_manager_user;

--
--	Name:	groups_id_seq;	Type:	SEQUENCE	OWNED	BY;	Schema:	public;	Owner:	band_manager_user
--

ALTER	SEQUENCE	public.groups_id_seq	OWNED	BY	public.groups.id;


--
--	Name:	notesheet_subgroup;	Type:	TABLE;	Schema:	public;	Owner:	band_manager_user
--

CREATE	TABLE	public.notesheet_subgroup	(
	notesheet_id	bigint	NOT	NULL,
	subgroup_id	bigint	NOT	NULL
);


ALTER	TABLE	public.notesheet_subgroup	OWNER	TO	band_manager_user;

--
--	Name:	notesheets;	Type:	TABLE;	Schema:	public;	Owner:	band_manager_user
--

CREATE	TABLE	public.notesheets	(
	id	bigint	NOT	NULL,
	filepath	text	NOT	NULL,
	track_id	bigint	NOT	NULL,
	instrument	text	NOT	NULL
);


ALTER	TABLE	public.notesheets	OWNER	TO	band_manager_user;

--
--	Name:	notesheets_id_seq;	Type:	SEQUENCE;	Schema:	public;	Owner:	band_manager_user
--

CREATE	SEQUENCE	public.notesheets_id_seq
	START	WITH	1
	INCREMENT	BY	1
	NO	MINVALUE
	NO	MAXVALUE
	CACHE	1;


ALTER	SEQUENCE	public.notesheets_id_seq	OWNER	TO	band_manager_user;

--
--	Name:	notesheets_id_seq;	Type:	SEQUENCE	OWNED	BY;	Schema:	public;	Owner:	band_manager_user
--

ALTER	SEQUENCE	public.notesheets_id_seq	OWNED	BY	public.notesheets.id;


--
--	Name:	subgroup_user;	Type:	TABLE;	Schema:	public;	Owner:	band_manager_user
--

CREATE	TABLE	public.subgroup_user	(
	subgroup_id	bigint	NOT	NULL,
	user_id	bigint	NOT	NULL
);


ALTER	TABLE	public.subgroup_user	OWNER	TO	band_manager_user;

--
--	Name:	subgroups;	Type:	TABLE;	Schema:	public;	Owner:	band_manager_user
--

CREATE	TABLE	public.subgroups	(
	id	bigint	NOT	NULL,
	group_id	bigint	NOT	NULL,
	name	text	NOT	NULL,
	description	text
);


ALTER	TABLE	public.subgroups	OWNER	TO	band_manager_user;

--
--	Name:	subgroups_id_seq;	Type:	SEQUENCE;	Schema:	public;	Owner:	band_manager_user
--

CREATE	SEQUENCE	public.subgroups_id_seq
	START	WITH	1
	INCREMENT	BY	1
	NO	MINVALUE
	NO	MAXVALUE
	CACHE	1;


ALTER	SEQUENCE	public.subgroups_id_seq	OWNER	TO	band_manager_user;

--
--	Name:	subgroups_id_seq;	Type:	SEQUENCE	OWNED	BY;	Schema:	public;	Owner:	band_manager_user
--

ALTER	SEQUENCE	public.subgroups_id_seq	OWNED	BY	public.subgroups.id;


--
--	Name:	track_event;	Type:	TABLE;	Schema:	public;	Owner:	band_manager_user
--

CREATE	TABLE	public.track_event	(
	track_id	bigint	NOT	NULL,
	event_id	bigint	NOT	NULL
);


ALTER	TABLE	public.track_event	OWNER	TO	band_manager_user;

--
--	Name:	tracks;	Type:	TABLE;	Schema:	public;	Owner:	band_manager_user
--

CREATE	TABLE	public.tracks	(
	id	bigint	NOT	NULL,
	name	text	NOT	NULL,
	group_id	bigint	NOT	NULL,
	description	text	NOT	NULL
);


ALTER	TABLE	public.tracks	OWNER	TO	band_manager_user;

--
--	Name:	tracks_id_seq;	Type:	SEQUENCE;	Schema:	public;	Owner:	band_manager_user
--

CREATE	SEQUENCE	public.tracks_id_seq
	START	WITH	1
	INCREMENT	BY	1
	NO	MINVALUE
	NO	MAXVALUE
	CACHE	1;


ALTER	SEQUENCE	public.tracks_id_seq	OWNER	TO	band_manager_user;

--
--	Name:	tracks_id_seq;	Type:	SEQUENCE	OWNED	BY;	Schema:	public;	Owner:	band_manager_user
--

ALTER	SEQUENCE	public.tracks_id_seq	OWNED	BY	public.tracks.id;


--
--	Name:	user_group;	Type:	TABLE;	Schema:	public;	Owner:	band_manager_user
--

CREATE	TABLE	public.user_group	(
	user_id	bigint	NOT	NULL,
	group_id	bigint	NOT	NULL
);


ALTER	TABLE	public.user_group	OWNER	TO	band_manager_user;

--
--	Name:	user_group_roles;	Type:	TABLE;	Schema:	public;	Owner:	band_manager_user
--

CREATE	TABLE	public.user_group_roles	(
	user_id	bigint	NOT	NULL,
	group_id	bigint	NOT	NULL,
	role	text	NOT	NULL
);


ALTER	TABLE	public.user_group_roles	OWNER	TO	band_manager_user;

--
--	Name:	users;	Type:	TABLE;	Schema:	public;	Owner:	band_manager_user
--

CREATE	TABLE	public.users	(
	id	bigint	NOT	NULL,
	first_name	text	NOT	NULL,
	last_name	text	NOT	NULL,
	email	text	NOT	NULL,
	password_hash	text	NOT	NULL
);


ALTER	TABLE	public.users	OWNER	TO	band_manager_user;

--
--	Name:	users_id_seq;	Type:	SEQUENCE;	Schema:	public;	Owner:	band_manager_user
--

CREATE	SEQUENCE	public.users_id_seq
	START	WITH	1
	INCREMENT	BY	1
	NO	MINVALUE
	NO	MAXVALUE
	CACHE	1;


ALTER	SEQUENCE	public.users_id_seq	OWNER	TO	band_manager_user;

--
--	Name:	users_id_seq;	Type:	SEQUENCE	OWNED	BY;	Schema:	public;	Owner:	band_manager_user
--

ALTER	SEQUENCE	public.users_id_seq	OWNED	BY	public.users.id;


--
--	Name:	announcements	id;	Type:	DEFAULT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.announcements	ALTER	COLUMN	id	SET	DEFAULT	nextval('public.announcements_id_seq'::regclass);


--
--	Name:	events	id;	Type:	DEFAULT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.events	ALTER	COLUMN	id	SET	DEFAULT	nextval('public.events_id_seq'::regclass);


--
--	Name:	groups	id;	Type:	DEFAULT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.groups	ALTER	COLUMN	id	SET	DEFAULT	nextval('public.groups_id_seq'::regclass);


--
--	Name:	notesheets	id;	Type:	DEFAULT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.notesheets	ALTER	COLUMN	id	SET	DEFAULT	nextval('public.notesheets_id_seq'::regclass);


--
--	Name:	subgroups	id;	Type:	DEFAULT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.subgroups	ALTER	COLUMN	id	SET	DEFAULT	nextval('public.subgroups_id_seq'::regclass);


--
--	Name:	tracks	id;	Type:	DEFAULT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.tracks	ALTER	COLUMN	id	SET	DEFAULT	nextval('public.tracks_id_seq'::regclass);


--
--	Name:	users	id;	Type:	DEFAULT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.users	ALTER	COLUMN	id	SET	DEFAULT	nextval('public.users_id_seq'::regclass);


--
--	Data	for	Name:	announcement_subgroup;	Type:	TABLE	DATA;	Schema:	public;	Owner:	band_manager_user
--

COPY	public.announcement_subgroup	(subgroup_id,	announcement_id)	FROM	stdin;
2	3
2	6
\.


--
--	Data	for	Name:	announcements;	Type:	TABLE	DATA;	Schema:	public;	Owner:	band_manager_user
--

COPY	public.announcements	(id,	title,	description,	priority,	group_id,	sender_id)	FROM	stdin;
1	Próba	Opis	1	1	1
3	Próba	Opis	1	1	1
4	Ogłoszenie	grupowe	Bez	podgrup	2	1	1
5	Próba	Opis	1	1	1
6	Próba	Opis	1	1	1
\.


--
--	Data	for	Name:	event_tracks;	Type:	TABLE	DATA;	Schema:	public;	Owner:	band_manager_user
--

COPY	public.event_tracks	(event_id,	track_id)	FROM	stdin;
4	1
1	1
5	1
6	1
7	1
8	1
9	1
10	1
11	1
12	1
13	1
14	1
15	1
16	1
18	6
1	2
17	6
\.


--
--	Data	for	Name:	event_users;	Type:	TABLE	DATA;	Schema:	public;	Owner:	band_manager_user
--

COPY	public.event_users	(event_id,	user_id)	FROM	stdin;
7	1
7	2
8	1
8	2
9	1
9	2
11	1
11	2
12	2
13	2
14	1
15	1
15	2
18	2
17	2
\.


--
--	Data	for	Name:	events;	Type:	TABLE	DATA;	Schema:	public;	Owner:	band_manager_user
--

COPY	public.events	(id,	title,	location,	description,	date,	group_id)	FROM	stdin;
2	Koncert	Noworoczny	Sala	koncertowa	Doroczny	koncert	orkiestry	2024-01-01	18:00:00+00	1
3	Koncert	Noworoczny	Sala	koncertowa	Doroczny	koncert	orkiestry	2024-01-01	18:00:00+00	1
4	Koncert	Noworoczny	Sala	koncertowa	Doroczny	koncert	orkiestry	2024-01-01	18:00:00+00	1
5	Koncert	Noworoczny	Sala	koncertowa	Doroczny	koncert	orkiestry	2024-01-01	18:00:00+00	1
6	Koncert	Noworoczny	Sala	koncertowa	Doroczny	koncert	orkiestry	2024-01-01	18:00:00+00	1
7	Koncert	Noworoczny	Sala	koncertowa	Doroczny	koncert	orkiestry	2024-01-01	18:00:00+00	1
8	Koncert	Noworoczny	Sala	koncertowa	Doroczny	koncert	orkiestry	2024-01-01	18:00:00+00	1
9	Koncert	Noworoczny	Sala	koncertowa	Doroczny	koncert	orkiestry	2024-01-01	18:00:00+00	1
10	Koncert	Noworoczny	Sala	koncertowa	Doroczny	koncert	orkiestry	2024-01-01	18:00:00+00	1
11	Koncert	Noworoczny	Sala	koncertowa	Doroczny	koncert	orkiestry	2024-01-01	18:00:00+00	1
12	Koncert	Noworoczny	Sala	koncertowa	Doroczny	koncert	orkiestry	2024-01-01	18:00:00+00	1
13	test	dla	tylko	2	Sala	koncertowa	Doroczny	koncert	orkiestry	2024-01-01	18:00:00+00	1
14	test	dla	tylko	1	Sala	koncertowa	Doroczny	koncert	orkiestry	2024-01-01	18:00:00+00	1
15	test	dla	wszystkich	Sala	koncertowa	Doroczny	koncert	orkiestry	2024-01-01	18:00:00+00	1
16	Koncert	Noworoczny	Sala	Koncertowa	Koncert	z	okazji	Nowego	Roku	2025-01-01	18:00:00+00	1
18	Koncert	Noworoczny	Sala	Koncertowa	Koncert	z	okazji	Nowego	Roku	2025-01-01	18:00:00+00	7
1	Wielki	Koncert	Noworoczny	Filharmonia	Koncert	z	okazji	Nowego	Roku	2025	2025-01-01	19:00:00+00	1
17	Wielki	Koncert	Noworoczny	Filharmonia	Koncert	z	okazji	Nowego	Roku	2025	2025-01-01	19:00:00+00	7
\.


--
--	Data	for	Name:	groups;	Type:	TABLE	DATA;	Schema:	public;	Owner:	band_manager_user
--

COPY	public.groups	(id,	name,	access_token,	description)	FROM	stdin;
1	Orkiestra	Dęta	a21494110d4431e4fa9844dbe128bb85	Główna	orkiestra
2	Orkiestra	Dęta	eabdbb474c118a9e9094100700dadd39	Główna	orkiestra
3	string	6d050d95776fd46430acd86a8328f7a7	string
4	gru	6d1f34f3156459012015b6bfb45e6a0b	string
5	gru	da630e21b20b9083e57d60ecfc727fd8	string
6	gru	c6f31b801a9aafbb38740eda791ca087	string
7	grupa	drugiego	usera	072367c0dd95fbe73b68ce06b6bc0126	grupa	drugiego	usera
8	grupa	drugiego	usera	c87f13c3dea5551ad6e5475cc5692422	grupa	drugiego	usera
\.


--
--	Data	for	Name:	notesheet_subgroup;	Type:	TABLE	DATA;	Schema:	public;	Owner:	band_manager_user
--

COPY	public.notesheet_subgroup	(notesheet_id,	subgroup_id)	FROM	stdin;
1	2
2	3
\.


--
--	Data	for	Name:	notesheets;	Type:	TABLE	DATA;	Schema:	public;	Owner:	band_manager_user
--

COPY	public.notesheets	(id,	filepath,	track_id,	instrument)	FROM	stdin;
1	/storage/nuty/marsz_trabka.pdf	1	trąbka
2	/nuty/marsz_trabka.pdf	6	Trąbkaa
\.


--
--	Data	for	Name:	subgroup_user;	Type:	TABLE	DATA;	Schema:	public;	Owner:	band_manager_user
--

COPY	public.subgroup_user	(subgroup_id,	user_id)	FROM	stdin;
2	2
2	3
\.


--
--	Data	for	Name:	subgroups;	Type:	TABLE	DATA;	Schema:	public;	Owner:	band_manager_user
--

COPY	public.subgroups	(id,	group_id,	name,	description)	FROM	stdin;
2	1	subgrupa	subgrup1
3	7	Sekcja	Dęta	Blaszana	Updated	Trąbki,	puzony,	tuby,	waltornie
\.


--
--	Data	for	Name:	track_event;	Type:	TABLE	DATA;	Schema:	public;	Owner:	band_manager_user
--

COPY	public.track_event	(track_id,	event_id)	FROM	stdin;
\.


--
--	Data	for	Name:	tracks;	Type:	TABLE	DATA;	Schema:	public;	Owner:	band_manager_user
--

COPY	public.tracks	(id,	name,	group_id,	description)	FROM	stdin;
1	Marsz	Triumfalny	1	Marsz	na	obchody	święta
2	Hallelujah	1	Famous	song	by	Leonard	Cohen
3	Amazing	Grace	1	Traditional	Christian	hymn
4	Marsz	Radeckiego	7	Marsz	wojskowy
5	Marsz	Radeckiego	1	Marsz	wojskowy	Straussa
6	Marsz	Radeckiego	7	Marsz	wojskowy	Straussa
\.


--
--	Data	for	Name:	user_group;	Type:	TABLE	DATA;	Schema:	public;	Owner:	band_manager_user
--

COPY	public.user_group	(user_id,	group_id)	FROM	stdin;
\.


--
--	Data	for	Name:	user_group_roles;	Type:	TABLE	DATA;	Schema:	public;	Owner:	band_manager_user
--

COPY	public.user_group_roles	(user_id,	group_id,	role)	FROM	stdin;
1	1	manager
1	2	manager
2	3	manager
3	3	member
3	4	manager
3	5	manager
3	6	manager
2	7	manager
2	8	manager
\.


--
--	Data	for	Name:	users;	Type:	TABLE	DATA;	Schema:	public;	Owner:	band_manager_user
--

COPY	public.users	(id,	first_name,	last_name,	email,	password_hash)	FROM	stdin;
2	Anna	Nowak	anna@example.com	$2a$10$WfsCl2mPVUjsTU.NS.tt/eJm7wrCJszZFLTs0J1tZb.h8eiQNYt0O
3	Jan	Kowalski	jan@examplee.com	$2a$10$Hb7VOtYvpyyNiVpjxIRGKO4.A85ZxfAP/KAfyS2Jk0i4UXRjnCVnm
36	ab	sd	string	$2a$10$klyhgtBqqqSgIwFLH0NoIubqvzkCvy9ZLNKjhfyvKIwcqaMd0FDfu
37	Jan	Kowalski	jan@test.pl	$2a$10$UJyT.9fesm5J4biQT0LXwO7h4syH0/DuA5uO9exfy5AufmKM3RDs.
1	Jan	Kowalski	jan@example.com	$2a$10$z.2uDq0.z6xspNu6oPtDheZ98eh8vYh5xT8KjGqq2XZD/R4ocuLqG
\.


--
--	Name:	announcements_id_seq;	Type:	SEQUENCE	SET;	Schema:	public;	Owner:	band_manager_user
--

SELECT	pg_catalog.setval('public.announcements_id_seq',	6,	true);


--
--	Name:	events_id_seq;	Type:	SEQUENCE	SET;	Schema:	public;	Owner:	band_manager_user
--

SELECT	pg_catalog.setval('public.events_id_seq',	18,	true);


--
--	Name:	groups_id_seq;	Type:	SEQUENCE	SET;	Schema:	public;	Owner:	band_manager_user
--

SELECT	pg_catalog.setval('public.groups_id_seq',	8,	true);


--
--	Name:	notesheets_id_seq;	Type:	SEQUENCE	SET;	Schema:	public;	Owner:	band_manager_user
--

SELECT	pg_catalog.setval('public.notesheets_id_seq',	2,	true);


--
--	Name:	subgroups_id_seq;	Type:	SEQUENCE	SET;	Schema:	public;	Owner:	band_manager_user
--

SELECT	pg_catalog.setval('public.subgroups_id_seq',	3,	true);


--
--	Name:	tracks_id_seq;	Type:	SEQUENCE	SET;	Schema:	public;	Owner:	band_manager_user
--

SELECT	pg_catalog.setval('public.tracks_id_seq',	6,	true);


--
--	Name:	users_id_seq;	Type:	SEQUENCE	SET;	Schema:	public;	Owner:	band_manager_user
--

SELECT	pg_catalog.setval('public.users_id_seq',	37,	true);


--
--	Name:	announcement_subgroup	announcement_subgroup_pkey;	Type:	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.announcement_subgroup
	ADD	CONSTRAINT	announcement_subgroup_pkey	PRIMARY	KEY	(subgroup_id,	announcement_id);


--
--	Name:	announcements	announcements_pkey;	Type:	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.announcements
	ADD	CONSTRAINT	announcements_pkey	PRIMARY	KEY	(id);


--
--	Name:	event_tracks	event_tracks_pkey;	Type:	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.event_tracks
	ADD	CONSTRAINT	event_tracks_pkey	PRIMARY	KEY	(event_id,	track_id);


--
--	Name:	event_users	event_users_pkey;	Type:	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.event_users
	ADD	CONSTRAINT	event_users_pkey	PRIMARY	KEY	(event_id,	user_id);


--
--	Name:	events	events_pkey;	Type:	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.events
	ADD	CONSTRAINT	events_pkey	PRIMARY	KEY	(id);


--
--	Name:	groups	groups_pkey;	Type:	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.groups
	ADD	CONSTRAINT	groups_pkey	PRIMARY	KEY	(id);


--
--	Name:	notesheet_subgroup	notesheet_subgroup_pkey;	Type:	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.notesheet_subgroup
	ADD	CONSTRAINT	notesheet_subgroup_pkey	PRIMARY	KEY	(notesheet_id,	subgroup_id);


--
--	Name:	notesheets	notesheets_pkey;	Type:	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.notesheets
	ADD	CONSTRAINT	notesheets_pkey	PRIMARY	KEY	(id);


--
--	Name:	subgroup_user	subgroup_user_pkey;	Type:	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.subgroup_user
	ADD	CONSTRAINT	subgroup_user_pkey	PRIMARY	KEY	(subgroup_id,	user_id);


--
--	Name:	subgroups	subgroups_pkey;	Type:	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.subgroups
	ADD	CONSTRAINT	subgroups_pkey	PRIMARY	KEY	(id);


--
--	Name:	track_event	track_event_pkey;	Type:	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.track_event
	ADD	CONSTRAINT	track_event_pkey	PRIMARY	KEY	(track_id,	event_id);


--
--	Name:	tracks	tracks_pkey;	Type:	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.tracks
	ADD	CONSTRAINT	tracks_pkey	PRIMARY	KEY	(id);


--
--	Name:	groups	uni_groups_access_token;	Type:	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.groups
	ADD	CONSTRAINT	uni_groups_access_token	UNIQUE	(access_token);


--
--	Name:	users	uni_users_email;	Type:	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.users
	ADD	CONSTRAINT	uni_users_email	UNIQUE	(email);


--
--	Name:	user_group	user_group_pkey;	Type:	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.user_group
	ADD	CONSTRAINT	user_group_pkey	PRIMARY	KEY	(user_id,	group_id);


--
--	Name:	user_group_roles	user_group_roles_pkey;	Type:	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.user_group_roles
	ADD	CONSTRAINT	user_group_roles_pkey	PRIMARY	KEY	(user_id,	group_id);


--
--	Name:	users	users_pkey;	Type:	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.users
	ADD	CONSTRAINT	users_pkey	PRIMARY	KEY	(id);


--
--	Name:	announcement_subgroup	fk_announcement_subgroup_announcement;	Type:	FK	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.announcement_subgroup
	ADD	CONSTRAINT	fk_announcement_subgroup_announcement	FOREIGN	KEY	(announcement_id)	REFERENCES	public.announcements(id)	ON	DELETE	CASCADE;


--
--	Name:	announcement_subgroup	fk_announcement_subgroup_subgroup;	Type:	FK	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.announcement_subgroup
	ADD	CONSTRAINT	fk_announcement_subgroup_subgroup	FOREIGN	KEY	(subgroup_id)	REFERENCES	public.subgroups(id)	ON	DELETE	CASCADE;


--
--	Name:	event_tracks	fk_event_tracks_event;	Type:	FK	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.event_tracks
	ADD	CONSTRAINT	fk_event_tracks_event	FOREIGN	KEY	(event_id)	REFERENCES	public.events(id);


--
--	Name:	event_tracks	fk_event_tracks_track;	Type:	FK	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.event_tracks
	ADD	CONSTRAINT	fk_event_tracks_track	FOREIGN	KEY	(track_id)	REFERENCES	public.tracks(id);


--
--	Name:	event_users	fk_event_users_event;	Type:	FK	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.event_users
	ADD	CONSTRAINT	fk_event_users_event	FOREIGN	KEY	(event_id)	REFERENCES	public.events(id)	ON	DELETE	CASCADE;


--
--	Name:	event_users	fk_event_users_user;	Type:	FK	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.event_users
	ADD	CONSTRAINT	fk_event_users_user	FOREIGN	KEY	(user_id)	REFERENCES	public.users(id)	ON	DELETE	CASCADE;


--
--	Name:	announcements	fk_groups_announcements;	Type:	FK	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.announcements
	ADD	CONSTRAINT	fk_groups_announcements	FOREIGN	KEY	(group_id)	REFERENCES	public.groups(id)	ON	DELETE	CASCADE;


--
--	Name:	events	fk_groups_events;	Type:	FK	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.events
	ADD	CONSTRAINT	fk_groups_events	FOREIGN	KEY	(group_id)	REFERENCES	public.groups(id)	ON	DELETE	CASCADE;


--
--	Name:	subgroups	fk_groups_subgroups;	Type:	FK	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.subgroups
	ADD	CONSTRAINT	fk_groups_subgroups	FOREIGN	KEY	(group_id)	REFERENCES	public.groups(id)	ON	DELETE	CASCADE;


--
--	Name:	tracks	fk_groups_tracks;	Type:	FK	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.tracks
	ADD	CONSTRAINT	fk_groups_tracks	FOREIGN	KEY	(group_id)	REFERENCES	public.groups(id)	ON	DELETE	CASCADE;


--
--	Name:	user_group_roles	fk_groups_user_roles;	Type:	FK	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.user_group_roles
	ADD	CONSTRAINT	fk_groups_user_roles	FOREIGN	KEY	(group_id)	REFERENCES	public.groups(id)	ON	DELETE	CASCADE;


--
--	Name:	notesheet_subgroup	fk_notesheet_subgroup_notesheet;	Type:	FK	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.notesheet_subgroup
	ADD	CONSTRAINT	fk_notesheet_subgroup_notesheet	FOREIGN	KEY	(notesheet_id)	REFERENCES	public.notesheets(id)	ON	DELETE	CASCADE;


--
--	Name:	notesheet_subgroup	fk_notesheet_subgroup_subgroup;	Type:	FK	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.notesheet_subgroup
	ADD	CONSTRAINT	fk_notesheet_subgroup_subgroup	FOREIGN	KEY	(subgroup_id)	REFERENCES	public.subgroups(id)	ON	DELETE	CASCADE;


--
--	Name:	subgroup_user	fk_subgroup_user_subgroup;	Type:	FK	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.subgroup_user
	ADD	CONSTRAINT	fk_subgroup_user_subgroup	FOREIGN	KEY	(subgroup_id)	REFERENCES	public.subgroups(id)	ON	DELETE	CASCADE;


--
--	Name:	subgroup_user	fk_subgroup_user_user;	Type:	FK	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.subgroup_user
	ADD	CONSTRAINT	fk_subgroup_user_user	FOREIGN	KEY	(user_id)	REFERENCES	public.users(id)	ON	DELETE	CASCADE;


--
--	Name:	track_event	fk_track_event_event;	Type:	FK	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.track_event
	ADD	CONSTRAINT	fk_track_event_event	FOREIGN	KEY	(event_id)	REFERENCES	public.events(id)	ON	DELETE	CASCADE;


--
--	Name:	track_event	fk_track_event_track;	Type:	FK	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.track_event
	ADD	CONSTRAINT	fk_track_event_track	FOREIGN	KEY	(track_id)	REFERENCES	public.tracks(id)	ON	DELETE	CASCADE;


--
--	Name:	notesheets	fk_tracks_notesheets;	Type:	FK	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.notesheets
	ADD	CONSTRAINT	fk_tracks_notesheets	FOREIGN	KEY	(track_id)	REFERENCES	public.tracks(id)	ON	DELETE	CASCADE;


--
--	Name:	user_group	fk_user_group_group;	Type:	FK	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.user_group
	ADD	CONSTRAINT	fk_user_group_group	FOREIGN	KEY	(group_id)	REFERENCES	public.groups(id)	ON	DELETE	CASCADE;


--
--	Name:	user_group	fk_user_group_user;	Type:	FK	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.user_group
	ADD	CONSTRAINT	fk_user_group_user	FOREIGN	KEY	(user_id)	REFERENCES	public.users(id)	ON	DELETE	CASCADE;


--
--	Name:	announcements	fk_users_announcements;	Type:	FK	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.announcements
	ADD	CONSTRAINT	fk_users_announcements	FOREIGN	KEY	(sender_id)	REFERENCES	public.users(id)	ON	DELETE	SET	NULL;


--
--	Name:	user_group_roles	fk_users_group_roles;	Type:	FK	CONSTRAINT;	Schema:	public;	Owner:	band_manager_user
--

ALTER	TABLE	ONLY	public.user_group_roles
	ADD	CONSTRAINT	fk_users_group_roles	FOREIGN	KEY	(user_id)	REFERENCES	public.users(id)	ON	DELETE	CASCADE;


--
--	PostgreSQL	database	dump	complete
--