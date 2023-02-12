# All the materials I discovered while doing my term paper on databases:



## Databases:

#### Theory:
- [SQL indexes (YouTube)](https://youtu.be/LpEwssOYRKA)
- [Что такое базы данных NoSQL?](https://aws.amazon.com/ru/nosql/)
- [Реляционные базы данных обречены (нет)?](https://habr.com/ru/post/103021/)
- [Отличия реляционных и нереляционных баз данных (и когда какую выбирать)](https://www.xelent.ru/blog/otlichiya-relyatsionnykh-i-nerelyatsionnykh-baz-dannykh/)
- [In-memory архитектура для веб-сервисов: основы технологии и принципы](https://habr.com/ru/company/headzio/blog/505792/)

#### InfluxDB:
- [Getting Started with Go and InfluxDB (Official)](https://www.influxdata.com/blog/getting-started-go-influxdb/)
- [Getting Started with Go and InfluxDB](https://thenewstack.io/getting-started-with-go-and-influxdb/)
- [Line protocol](https://docs.influxdata.com/influxdb/v2.0/reference/syntax/line-protocol/)
- [Create a token](https://docs.influxdata.com/influxdb/cloud/security/tokens/create-token/#manage-tokens-in-the-influxdb-ui)
- [InfluxDB Client Go (GitHub)](https://github.com/influxdata/influxdb-client-go)



## Golang:
- [Standard Go Project Layout (GitHub)](https://github.com/golang-standards/project-layout/blob/master/README_ru.md)
- [Fyne tutorials (YouTube)](https://youtube.com/playlist?list=PLjpijTpXl1_po-ld8jORR9k5NornDNKQk)
- [UUID package for Go language (GitHub source code)](https://github.com/satori/go.uuid)
- [Dependency Injection in Go: The better way](https://faun.pub/dependency-injection-in-go-the-better-way-833e2cc65fab)
- [4 примера iota-перечислений](https://habr.com/ru/company/nixys/blog/492056/)
- [Практичный Go: советы по написанию поддерживаемых программ в реальном мире](https://habr.com/ru/post/441842/)
- [Golang Clean Architecture REST API example (GitHub)](https://github.com/AleksK1NG/Go-Clean-Architecture-REST-API)
- [Clean architecture example (OTUS)](https://github.com/OtusGolang/home_work/tree/master/hw12_13_14_15_calendar)
- [Tutorial: Developing a RESTful API with Go and Gin](https://go.dev/doc/tutorial/web-service-gin)
- [How to Wrap and Unwrap Errors in Golang](https://rollbar.com/blog/golang-wrap-and-unwrap-error/)
- [How to generate hash number of a string in Go? (stackoverflow)](https://stackoverflow.com/questions/13582519/how-to-generate-hash-number-of-a-string-in-go)



## Python:

- [Что такое yield в Python? (proglib)](https://proglib.io/p/chto-takoe-yield-v-python-samyy-populyarnyy-vopros-na-stakoverflou-po-pitonu-2022-03-21)
- [Приватные методы без нижнего подчеркивания и интерфейсы в Python (habr)](https://habr.com/ru/post/443192/)



## gRPC:

#### Introduction:
- [gRPC vs REST, что выбрать для нового сервера? (habr)](https://habr.com/ru/post/565020/)
- [Why gRPC / Load balancing options](https://grpc.io/blog/grpc-load-balancing/#why-grpc)
- [gRPC vs REST performance benchmark (Protobuf - JSON - GKE - Terraform - Golang - Kubernetes - k6) (YouTube)](https://www.youtube.com/watch?v=ZwP4ly03n00)
- [gRPC vs. REST: How Does gRPC Compare with Traditional REST APIs?](https://blog.dreamfactory.com/grpc-vs-rest-how-does-grpc-compare-with-traditional-rest-apis/)
- [A curated list of useful resources for gRPC (GitHub)](https://github.com/grpc-ecosystem/awesome-grpc)
- [Concurrency (gRPC Golang) (GitHub)](https://github.com/grpc/grpc-go/blob/master/Documentation/concurrency.md)
- [Best practices for reusing connections, concurrency (gRPC Golang) (GitHub issue)](https://github.com/grpc/grpc-go/issues/682)
- [GRPC Health Checking Protocol (GitHub)](https://github.com/grpc/grpc/blob/master/doc/health-checking.md)
- [Calling grpc server from multi-threaded client (Java) (stackoverflow)](https://stackoverflow.com/questions/72129487/calling-grpc-server-from-multi-threaded-client)

#### Tutorials:
- [gRPC quick start (official)](https://grpc.io/docs/languages/go/quickstart/)
- [gRPC-Go (GitHub source code)](https://github.com/grpc/grpc-go)
- [gRPC-Go (documentation)](https://pkg.go.dev/google.golang.org/grpc)
- [Basics tutorial (gRPC Golang) (official)](https://grpc.io/docs/languages/go/basics/)
- [Basics tutorial (gRPC Python) (official)](https://grpc.io/docs/languages/python/basics/)
- [Generated-code reference (Golang) (official) (about thread-safety and generated methods)](https://grpc.io/docs/languages/go/generated-code/)
- [Как создать простой микросервис на Golang и gRPC и выполнить его контейнеризацию с помощью Docker (habr)](https://habr.com/ru/post/461279/)
- [Go gRPC Beginners Tutorial (unofficial)](https://tutorialedge.net/golang/go-grpc-beginners-tutorial/)
- [Golang gRPC example with making gRPC requests by hand](https://earthly.dev/blog/golang-grpc-example/)

#### Attempts to make/describe an application on gRPC with multiple servers:
- [\[golang\] Multiple servers (and clients) over the same (grpc) connection (Google groups)](https://groups.google.com/g/grpc-io/c/3Mtx9pf1qnY)
- [make it easier to spawn multiple grpc server instances (i.e. kestrel) (GitHub issue C#)](https://github.com/grpc/grpc-dotnet/issues/943)

#### Troubleshooting:
- [mustEmbedUnimplemented*** method appear in grpc-server (GitHub issue)](https://github.com/grpc/grpc-go/issues/3794)
- [Embedding in Golang (documentation)](https://go.dev/doc/effective_go#embedding)
- [Error "protoc-gen-go: program not found or is not executable" (stackoverflow)](https://stackoverflow.com/questions/57700860/error-protoc-gen-go-program-not-found-or-is-not-executable)
- [protoc-gen-go: unable to determine Go import path for "simple.proto" (stackoverflow)](https://stackoverflow.com/questions/70586511/protoc-gen-go-unable-to-determine-go-import-path-for-simple-proto)
- [Fix read behavior when timeout is set by service config (GitHub issue)](https://github.com/grpc/grpc-go/issues/1818)



## Software design:

#### Patterns:
- [ActiveRecord pattern (Wikipedia)](https://ru.wikipedia.org/wiki/ActiveRecord)
- [Dependency injection (design pattern) (Wikipedia)](https://en.wikipedia.org/wiki/Dependency_injection)
- [What is Dependency Injection?](https://www.growin.com/blog/what-is-dependency-injection/#:~:text=Dependency%20Injection%20(DI%20is%20a,for%20example%20as%20a%20service.)
- [Template method pattern (Wikipedia)](https://en.wikipedia.org/wiki/Template_method_pattern)
- [Sentinel value (Wikipedia)](https://en.wikipedia.org/wiki/Sentinel_value)
- [Паттерн репозиторий (Github gist)](https://gist.github.com/maestrow/594fd9aee859c809b043)
- [The Repository Pattern Explained (For Dummies)](https://blog.sapiensworks.com/post/2014/06/02/The-Repository-Pattern-For-Dummies.aspx)
- [Repository Pattern](https://lyz-code.github.io/blue-book/architecture/repository_pattern/)
- [Are the Repository Pattern and Active Record pattern compatible? (stackoverflow)](https://softwareengineering.stackexchange.com/questions/284865/are-the-repository-pattern-and-active-record-pattern-compatible)

#### Microservice architechture:
- [Микросервисы и микросервисная архитектура](https://www.atlassian.com/ru/microservices)
- [26 основных паттернов микросервисной разработки](https://mcs.mail.ru/blog/26-osnovnyh-patternov-mikroservisnoj-razrabotki)
- [Сравнение микросервисной и монолитной архитектур](https://www.atlassian.com/ru/microservices/microservices-architecture/microservices-vs-monolith)

#### Other:
- [?](https://habr.com/ru/company/funcorp/blog/372199/)
- [Бессерверные вычисления (Wikipedia)](https://ru.wikipedia.org/wiki/%D0%91%D0%B5%D1%81%D1%81%D0%B5%D1%80%D0%B2%D0%B5%D1%80%D0%BD%D1%8B%D0%B5_%D0%B2%D1%8B%D1%87%D0%B8%D1%81%D0%BB%D0%B5%D0%BD%D0%B8%D1%8F)
- [Что такое Zookeeper](https://www.bigdataschool.ru/wiki/zookeeper)



## Desktop GUI:

#### Design:
- [An arctic, north-bluish color palette.](https://www.nordtheme.com/)
- [Цвет в дизайне интерфейсов: инструкция по применению (habr)](https://habr.com/ru/company/iloveip/blog/337060/)

#### Libraries:
- [List of Golang libraries for building GUI Applications (GitHub)](https://github.com/avelino/awesome-go#gui)
- [GUI (List of Golang libraries for building web/desktop GUI)](https://golangr.com/gui/)
- [Fyne.io](https://fyne.io/)
- [Fyne Tutorials (YouTube](https://www.youtube.com/playlist?list=PLjpijTpXl1_po-ld8jORR9k5NornDNKQk)
- [therecipe/qt (GitHub) (it allows to write Qt applications entirely in Go, JavaScript/TypeScript, Dart/Flutter, Haxe and Swift)](https://github.com/therecipe/qt)
- [go-gtk (GitHub)](https://github.com/mattn/go-gtk/)
- [GUI на Golang: GTK+ 3 (habr)](https://habr.com/ru/post/420035/)
- [webview (GitHub) (webview library for C/C++/Go to build GUIs)](https://github.com/webview/webview)
- [ui: platform-native GUI library for Go (GitHub)](https://github.com/andlabs/ui)

#### Qt:
- [QTextCursor::insertBlock](http://doc.crossplatform.ru/qt/4.7.x/qtextcursor.html#insertBlock)
- [Выделенный текст в QTextEdit сделать жирным (cyberforum)](https://www.cyberforum.ru/qt/thread2240270.html)
- [Чтение чисел из QTextEdit (cyberforum)](https://www.cyberforum.ru/qt/thread384322.html)
- [Форматирование выбранного текста (cyberforum)](https://www.cyberforum.ru/qt/thread1605957.html)
- [How to set a GUI Theme to a Qt Widgets Application (YouTube)](https://www.youtube.com/watch?v=zjWfDEUsobQ)
- [Creating PyQt Layouts for GUI Python Applications (YouTube)](https://www.youtube.com/watch?v=MY29YV9Wk7I)
- [Реализация простейшего текстового редактора на Qt (YouTube)](https://www.youtube.com/watch?v=FRcbOZ8mTJw)



## Web:
- [A basic tutorial introduction to gRPC-web.](https://grpc.io/docs/platforms/web/basics/)
- [Building microservices in Go and Python using gRPC and TLS/SSL authentication](https://gustavoh.medium.com/building-microservices-in-go-and-python-using-grpc-and-tls-ssl-authentication-cfcee7c2b052)
- [clay (Minimal server platform for gRPC and REST+Swagger APIs in Go)](https://github.com/utrack/clay)
- [Building a realtime dashboard with ReactJS, Go, gRPC, and Envoy.](https://medium.com/swlh/building-a-realtime-dashboard-with-reactjs-go-grpc-and-envoy-7be155dfabfb)
- [Using GRPC with TLS, Golang and React (No Envoy)](https://programmingpercy.tech/blog/using-grpc-tls-go-react-no-reverse-proxy/)
- [Small Go/React/TypeScript gRPC-Web example (GitHub)](https://github.com/johanbrandhorst/grpc-web-go-react-example)
- [Build a chat app using gRPC and React (with Envoy)](https://daily.dev/blog/build-a-chat-app-using-grpc-and-reactjs)
- [Detailed full-stack flow with gRPC-Web, Go and React (Reddit)](https://www.reddit.com/r/golang/comments/sqthd7/go_and_grpc_is_just_so_intuitive_heres_a_detailed/)
- [gRPC-Web Demo (with Envoy) (GitHub)](https://github.com/uid4oe/grpc-web-demo)
- [Streaming data with gRPC](https://itnext.io/streaming-data-with-grpc-2eb983fdee11)
- [Securing your GitHub Pages site with HTTPS](https://docs.github.com/en/pages/getting-started-with-github-pages/securing-your-github-pages-site-with-https)
- [Build CRUD gRPC Server API & Client with Golang and MongoDB](https://codevoweb.com/crud-grpc-server-api-client-with-golang-and-mongodb/)
- [How to use gRPC-web with React](https://morioh.com/p/ae48b33d10a0)
- [Building a Chat Application in Go with ReactJS / Dockerizing your Backend](https://tutorialedge.net/projects/chat-system-in-go-and-react/part-6-dockerizing-your-backend/)
- [Full Stack Application with Go, Gin, React, and MongoDB](https://medium.com/geekculture/full-stack-application-with-go-gin-react-and-mongodb-37b63ef71133)
- [Зачем нужен обратный прокси сервер в 5 актах](https://itnan.ru/post.php?c=1&p=538936)
- [Defining Stateful vs Stateless Web Services](https://nordicapis.com/defining-stateful-vs-stateless-web-services/)
- [MEAN vs. MERN vs. MEVN Stacks: What’s the Difference?](https://kenzie.snhu.edu/blog/mean-vs-mern-vs-mevn-stacks-whats-the-difference/)
- [curl only results in "Empty reply from server" (GitHub issue)](https://github.com/moby/moby/issues/2522)



## Deploy:
- [How We Deploy Python Code](https://www.nylas.com/blog/packaging-deploying-python/)
- [Deploying Go + React to Heroku using Docker, Part 1](https://levelup.gitconnected.com/deploying-go-react-to-heroku-using-docker-9844bf075228)
- [Deploying Go + React to Heroku using Docker, Part 2 (the database)](https://dean-baker.medium.com/deploying-go-react-to-heroku-using-docker-part-2-the-database-afaaaae66f81)
- [5 Heroku Alternatives for Free Full Stack Hosting](https://www.makeuseof.com/heroku-alternatives-free-full-stack-hosting/)
- [Getting Started with GitHub Pages (YouTube)](https://www.youtube.com/watch?v=QyFcl_Fba-k&t=0s)
- [Деплой React-приложения на Vercel, Netlify, Heroku, GitHub Pages, Surge (YouTube)](https://www.youtube.com/watch?v=-pJN9faoa8E)
- [Deploy to a Server on Render.com](https://blitzjs.com/docs/deploy-render)
- [Deployment types (envoy)](https://www.envoyproxy.io/docs/envoy/latest/intro/deployment_types/deployment_types)

## Docker:
- [Multi-stage builds](https://docs.docker.com/build/building/multi-stage/)
- [Build your Go image](https://docs.docker.com/language/golang/build-images/)
- [Dockerfile reference](https://docs.docker.com/engine/reference/builder/)
- [Tutorial: Create multi-container apps with MySQL and Docker Compose (Microsoft)](https://learn.microsoft.com/en-us/visualstudio/docker/tutorials/tutorial-multi-container-app-mysql)
- [Using Docker with Postgres: Tutorial and Best Practices](https://earthly.dev/blog/postgres-docker/)
- [Networking in Compose](https://docs.docker.com/compose/networking/)
- [Use bridge networks](https://docs.docker.com/network/bridge/)
- [Using Redis with docker and docker-compose for local development a step-by-step tutorial](https://geshan.com.np/blog/2022/01/redis-docker/)



## Other:
- [46 KILLER GOLANG Projects in 46 Different Videos (Youtube)](https://www.youtube.com/playlist?list=PL5dTjWUk_cPYztKD7WxVFluHvpBNM28N9)
- [Использование диаграммы вариантов использования UML при проектировании программного обеспечения](https://habr.com/ru/post/566218/)
- [ritual (GitHub) (it allows to use C++ libraries from Rust)](https://github.com/rust-qt/ritual)
- [GTK-server example scripts (Какой-то древний ужас)](https://www.gtk-server.org/examples.html)
- [Теги strong и b, важность и выделение](https://htmlacademy.ru/courses/301/run/17)
