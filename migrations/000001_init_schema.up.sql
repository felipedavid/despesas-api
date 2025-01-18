CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    phone_number TEXT,
    birth_date DATE,
    job_title TEXT,
    company_name TEXT,
    document TEXT,
    document_type TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE address (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users (id) NOT NULL,
    full_address TEXT NOT NULL,
    state TEXT NOT NULL,
    primary_address TEXT NOT NULL,
    country TEXT NOT NULL,
    postal_code TEXT NOT NULL,
    city TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE phone_number (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users (id) NOT NULL,
    value TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE email (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users (id) NOT NULL,
    value TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE plan (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    price INT NOT NULL,
    duration_days INT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW (),
    updated_at TIMESTAMPTZ DEFAULT NOW ()
);

CREATE TABLE subscription (
    id SERIAL PRIMARY KEY,
    plan_id INT NOT NULL REFERENCES plan (id),
    user_id INT NOT NULL REFERENCES users (id),
    start_date TIMESTAMPTZ NOT NULL,
    end_date TIMESTAMPTZ NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE connector (
    id SERIAL PRIMARY KEY,
    external_id INTEGER NOT NULL, -- pluggy connector id
    name TEXT NOT NULL,
    primary_color TEXT NOT NULL,
    country TEXT NOT NULL,
    type TEXT NOT NULL,
    status TEXT NOT NULL,
    stage TEXT NOT NULL,
    is_open_finance BOOL NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE fi_connection (
    id SERIAL PRIMARY KEY,
    external_id TEXT NOT NULL, -- pluggy item id
    connector_id INTEGER REFERENCES connector (id) NOT NULL,
    status TEXT NOT NULL,
    execution_status TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE bank_account_data (
    id SERIAL PRIMARY KEY,
    transfer_number TEXT NOT NULL,
    closing_balance INTEGER NOT NULL,
    automatically_invested_balance INTEGER,
    overdraft_contracted_limit INTEGER,
    overdraft_used_limit INTEGER,
    unarranged_overdraft_amount INTEGER,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE credit_account_data (
    id SERIAL PRIMARY KEY,
    level TEXT NOT NULL,
    brand TEXT NOT NULL,
    balance_close_date DATE NOT NULL,
    available_credit_limit INTEGER NOT NULL,
    balance_foreign_currency INTEGER NOT NULL,
    minimum_payment INTEGER NOT NULL,
    credit_limit INTEGER NOT NULL,
    is_limit_flexible INTEGER NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE account (
    id SERIAL PRIMARY KEY,
    external_id TEXT, -- pluggy account id
    type TEXT NOT NULL,
    subtype TEXT NOT NULL,
    name TEXT NOT NULL,
    balance INT NOT NULL,
    currency_code TEXT NOT NULL,
    number TEXT,
    owner TEXT,
    tax_number TEXT,
    bank_account_data_id INTEGER REFERENCES bank_account_data (id),
    credit_account_data_id INTEGER REFERENCES credit_account_data (id),
    fi_connection_id INTEGER REFERENCES fi_connection (id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE category (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    user_id INTEGER REFERENCES users (id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE transaction (
    id SERIAL PRIMARY KEY,
    external_id TEXT, -- pluggy transaction id
    user_id INTEGER REFERENCES users (id) NOT NULL,
    account_id INTEGER REFERENCES account (id) NOT NULL,
    description TEXT NOT NULL,
    amount INT NOT NULL,
    currency_code TEXT NOT NULL,
    transaction_date TIMESTAMPTZ NOT NULL,
    category_id INTEGER REFERENCES category (id),
    status TEXT,
    type TEXT,
    operation_type TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMPTZ
);
