-- Create new game 'whisqy' if it does not exist
INSERT INTO games(name)
SELECT 'whisqy'
WHERE NOT EXISTS (
  SELECT 1 FROM games WHERE name = 'whisqy'
);

