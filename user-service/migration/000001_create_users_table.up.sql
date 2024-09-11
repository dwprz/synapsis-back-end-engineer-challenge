    CREATE TYPE "UserRole" AS ENUM( 'USER', 'ADMIN');

    CREATE TABLE IF NOT EXISTS users (
        user_id VARCHAR(21) PRIMARY KEY,
        email VARCHAR(100) NOT NULL UNIQUE,
        full_name VARCHAR(100) NOT NULL,
        whatsapp VARCHAR(20) NULL,
        role "UserRole" NOT NULL DEFAULT 'USER',
        password VARCHAR(100) NOT NULL,
        refresh_token VARCHAR(500) NULL UNIQUE,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

    CREATE INDEX IF NOT EXISTS email_hash_index ON users USING HASH (email);
    
    CREATE INDEX IF NOT EXISTS refresh_token_hash_index ON users USING HASH (refresh_token);