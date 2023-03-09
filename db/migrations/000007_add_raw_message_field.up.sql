ALTER TABLE messages
    ADD COLUMN raw_message BYTEA;
ALTER TABLE erc_to_native_messages
    ADD COLUMN raw_message BYTEA;

ALTER TABLE messages
    ALTER COLUMN raw_message TYPE BLOB;
ALTER TABLE erc_to_native_messages
    ALTER COLUMN raw_message TYPE BLOB;
