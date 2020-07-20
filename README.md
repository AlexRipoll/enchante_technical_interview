# Enchainte_technical_interview

## Resolution for enchainté's technical interview

### Application execution

#### First steps

1. [Install Docker](https://www.docker.com/get-started)
2. Clone this project: `git clone https://github.com/AlexRipoll/enchante_technical_interview.git`
3. Move to the project folder: `cd enchante_technical_interview`

#### Environment setup

1. create environment

    `make start`
       
2. update DB schema

    `docker exec -i $(docker-compose ps -q db) mysql -u root -proot enchainte_db < config/mysql/database.sql`
    
3. stop and delete environment

    `make down`
    
And that's it!

Host:

    http://localhost:9000/
    
Postman Doc:
    
    https://explore.postman.com/4DNpR32nDDWP3Y
    
    https://www.getpostman.com/collections/f37cccd389730c6cf513
    
