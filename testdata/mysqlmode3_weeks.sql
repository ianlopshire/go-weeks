CREATE TEMPORARY TABLE IF NOT EXISTS d0 AS select date '2020-12-31' - INTERVAL (a.a + (10 * b.a) + (100 * c.a) + (1000 * d.a) + (10000 *e.a)) DAY as dt
                                           from (select 0 as a union all select 1 union all select 2 union all select 3 union all select 4 union all select 5 union all select 6 union all select 7 union all select 8 union all select 9) as a
                                                    cross join (select 0 as a union all select 1 union all select 2 union all select 3 union all select 4 union all select 5 union all select 6 union all select 7 union all select 8 union all select 9) as b
                                                    cross join (select 0 as a union all select 1 union all select 2 union all select 3 union all select 4 union all select 5 union all select 6 union all select 7 union all select 8 union all select 9) as c
                                                    cross join (select 0 as a union all select 1 union all select 2 union all select 3 union all select 4 union all select 5 union all select 6 union all select 7 union all select 8 union all select 9) as d
                                                    cross join (select 0 as a union all select 1 union all select 2 union all select 3 union all select 4 union all select 5 union all select 6 union all select 7 union all select 8 union all select 9) as e
                                           WHERE date '2020-12-31' - INTERVAL (a.a + (10 * b.a) + (100 * c.a) + (1000 * d.a) + (10000 *e.a)) DAY between '1989-12-31' and '2020-12-31';


CREATE TEMPORARY TABLE IF NOT EXISTS d1 SELECT * FROM d0;

CREATE TEMPORARY TABLE IF NOT EXISTS Mode3Weeks1 AS
SELECT ROW_NUMBER() over w AS rownum,
       year,
       week,
       start
FROM (SELECT CAST(LEFT(YEARWEEK(d1.dt, 3), 4) AS UNSIGNED) as year,
             WEEK(d1.dt, 3) as week,
             d1.dt as start
      FROM d1
               INNER JOIN d0 ON d0.dt = (d1.dt - INTERVAL 1 DAY)
      where YEARWEEK(d0.dt, 3) != YEARWEEK(d1.dt, 3)
      ORDER BY d1.dt) as weeks
WINDOW w AS (ORDER BY start);

CREATE TEMPORARY TABLE IF NOT EXISTS Mode3Weeks2 AS SELECT * FROM Mode3Weeks1;

SELECT
    w1.year,
    w1.week,
    DATE_FORMAT(w1.start, '%Y-%m-%dT%T%Z') as start,
    DATE_FORMAT(w2.start,'%Y-%m-%dT%T%Z' ) as end
FROM Mode3Weeks1 w1
         INNER JOIN Mode3Weeks2 w2 on w1.rownum = (w2.rownum -1);