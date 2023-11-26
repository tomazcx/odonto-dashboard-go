CREATE TABLE
    IF NOT EXISTS appointments (
        id VARCHAR(191) PRIMARY KEY,
        date VARCHAR(191) NOT NULL,
        teeth VARCHAR(191) NOT NULL,
        proccedure VARCHAR(191) NOT NULL,
        clientId VARCHAR(191) NOT NULL,
        CONSTRAINT fk_client_appointments FOREIGN KEY(id) REFERENCES clients(id) ON DELETE CASCADE
    );