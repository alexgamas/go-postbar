curl -X GET -v http://localhost:8181/health

curl -X GET -v http://localhost:8181/posts


curl -X GET -v http://localhost:8181/posts/5bca15c977376163c170b5db

curl -X POST -v http://localhost:8181/posts \
  -H 'Content-Type: application/json' \
  -d '{
    "title" : "Go (linguagem de programação)",
    "message" : "Go é uma linguagem de programação criada pela Google e lançada em código livre em novembro de 2009. É uma linguagem compilada e focada em produtividade e programação concorrente,[4] baseada em trabalhos feitos no sistema operacional chamado Inferno.[5] O projeto inicial da linguagem foi feito em setembro de 2007 por Robert Griesemer, Rob Pike e Ken Thompson.[4] Atualmente, há implementações para Windows, Linux, Mac OS X e FreeBSD."
  }'

curl -X POST -v http://localhost:8181/posts \
  -H 'Content-Type: application/json' \
  -d '{
    "title" : "Making a RESTful JSON API in Go - The New Stack",
    "message" : "SpringOne2GX is a conference for app developers, solution and data architects. Sessions are specifically tailored for developers and architects using the popular open source Spring IO Projects, Groovy & Grails, Cloud Foundry, RabbitMQ, Redis, Geode, Hadoop and Tomcat technologies."
  }'

curl -X POST -v http://localhost:8181/posts \
  -H 'Content-Type: application/json' \
  -d '{
    "title" : "Building a REST Service with Golang - stevenwhite.com",
    "message" : "This is the second in a series of high level tutorials on setting up a REST service using golang, focused on using public packages and defining a basic webserver. Packages Go allows for external packages to be accessed in your source via the import statement."
  }'

curl -X POST -v http://localhost:8181/posts \
  -H 'Content-Type: application/json' \
  -d '{
    "title" : "The world''s fastest framework for building websites | Hugo",
    "message" : "Hugo is the fastest tool of its kind. At <1 ms per page, the average site builds in less than a second."
  }'

curl -X POST -v http://localhost:8181/posts \
  -H 'Content-Type: application/json' \
  -d '{
    "title" : "Gobot - Golang framework for robotics, drones.",
    "message" : "Gobot is a framework for robotics, drones, and the Internet of Things (IoT), written in the Go programming language . Gobot - Golang framework for robotics, drones, and the Internet of Things (IoT) Gobot"
  }'

curl -X POST -v http://localhost:8181/posts \
  -H 'Content-Type: application/json' \
  -d '{
    "title" : "Go Web Programming Bootcamp - Go Resources"
    "message" : "Go Web Programming Bootcamp. The following videos were put together in July of 2015 for the Summer Web Bootcamp at Fresno City College. The videos cover the material in An Introduction to Programming in Go along with the basics of server-side web development with Go and Google App Engine."
  }'

curl -X POST -v http://localhost:8181/posts \
  -H 'Content-Type: application/json' \
  -d '{
    "title" : "Create A Simple RESTful API With Golang",
    "message" : "Learn how to create a very simple RESTful API that works with and manipulates mock data using the very fast Golang development langauge."
  }'

