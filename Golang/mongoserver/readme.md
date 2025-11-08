1. install mongodb in vm and some of the useful commands
    - mongosh
    - show dbs
    - use netflix
    - show collections
    - db.watchlist.find()

2. go run main.go will start server and connect with db
3. Use below rest calls using postman
    - GET - http://10.50.14.20:4000/api/movies
    - DELTE - http://10.50.14.20:4000/api/deleteallmovie
    - POST - http://10.50.14.20:4000/api/movie
        {
            "movie" : "fds",
            "watched": false
        }
