CREATE TABLE IF NOT EXISTS "files"
(
    "id"       bigint PRIMARY KEY,
    "name"     text NOT NULL ,
    "mimetype" text NOT NULL ,
    "data"     bytea NOT NULL
);

CREATE TABLE IF NOT EXISTS "assets"
(
    "id"         bigint NOT NULL UNIQUE,
    "name"       varchar(64) NOT NULL,
    "url_path"   text PRIMARY KEY,
    "file_id"    bigint REFERENCES "files" ("id") ON DELETE CASCADE NOT NULL,
    "creator_id" bigint NOT NULL
);

CREATE SEQUENCE IF NOT EXISTS global_snowflake_id_seq;

CREATE OR REPLACE FUNCTION create_snowflake()
    RETURNS bigint
    LANGUAGE 'plpgsql'
AS
$BODY$
DECLARE
    our_epoch  bigint := 1577836800;
    seq_id     bigint;
    now_millis bigint;
    -- the id of this DB shard, must be set for each
    -- schema shard you have - you could pass this as a parameter too
    shard_id   int    := 1;
    result     bigint := 0;
BEGIN
    SELECT nextval('global_snowflake_id_seq') % 1024 INTO seq_id;

    SELECT FLOOR(EXTRACT(EPOCH FROM clock_timestamp()) * 1000) INTO now_millis;
    result := (now_millis - our_epoch) << 22;
    result := result | (shard_id << 9);
    result := result | (seq_id);
    return result;
END;
$BODY$;
