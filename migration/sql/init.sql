SET client_min_messages = warning;
SET client_encoding = 'UTF8';


CREATE TABLE message (
       id bigserial PRIMARY KEY,
       resource character varying NOT NULL,
       topic character varying NOT NULL,

       external_id character varying NOT NULL,
       answered_to character varying NOT NULL DEFAULT '',
       posted_time timestamp WITH time zone NOT NULL,

       author character varying NOT NULL,
       tree_path character varying NOT NULL,

       creation_time timestamp WITH time zone NOT NULL DEFAULT current_timestamp,
       UNIQUE(resource, topic, external_id)
);
