CREATE TABLE IF NOT EXISTS bill(
  id serial PRIMARY Key,
  name VARCHAR(50),
  amount int,
  category VARCHAR(50),
  importance VARCHAR(50),
  frequency VARCHAR(50),
  dueDate int
);

INSERT INTO bill (name, amount, category, importance, frequency, dueDate)
VALUES
    ('Mortgage', 3245, 'living expense', 'high', 'monthly', 1),
    ('Goddard Tuition', 2100, 'daycare', 'high', 'monthly', 1),
    ('Groceries', 200, 'living expense', 'high', 'weekly', 1),
    ('Tahoe', 700, 'living expenses', 'high', 'monthly', 1),
    ('TECO', 300, 'living expense', 'high', 'monthly', 20),
    ('Water', 80, 'living expense', 'high', 'monthly', 15),
    ('Frontier', 70, 'living expense', 'high', 'monthly', 18),
    ('Car Insurance', 540, 'living expense', 'high', 'monthly', 10),
    ('Life Insurance', 95, 'living expense', 'high', 'monthly', 19)
;