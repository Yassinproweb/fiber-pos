CREATE TABLE IF NOT EXISTS orders (
  id integer primary key autoincrement,
  name text not null,
  type text not null,
  status text not null,
  items integer not null,
  cost integer not null,
  custname text not null,
  custnumber text not null,
  destination text not null,
  datetime
);
