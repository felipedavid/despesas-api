
CREATE TABLE users (
    id UUID DEFAULT (gen_random_uuid()) PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password BYTEA NOT NULL,
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

CREATE TABLE token (
    hash BYTEA PRIMARY KEY,
    user_id UUID REFERENCES users (id) NOT NULL,
    expiry TIMESTAMPTZ NOT NULL,
    scope TEXT NOT NULL
);

CREATE TABLE address (
    id UUID DEFAULT (gen_random_uuid()) PRIMARY KEY,
    user_id UUID REFERENCES users (id) NOT NULL,
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
    id UUID DEFAULT (gen_random_uuid()) PRIMARY KEY,
    user_id UUID REFERENCES users (id) NOT NULL,
    value TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE email (
    id UUID DEFAULT (gen_random_uuid()) PRIMARY KEY,
    user_id UUID REFERENCES users (id) NOT NULL,
    value TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE plan (
    id UUID DEFAULT (gen_random_uuid()) PRIMARY KEY,
    name TEXT NOT NULL,
    price INT NOT NULL,
    duration_days INT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW (),
    updated_at TIMESTAMPTZ DEFAULT NOW ()
);

CREATE TABLE subscription (
    id UUID DEFAULT (gen_random_uuid()) PRIMARY KEY,
    plan_id UUID NOT NULL REFERENCES plan (id),
    user_id UUID NOT NULL REFERENCES users (id),
    start_date TIMESTAMPTZ NOT NULL,
    end_date TIMESTAMPTZ NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE connector (
    id UUID DEFAULT (gen_random_uuid()) PRIMARY KEY,
    external_id UUID NOT NULL, -- pluggy connector id
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
    id UUID DEFAULT (gen_random_uuid()) PRIMARY KEY,
    external_id UUID NOT NULL, -- pluggy item id
    connector_id UUID REFERENCES connector (id) NOT NULL,
    status TEXT NOT NULL,
    execution_status TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE bank_account_data (
    id UUID DEFAULT (gen_random_uuid()) PRIMARY KEY,
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
    id UUID DEFAULT (gen_random_uuid()) PRIMARY KEY,
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
    id UUID DEFAULT (gen_random_uuid()) PRIMARY KEY,
    type TEXT NOT NULL,
    name TEXT NOT NULL,
    balance INT NOT NULL,
    currency_code TEXT NOT NULL,
    user_id UUID REFERENCES users (id) NOT NULL,
    external_id UUID, -- pluggy account id
    subtype TEXT,
    number TEXT,
    owner TEXT,
    tax_number TEXT,
    bank_account_data_id UUID REFERENCES bank_account_data (id),
    credit_account_data_id UUID REFERENCES credit_account_data (id),
    fi_connection_id UUID REFERENCES fi_connection (id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE category (
    id UUID DEFAULT (gen_random_uuid()) PRIMARY KEY,
    name TEXT NOT NULL,
    default_category BOOLEAN NOT NULL DEFAULT false,
    user_id UUID REFERENCES users (id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMPTZ
);

INSERT INTO category (name)
VALUES
    ('wage'),
    ('transportation'),
    ('food'),
    ('education'),
    ('entertainment');

CREATE TABLE transaction (
    id UUID DEFAULT (gen_random_uuid()) PRIMARY KEY,
    external_id UUID, -- pluggy transaction id
    user_id UUID REFERENCES users (id) NOT NULL,
    account_id UUID REFERENCES account (id),
    description TEXT,
    amount INT NOT NULL,
    currency_code TEXT NOT NULL,
    transaction_date TIMESTAMPTZ NOT NULL,
    category_id UUID REFERENCES category (id),
    status TEXT,
    type TEXT,
    operation_type TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMPTZ
);
