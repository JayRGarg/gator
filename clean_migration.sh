cd /Users/jayrgarg/projects/gator/sql/schema
goose postgres "postgres://jayrgarg:@localhost:5432/gator" down-to 0
goose postgres "postgres://jayrgarg:@localhost:5432/gator" up-to 5
cd /Users/jayrgarg/projects/gator
