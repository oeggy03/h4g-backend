# Buddy4Good

## Problem Statement 

*"How might we, as a community empower, equip and co-create with persons with disabilities to enable them to access, use and enjoy places, services, produces and/or information, whether physical or virtual, so that persons with disabilities can connect to and be included in the wider community?"*

This is the problem statement from **SGEnable**, under **pillar 1**

From our research, we find that people with disabilities have more trouble getting around and mingling with others, compared to most of us. In addition, not all of these people have friends or family who are willing and able to bring them out for activities, or help them enjoy the many wonders of the world like the way we do.

However, there is no lack of helpful, kind-hearted people in Singapore. Yet, there is a lack of platforms aimed at connecting both groups. Hence, we have created Buddy4Good, which aims to encourage both disabled and able-bodied people to go out and enjoy activities together.

**Buddy4Good** is a platform for both persons with disabilities and able-bodied persons to connect with each other. Buddy4Good encourages the buddying up of able-bodied people (Best Buddy) with disabled people (Special Friend) to spend time doing activities together. These may be workshops, exercising, bird-watching, or any interest either party may have.

## About the backend
The backend was built with Go, and uses jwt tokens for authentication and security. The password is encrypted with bcrypt before being stored in the database to improve the backend security.

For many of the functions, the backend checks that the user is signed in / the correct user before sending back any information.

## Our presentation slides
https://bit.ly/h4gbuddy4g

## Set-up instructions
**Notes:**

- All of this was set up on a Windows PC. If you are using another OS, I apologise but I will not be able to give specific instructions for it.
- We have the frontend and backend in separate repos.
- Frontend: https://github.com/oeggy03/h4g-frontend Backend(this repo): https://github.com/oeggy03/h4g-backend

## Set-up for backend

1. Install Go, MySQL server and MySQL workbench
2. Clone this backend repo
3. Using the terminal, run ```go mod tidy``` in the project directory
4. Create the database "h4g_db" locally using MySQL. (Run CREATE DATABASE h4g_db)
5. Change the username and password in the .env file to your own for the MySQL server
6. (Optional) In the "SQLfile" folder, there are 4 .sql files. You can choose to run them in MySQL workbench to get a pre-made set of users, comments and activities. Otherwise, you will have to create accounts, posts and comments on your own. 
8. Using the terminal, run ```go run main.go```. You may get a popup from Windows Defender, click agree. (Note: if this step fails, it may be because of your antivirus / firewall)

**If you chose to use the pre-made SQL files**
An "best buddy" account you could use:
username: oeggy03
password: hello

An "special friend" account you could use:
username: oeggy034
password: hello

## Starting the web app
Note that the frontend at https://github.com/oeggy03/h4g-frontend has to be set up before starting the app.

1. Run `go run main.go` in the project directory for the backend
2. Run `npm start` in the project directory for the frontend


## Future Plans

- We were unable to deploy the product online due to the limited time we had and the costs involved.
- We plan to host the MySQL database on Google Cloud SQL rather than MySQL server for easier setup.
- Feature wise, we plan to implement badges and a point system for people who create or join many activities, in order to encourage use of the platform
- We could also add a feature, which displays all handicapped-friendly areas in Singapore. However, we will need an api for that.
- Set up a backend middleware to improve backend authentication and security.
