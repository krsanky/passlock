-- change to suit db type or context (this is for a webapp and postgres)

-- drop table passlock;
CREATE TABLE passlock (                                                                     
    id SERIAL PRIMARY KEY,
    account_id INT REFERENCES account ON DELETE CASCADE NOT NULL,

    title text NOT NULL,
    password text NOT NULL,
    ts TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    UNIQUE (account_id, title)
);

