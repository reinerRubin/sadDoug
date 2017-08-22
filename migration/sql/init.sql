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


-- SELECT gs, COUNT(m.id)
-- FROM generate_series('2017-08-10 00:00:00+03', '2017-08-19 00:00:00+03', interval '3 hour') AS gs
-- LEFT JOIN message AS m ON m.posted_time - gs < '3 hour' AND m.posted_time - gs > '0 hour'
-- GROUP BY gs
-- ORDER BY gs
