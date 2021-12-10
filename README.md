# Planning Poker

## Database
Ensure SQLite3 is installed and run the following to initialize the database:

```
sqlite3 planning-poker.db < db/create_games.sql
```

Seed the database:

```
sqlite3 planning-poker.db < db/seed_games_sd.sql
sqlite3 planning-poker.db < db/seed_games_whisqy.sql
```
