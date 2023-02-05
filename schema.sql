CREATE TABLE products
(
    id TEXT,
    name TEXT NOT NULL,
    image TEXT,
    details TEXT,
    sizes TEXT [],
    colours TEXT [],
    video_url TEXT,
    availability BOOLEAN,
    star INT [],
    labels TEXT [],
    discount NUMERIC(10,2) DEFAULT 0.00,
    price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    brands TEXT [],
    categories JSONB,
    created_date TIMESTAMP,
    updated_date TIMESTAMP,
    CONSTRAINT products_pkey PRIMARY KEY (id)
);

CREATE TABLE Articles
(
    id TEXT,
    title TEXT NOT NULL,
    image TEXT,
    article TEXT,
    created_date TIMESTAMP,
    updated_date TIMESTAMP,
    categories JSONB,
    CONSTRAINT articles_pkey PRIMARY KEY (id)
);

CREATE TABLE Vendors
(
    id TEXT,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    phone TEXT,
    address TEXT,
    logo_url TEXT,
    created_date TIMESTAMP,
    updated_date TIMESTAMP,
    CONSTRAINT vendors_pkey PRIMARY KEY (id)
);

CREATE TABLE Categories
(
    id TEXT,
    main_category TEXT NOT NULL,
    sub_category TEXT,
    CONSTRAINT category_pkey PRIMARY KEY (id)
);


CREATE TABLE Article_Categories
(
    id TEXT,
    article_category TEXT,
    CONSTRAINT article_category_pkey PRIMARY KEY (id)
);