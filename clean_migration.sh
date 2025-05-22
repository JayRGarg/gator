cd /Users/jayrgarg/projects/gator/sql/schema
goose postgres "postgres://jayrgarg:@localhost:5432/gator" down
goose postgres "postgres://jayrgarg:@localhost:5432/gator" up
cd /Users/jayrgarg/projects/gator
