<h1 align="center">Mark URL</h1>
<p align="center">A url shortner written in Go</p>
<p align="center">
  <a href="#-introduction">Introduction</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-project">Project</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-how-to-run">How to Run</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-license">License</a>
</p>


## ⚙️ Techs 
- [Go](https://go.dev/)
- [Postgres](https://www.postgresql.org/)
- [Docker](https://www.docker.com/)
- [Gin](https://github.com/gin-gonic/gin)
- [Validator](https://pkg.go.dev/gopkg.in/validator.v2)
- [Gorm](https://github.com/go-gorm/gorm)

---

## 💻 Project

This project is focused on designing a system for a URL shortener. My approach was to begin with a basic application, and then enhance it by using Docker containers. This will allow for easy deployment of the application on various Platform as a Service (PaaS) providers such as Heroku and AWS Beanstalk.

The above quotation is only valid when our application has few users. For scalability, we need to deploy it on a service like AWS EC2.

The URL shortener takes a long URL and generates a unique, shorter URL. When a request is made to this shortened URL, the system redirects the user to the original, longer URL.the short url the system redirect to the origanal url.

For example:
```
https://www.linkedin.com/in/rafael-mgr/
```

And we receive the following
```
https://tinyurl.com/bdfwpx29
```

URL shortening is used to optimize links between devices, track individual links to analyze the audience, measure the performance of advertising campaigns, and hide original affiliate URLs. All of this is done to improve user experience.

You can check the system desing where a take the idea []()

## Requirements
1. Given a URL, our system must generate a short and unique "alias" of the name. This link should be short enough to be easily copied and placed between applications.
2. When users access the short link, our service must redirect them to the original link.
3. Users must have the option to get a customized link for their URL
4. Links must expire after a standard interval. Users should be able to specify the expiration time.

---
## 🚀 How to run

To run the this project 

- Change the env variables
- And with docker you can use `docker-compose up -d --build`

> This project using an Postgres as default database.


---
## 📄 License
The projects is under the MIT license. See the file [LICENSE](LICENSE) fore more details

---

Made with ♥ by Rafael and [Rocketeseat](https://youtu.be/w_el04y0cHo])👋 


[![Linkedin Badge](https://img.shields.io/badge/-Rafael-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https://www.linkedin.com/in/tgmarinho/)](https://www.linkedin.com/in/rafael-mgr/)
[![Gmail Badge](https://img.shields.io/badge/-Gmail-red?style=flat-square&link=mailto:nelsonsantosaraujo@hotmail.com)](mailto:ribeirorafaelmatehus@gmail.com)
