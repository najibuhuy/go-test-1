> go mod init
> go get github.com/gofiber/fiber/v2
> go get github.com/jackc/pgx/v5
> go get github.com/jackc/pgx/v5/pgxpool
> go get github.com/joho/godotenv
> go get go.uber.org/zap 

#How to run
1. Create a .env file in the root directory with your database credentials.
2. docker-compose -f deployment.yml up --build -d
3. Run query in bank_account.sql to create a table in database
4. You can test it in postman, import "test golang.postman_collection.json"