CREATE TABLE bank_accounts (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name TEXT NOT NULL,
    id_number TEXT UNIQUE NOT NULL,
    phone TEXT UNIQUE NOT NULL,
    bank_no TEXT UNIQUE NOT NULL,
    balance NUMERIC DEFAULT 0.0