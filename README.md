# dbmagritte

_created by Austin Poor_

A database migration tool written in Go.

## Usage

Commands:
* `init`: Set up the repo by creating a `.dbmagritte.yaml` file and a `migrations/` directory.
* `new`: Create a new migration.
* `info`: Get the current state of the DB.
* `validate`: Check if the DB is in a valid state? Optional `validate.sql` file?

* `up`: Move forward in the migration tree.
* `down`: Move back in the migration tree.
* `reset`: Roll-back all migrations
* `fast-forward`: Move to the newest migration (assuming there's just one)

## Notes

`new` creates a new directory in the `migrations/` directory named after the current git hash.

The directory needs to have an `up.sql` (to make a migration) and a `down.sql` (to undo the migration). It can also, optionally, have a `validate.sql` file to check if the DB is in a valid state.

**Q:** Should there be a config file in each directory? That way there can be multiple `up`/`down`/`validate` SQL files per migration. They will be run in order. Might solve issues of validating with multiple SQL statements in one file?

Create global flags to set the path to the project root, so it can be run from somewhere else.
