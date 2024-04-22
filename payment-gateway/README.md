# Payment service 
## Overview
This respository was created by Mijael Balderrama for a payment service, this service send a comunication to another bank service 

## how to do it.

1. `install docker`
2. `go to payment-gateway folder`
3. `execute the comand:`

        docker compose up
    this command will compile the application and rise up the database
4. in the project folder go to `db/migration/db-up.sql` file and copy the content
5. open any postgressql manager `I recomend dbeaver-ce`

5. connet to database, all credentials are in `env` file
6. open a sql-script and paste the sql script and execute
