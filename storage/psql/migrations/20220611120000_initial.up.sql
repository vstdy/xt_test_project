CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- btc_usdt table
CREATE TABLE btc_usdt
(
    "id"        UUID DEFAULT uuid_generate_v4(),
    "buy"       INT         NOT NULL,
    "sell"      INT         NOT NULL,
    "timestamp" TIMESTAMPTZ NOT NULL,
    PRIMARY KEY ("id")
);

CREATE INDEX btc_usdt_timestamp_idx ON btc_usdt ("timestamp");

-- cur_rub table
CREATE TABLE cur_rub
(
    "id"   UUID DEFAULT uuid_generate_v4(),
    "date" DATE NOT NULL,
    "aud"  INT  NOT NULL,
    "azn"  INT  NOT NULL,
    "gbp"  INT  NOT NULL,
    "amd"  INT  NOT NULL,
    "byn"  INT  NOT NULL,
    "bgn"  INT  NOT NULL,
    "brl"  INT  NOT NULL,
    "huf"  INT  NOT NULL,
    "hkd"  INT  NOT NULL,
    "dkk"  INT  NOT NULL,
    "usd"  INT  NOT NULL,
    "eur"  INT  NOT NULL,
    "inr"  INT  NOT NULL,
    "kzt"  INT  NOT NULL,
    "cad"  INT  NOT NULL,
    "kgs"  INT  NOT NULL,
    "cny"  INT  NOT NULL,
    "mdl"  INT  NOT NULL,
    "nok"  INT  NOT NULL,
    "pln"  INT  NOT NULL,
    "ron"  INT  NOT NULL,
    "xdr"  INT  NOT NULL,
    "sgd"  INT  NOT NULL,
    "tjs"  INT  NOT NULL,
    "try"  INT  NOT NULL,
    "tmt"  INT  NOT NULL,
    "uzs"  INT  NOT NULL,
    "uah"  INT  NOT NULL,
    "czk"  INT  NOT NULL,
    "sek"  INT  NOT NULL,
    "chf"  INT  NOT NULL,
    "zar"  INT  NOT NULL,
    "krw"  INT  NOT NULL,
    "jpy"  INT  NOT NULL,
    PRIMARY KEY ("id"),
    UNIQUE ("date")
);

-- cur_btc table
CREATE TABLE cur_btc
(
    "id"        UUID DEFAULT uuid_generate_v4(),
    "timestamp" TIMESTAMPTZ NOT NULL,
    "rub"       INT         NOT NULL,
    "aud"       INT         NOT NULL,
    "azn"       INT         NOT NULL,
    "gbp"       INT         NOT NULL,
    "amd"       INT         NOT NULL,
    "byn"       INT         NOT NULL,
    "bgn"       INT         NOT NULL,
    "brl"       INT         NOT NULL,
    "huf"       INT         NOT NULL,
    "hkd"       INT         NOT NULL,
    "dkk"       INT         NOT NULL,
    "usd"       INT         NOT NULL,
    "eur"       INT         NOT NULL,
    "inr"       INT         NOT NULL,
    "kzt"       INT         NOT NULL,
    "cad"       INT         NOT NULL,
    "kgs"       INT         NOT NULL,
    "cny"       INT         NOT NULL,
    "mdl"       INT         NOT NULL,
    "nok"       INT         NOT NULL,
    "pln"       INT         NOT NULL,
    "ron"       INT         NOT NULL,
    "xdr"       INT         NOT NULL,
    "sgd"       INT         NOT NULL,
    "tjs"       INT         NOT NULL,
    "try"       INT         NOT NULL,
    "tmt"       INT         NOT NULL,
    "uzs"       INT         NOT NULL,
    "uah"       INT         NOT NULL,
    "czk"       INT         NOT NULL,
    "sek"       INT         NOT NULL,
    "chf"       INT         NOT NULL,
    "zar"       INT         NOT NULL,
    "krw"       INT         NOT NULL,
    "jpy"       INT         NOT NULL,
    PRIMARY KEY ("id")
);

CREATE INDEX cur_btc_timestamp_idx ON cur_btc ("timestamp");
