PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE urls (
    id INTEGER NOT NULL PRIMARY KEY,
    path TEXT NOT NULL,
    url TEXT NOT NULL
);
INSERT INTO urls VALUES(1,'/cicd','https://github.com/devvspaces/server_eyes');
INSERT INTO urls VALUES(2,'/bulk_mailer','https://github.com/devvspaces/bulk_emailer');
INSERT INTO urls VALUES(3,'/learn-py-decorators','https://thecodeway.hashnode.dev/learn-python-decorators-from-basic-to-pro-in-10-mins');
COMMIT;
