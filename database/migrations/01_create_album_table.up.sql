USE recordings;

CREATE TABLE IF NOT EXISTS albums (
    id      INT AUTO_INCREMENT NOT NULL,
    title   VARCHAR(128) NOT NULL,
    artist  VARCHAR(255) NOT NULL,
    price   DECIMAL(5,2) NOT NULL,
    PRIMARY KEY (`id`)
)