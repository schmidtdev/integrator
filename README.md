
<h1 align="center" style="font-weight: bold;">Integrator</h1>

<p align="center">
<a href="#tech">Technologies</a>
<a href="#started">Getting Started</a>
<!-- <a href="#routes">API Endpoints</a> -->
</p>


<p align="center">A simple integration manager for systems data</p>

<h2 id="technologies">💻 Technologies</h2>

- Golang
- PostgreSQL
- GORM

<h2 id="started">🚀 Getting started</h2>

Configure your DB instance and your .env, then ```go mod tidy``` and run!

<h3>Prerequisites</h3>

Here you list all prerequisites necessary for running your project. For example:

- [Golang](https://go.dev/)
- [GORM](https://gorm.io/docs/)
- [PostgreSQL](https://www.postgresql.org/)

<h3>Cloning</h3>

How to clone your project

```bash
git clone github.com/schmidtdev/integrator
```

<h3>Config .env variables</h2>

Use the `.env.example` as reference to create your configuration file `.env` with your PostgreSQL Credentials

```yaml
DB_HOST=localhost
DB_PORT=5432
DB_NAME=integrator_db
DB_USER=your_user
DB_PASSWORD=your_pwd
```

<h3>Starting</h3>

How to start the project

```bash
cd integrator
go mod tidy
go run . (or 'air' if using hot reload)
```

<!-- <h2 id="routes">📍 API Endpoints</h2>

Here you can list the main routes of your API, and what are their expected request bodies.
​
| route               | description                                          
|----------------------|-----------------------------------------------------
| <kbd>GET /authenticate</kbd>     | retrieves user info see [response details](#get-auth-detail)
| <kbd>POST /authenticate</kbd>     | authenticate user into the api see [request details](#post-auth-detail)

<h3 id="get-auth-detail">GET /authenticate</h3>

**RESPONSE**
```json
{
  "name": "Fernanda Kipper",
  "age": 20,
  "email": "her-email@gmail.com"
}
```

<h3 id="post-auth-detail">POST /authenticate</h3>

**REQUEST**
```json
{
  "username": "fernandakipper",
  "password": "4444444"
}
```

**RESPONSE**
```json
{
  "token": "OwoMRHsaQwyAgVoc3OXmL1JhMVUYXGGBbCTK0GBgiYitwQwjf0gVoBmkbuyy0pSi"
}
``` -->