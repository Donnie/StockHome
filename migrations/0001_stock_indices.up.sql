PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS stocks (
  id INTEGER PRIMARY KEY,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  deleted_at DATETIME,

  description TEXT,
  name TEXT NOT NULL,
  sector TEXT,
  symbol VARCHAR(15) NOT NULL UNIQUE
);

CREATE INDEX IF NOT EXISTS idx_stocks_created_at ON stocks(created_at);
CREATE INDEX IF NOT EXISTS idx_stocks_updated_at ON stocks(updated_at);
CREATE INDEX IF NOT EXISTS idx_stocks_deleted_at ON stocks(deleted_at);

CREATE INDEX IF NOT EXISTS idx_stocks_name ON stocks(name);
CREATE INDEX IF NOT EXISTS idx_stocks_sector ON stocks(sector);

CREATE TABLE IF NOT EXISTS candles (
  id INTEGER PRIMARY KEY,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  deleted_at DATETIME,

  date DATETIME NOT NULL,
  close INTEGER NOT NULL,
  high INTEGER NOT NULL,
  low INTEGER NOT NULL,
  open INTEGER NOT NULL,
  volume BIGINT NOT NULL,
  stock_id INTEGER REFERENCES stocks (id) ON DELETE CASCADE,
  UNIQUE(date, stock_id)
);

CREATE INDEX IF NOT EXISTS idx_candles_created_at ON candles(created_at);
CREATE INDEX IF NOT EXISTS idx_candles_updated_at ON candles(updated_at);
CREATE INDEX IF NOT EXISTS idx_candles_deleted_at ON candles(deleted_at);

CREATE INDEX IF NOT EXISTS idx_candles_close ON candles(close);
CREATE INDEX IF NOT EXISTS idx_candles_date ON candles(date);
CREATE INDEX IF NOT EXISTS idx_candles_high ON candles(high);
CREATE INDEX IF NOT EXISTS idx_candles_low ON candles(low);
CREATE INDEX IF NOT EXISTS idx_candles_open ON candles(open);
CREATE INDEX IF NOT EXISTS idx_candles_stock_id ON candles(stock_id);
CREATE INDEX IF NOT EXISTS idx_candles_volume ON candles(volume);

CREATE TABLE IF NOT EXISTS indices (
  id INTEGER PRIMARY KEY,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  deleted_at DATETIME,

  description TEXT,
  name TEXT NOT NULL,
  symbol VARCHAR(15) NOT NULL UNIQUE
);

CREATE INDEX IF NOT EXISTS idx_indices_created_at ON indices(created_at);
CREATE INDEX IF NOT EXISTS idx_indices_updated_at ON indices(updated_at);
CREATE INDEX IF NOT EXISTS idx_indices_deleted_at ON indices(deleted_at);
CREATE INDEX IF NOT EXISTS idx_indices_name ON indices(name);

CREATE TABLE indices_stocks (
  index_id INTEGER REFERENCES indices (id) ON DELETE CASCADE,
  stock_id INTEGER REFERENCES stocks (id) ON DELETE CASCADE,
  PRIMARY KEY (index_id, stock_id)
);

CREATE INDEX IF NOT EXISTS idx_indices_stocks_index_id ON indices_stocks(index_id);
CREATE INDEX IF NOT EXISTS idx_indices_stocks_stock_id ON indices_stocks(stock_id);

