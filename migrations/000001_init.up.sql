CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    cellphone TEXT,
    tax_number TEXT,
    birth_date DATE,
    created_at TIMESTAMPZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMPZ
);

CREATE TABLE plan (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    price INT NOT NULL,
    duration_days INT NOT NULL,
    created_at TIMESTAMPZ DEFAULT NOW (),
    updated_at TIMESTAMPZ DEFAULT NOW ()
);

CREATE TABLE subscription (
    id SERIAL PRIMARY KEY,
    plan_id INT NOT NULL REFERENCES plan (id),
    user_id INT NOT NULL REFERENCES users (id),
    start_date TIMESTAMPZ NOT NULL,
    end_date TIMESTAMPZ NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMPZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE transaction (
    id SERIAL PRIMARY KEY,
    external_id TEXT, -- pluggy transaction id
    user_id INTEGER REFERENCES users (id) NOT NULL,
    account_id INTEGER REFERENCES account (id) NOT NULL,
    description TEXT NOT NULL,
    amount INT NOT NULL,
    currency_code TEXT NOT NULL,
    transaction_date TIMESTAMPZ NOT NULL,
    category_id INTEGER REFERENCES category (id),
    created_at TIMESTAMPZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMPZ
);

CREATE TABLE account (
    id SERIAL PRIMARY KEY,
    external_id TEXT, -- pluggy account id
    type ACCOUNT_TYPE NOT NULL,
    subtype ACCOUNT_SUBTYPE,
    name TEXT NOT NULL,
    balance INT NOT NULL,
    currency_code TEXT NOT NULL,
    number TEXT,
    owner TEXT,
    tax_number TEXT,
    bank_data_id INTEGER REFERENCES bank_data (id),
    credit_data_id INTEGER REFERENCES credit_data (id),
    fi_connection_id INTEGER REFERENCES fi_connection (id),
    created_at TIMESTAMPZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMPZ
);

CREATE TABLE bank_account_data (
    id SERIAL PRIMARY KEY,
    transfer_number TEXT NOT NULL,
    closing_balance INTEGER NOT NULL,
    automatically_invested_balance INTEGER,
    overdraft_contracted_limit INTEGER,
    overdraft_used_limit INTEGER,
    unarranged_overdraft_amount INTEGER,
    created_at TIMESTAMPZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPZ DEFAULT CURRENT_TIMESTAMP NOT NULL
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
    created_at TIMESTAMPZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE fi_connection (
    id SERIAL PRIMARY KEY,
    external_id TEXT NOT NULL, -- pluggy item id
    connector_id INTEGER REFERENCES connector (id) NOT NULL,
    status TEXT NOT NULL,
    execution_status TEXT NOT NULL,
    created_at TIMESTAMPZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMPZ
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
    created_at TIMESTAMPZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
);

CREATE TABLE category (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    user_id INTEGER REFERENCES users (id),
    created_at TIMESTAMPZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMPZ
);
