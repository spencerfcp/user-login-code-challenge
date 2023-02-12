# Thank You

Thank you for taking the time out to review this. If you have any questionns, don't hestiate to let me know. Looking forward to chatting with you!

Application has also been deployed to EC2, available here: http://ec2-52-3-241-109.compute-1.amazonaws.com/

# Getting Started

This project utilizes Next.js for the Frontend, Go for the backend, and Docker/Postgres for the database.

To get started, need to set up the enviornment which can be done through running:

- `./setup-dev-machine.sh`

#### The following links direnv up to whatever shell you're using.
- `touch ~/.bash_profile`
- `echo "eval \"\$(direnv hook bash)\"" >> ~/.bash_profile`
- `echo "eval \"\$(direnv hook zsh)\"" >> ~/.zshrc`

To start database, in `database`:

- make sure docker is running
- run `docker build -t scoir_db:v1 .` - run `docker compose up -d`

To start backend, in `backend`:

- run `direnv allow`
- run `go run main.go`

To start frontend, in `frontend`:

- run `pnpm install`
- run `pnpm dev`

You should now be able to access and use the application at http://localhost:3000

While you can create a new user, a test user has already been created for you: 

Username: test
Password: pass

# Assumptions about this challenge.

- Development will be done on a mac
- No SSL certificate is necessary
- Text-based content doesn't need to be translated.
- No session management is required, only informing the user of a successful login.
- Should have test coverage
- API should be built with future requests in mind.
- Storing .env in github is alright for sake of example.
