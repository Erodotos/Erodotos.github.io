Blak Post

This repository contains the project files of this website. To run this website on a new environment (ie WAMP) you will need to:

 * Clone this repository
 * Run the compiler script [`compile.sh`] to compile required core files
 * Import the database from the repository [`SQL/databasename.sql`] into your local MySQL
 * Edit database configuration in [`mysite/_configDB.php`] and update to match your MySQL credentials
 * Run a build [`dev/build?flush=all`]