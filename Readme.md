# About

This is the back end to a website. The content of the website is stored in a Postgres database.

When someone requests a page a list of all pages is retrieved from the database and that list is compared against the requested page. If a match is found the page name from that came from the database is used moving forward. This was done to avoid SQL injection and to not trust the users input in the URL.

This backend also supports Golang's templates. 