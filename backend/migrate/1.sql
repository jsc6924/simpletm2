ALTER TABLE User ADD COLUMN token TEXT;

UPDATE User
SET token = (
    SELECT token
    FROM APIToken
    WHERE APIToken.user_id = User.id
    LIMIT 1
);

DROP TABLE APIToken;

ALTER TABLE User RENAME TO User_old;
CREATE TABLE IF NOT EXISTS "User" (
    id VARCHAR(255) NOT NULL,
    salt TEXT NOT NULL,
    token TEXT NOT NULL,  -- Make token NOT NULL
    PRIMARY KEY (id)
);
INSERT INTO User (id, salt, token)
SELECT id, salt, token
FROM User_old;
DROP TABLE User_old;


ALTER TABLE Permission RENAME TO Permission_old;

CREATE TABLE Permission
(
    user_id VARCHAR(255) NOT NULL,
    game_id TEXT NOT NULL,
    permission INT NOT NULL, -- (none=0, read=1, write=2, admin=3)
    PRIMARY KEY (user_id, game_id),
    FOREIGN KEY (user_id) REFERENCES User(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    FOREIGN KEY (game_id) REFERENCES Game(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

INSERT INTO Permission (user_id, game_id, permission)
SELECT user_id, game_id, permission
FROM Permission_old;

DROP TABLE Permission_old;

