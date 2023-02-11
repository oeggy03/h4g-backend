# Buddy4Good

Buddy4Good is a platform for both persons with disabilities and able-bodied persons to connect with each other. Buddy4Good encourages the buddying up of an able bodied person (Best Buddy) with a disabled person (Special Friend) to spend time doing activities together. These may be workshops, exercising, bird-watching, or any interest either party may have.


## Set-up for backend

1. Install Go and MySQL server
2. Clone `https://github.com/oeggy03/h4g-backend`
3. Run `go mod tidy` in the project directory
4. Create the database `h4g_db` locally
5. Change `password` in the .env file to the password for MySQL database


## Starting the web app
Note that the frontend at https://github.com/oeggy03/h4g-frontend has to be set up before starting the app.

1. Run `go run main.go` in the project directory for the backend
2. Run `npm start` in the project directory for the frontend


## Future Plans

- Host the database on Google Cloud SQL rather than MySQL for easier setup
- Badges and a point system for people who create or join many activities
