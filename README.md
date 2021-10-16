# STEELMAN
# DB
STEELMAN database is managed with [migrate](https://github.com/golang-migrate/migrate).
To create an update in the database like an addition, removal or edition of a table or a column in a table you must create a _migration_ file running the following command:
```bash
$migrate create -ext sql -dir db/migrations -seq <migration>
```
After running the command, two files in `db/migrations` will be created, the _up_ file with the extension _up.sql_ and the _down_ file with the extension _down.sql_. In the _up_ file you must write the sql code to udpate the database and in the _down_ file you must write the sql code to rollback the intended update.

Docker-compose will take care of the rest, you just have to run the `migrate` container to update the database. 

Note: `<migration>` must have a representative name like _create_user_table_ or _remove_id_column_user_table_

For more info about how to use _migrate_: https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md
