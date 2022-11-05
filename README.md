# Golang-Curmin
*This repository contains the necessary codes &amp; info about REST API which is used to backend services for a minimal currecy application in Golang. It has some basic CRUD operations on user system. Using Postgres*

# About the project

*The functions of this service are as follows;*

1. List Latest Currency Data
    - API should be able to list latest currency data by base. (Target optional)
2. List Currency Names
    - API can list currency names and codes
3. List Currency by time
    - API can list users by time range
4. Insert Currency Names
    - API can insert currency names and codes
5. Insert Currency Data
    - API can insert only targeted data in request paylaod
6. Insert to All Currency Data
    - API can insert all data in request paylaod

---

### Prerequisites and Installation

- #### **Docker:** 
  You can install Docker Desktop by following the instructions on the [Docker Desktop website](https://desktop.docker.com/).

  If you are using **MacOS**, you can install Docker by following commands:
    - *`brew install docker`*
    - *`docker run hello-world`*
  
 ---

- #### **PostgreSQL:**

  Plese refer [**here**](https://github.com/denizedizcan/Golang-Curmin/tree/main/common/db/README.md) for the initialization scripts.

---

- #### **Go:**
  
  If you haven't done already, You need to install Go by following the instructions on the [Go website](https://golang.org/doc/install).

---
- #### **Project:**

    You have to clone the project from the [Github repository](https://github.com/denizedizcan/Golang-Curmin) and run the below script to build the project

   *`https://github.com/denizedizcan/Golang-Curmin.git`*

---
- #### **Run the Project:**

    You can run the project directly by running the below script from the same directory as the project:

    *`go run main.go`*