CREATE TABLE
    IF NOT EXISTS clients (
        id VARCHAR(191) PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        age INTEGER NOT NULL,
        phoneNumber VARCHAR(255) NOT NULL,
        budget VARCHAR(255),
        budgetDescription VARCHAR(255),
        anamnese VARCHAR(255)
    )