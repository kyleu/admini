-- {% func ExampleDatabase() %}
create extension if not exists hstore;

drop table if exists simple;

create table simple (
  id serial primary key,
  t text, v varchar(128), cf char(128), bp bpchar, ss smallserial, bs bigserial,
  b bool, c char, i2 int2, i4 int4, o oid, i8 int8, ib numeric(100, 100), m money, f4 float4, f8 float8,
  d date, tm time, tmz time with time zone, ts timestamp, tz timestamp with time zone, iv interval,
  u uuid, j json, x xml,
  bf bit(4), bt varbit, cd cidr, it inet, ma macaddr,
  p point, pg polygon, bx box, cr circle, ph path,
  tq tsquery, tv tsvector
);

insert into simple (
  t, v, cf, bp, ss, bs,
  b, c, i2, i4, o, i8, ib, m, f4, f8,
  d, tm, tmz, ts, tz, iv,
  u, j, x, bf, bt, cd, it, ma,
  p, pg, bx, cr, ph,
  tq, tv
) values (
  'text', 'varchar', 'chars', 'padded', 0, 0,
  true, 'x', 2, 4, 0, 8, 0, 19.99, 0.4, 0.8,
  '2000-01-01', '04:00:00', '06:00:00', '2000-01-01 00:00:01', '2000-02-02 00:00:02', '1 hour',
  '00000000-0000-0000-0000-000000000000', '{}'::json, '<a><b/></a>', '0101', '01010101', '192.168/24', '192.168.0.1', '08:00:2b:01:02:03',
  '(1, 1)', '((0, 0), (1, 1), (2, 2))', '((0, 0), (1, 1))', '((0, 0), 1)', '((0, 0), (1, 1), (2, 2))',
  'foo & bar', 'foo bar baz'
);

drop table if exists complex;

create table complex (
  id serial primary key,
  ay bytea,
  ab bool[],
  ao oid[],
  af float8[],
  aj json[],
  at text[],
  ats timestamp[],
  h hstore
);

insert into complex (
  ay, ab, ao, af, aj, at, ats, h
) values (
  '{ 0, 1, 2 }', '{true, false}', '{0, 1}', '{0.1, 0.1}', '{ "{}", "{}" }', '{ "a", "b", "c" }', '{ 2000-01-01, 2000-01-02 }', 'a => x, b => y, c => z'
);

create table pktest (
  id integer,
  foo varchar(100),
  bar uuid unique,
  primary key (id, foo)
);

insert into pktest (
  id, foo, bar
) values (
  0, 'a', '00000000-0000-0000-0000-000000000000'
);

create view testview as select * from simple;

-- {% endfunc %}
