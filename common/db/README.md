# **POSTGRESQL**

You have to make sure that you are in the **docs directory inside the project**.

`cd /Golang-Curmin/common/db`

**init-db.sql** file is used to initialize the database.

---

#### **Create Postgresql Image**

`docker build -t {YOUR-DESIRED-IMAGE-NAME} .`

**Example:**

`docker build -t postgresql-img:latest .`

---

#### **Create Postgresql Container**

`docker run -d --name curmin  -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD={YOUR-DESIRED-PASSWORD} -e POSTGRES_DB=postgres -p 5432:5432 {YOUR-DESIRED-IMAGE-NAME}:latest`

**Example:**

`docker run -d --name curmin  -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=passw -e POSTGRES_DB=postgres -p 5432:5432 postgresql-img:latest`

**Note:** You can give volume path as projects db folder.

---