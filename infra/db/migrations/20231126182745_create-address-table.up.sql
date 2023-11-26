CREATE TABLE
    IF NOT EXISTS addresses
    (
        id VARCHAR(191) PRIMARY KEY,
        city VARCHAR(191) NOT NULL,
        district VARCHAR(191),
        streetAndNumber VARCHAR(191),
        clientId VARCHAR(191),
        CONSTRAINT fk_client_address FOREIGN KEY(clientId) REFERENCES clients(id) ON DELETE CASCADE
    );