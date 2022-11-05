# **POSTGRESQL**

You have to make sure that you are in the **docs directory inside the project**.

`cd /Golang-Curmin/common/db`

---

#### **Create Postgresql Image**

`docker build -t postgresql-img:latest .`

---

#### **Create Postgresql Container**

`docker run -d --name curmin  -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=passw -e POSTGRES_DB=postgres -p 5432:5432 postgresql-img:latest`

---