# EmailSender


# Configuration

## Docker-Compose File
db:
    image: 

For mysql, copy the following after the image:  - mysql/mysql-server

For postgres, copy the following after the image: - postgres

For sq;-server, copy the following after the image: - mcr.microsoft.com/mssql/server

## .env

Create a .env file and put the following values in it with the information required

![env_file](https://i.imgur.com/u1cA0t1.png)    ![docker-compose](https://i.imgur.com/oQWkiSt.png)

Some of those values are used for the docker-compose.yml file