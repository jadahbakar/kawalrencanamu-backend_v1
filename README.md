<!--
https://github.com/othneildrew/Best-README-Template
https://gist.github.com/lukas-h/2a5d00690736b4c3a7ba
https://github.com/aleen42/badges
https://github.com/Naereen/badges
-->

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]


<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/jadahbakar/kawalrencanamu-backend">
    <img src="images/kawalrencanamu.png" alt="Logo" width="158" height="137">
  </a>
  <h2 align="center">Kawal Rencanamu - Backend</h2> <br />
</p>



<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#project-structure">Project structure</a></li>
    <li><a href="#features">Features</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgements">Acknowledgements</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

This system is using for [Kawal Rencanamu][company-website] frontend and mobile system. right now we're using Monolith architecture, but it does not rule out the possibility we rewrite using Microservice architecture.

<!-- BUILD WITH -->
#### Built With

This section we explain what is backend is develop with 
* [Golang](https://golang.org)
* [Fiber](https://docs.gofiber.io/)
* [Viper](https://github.com/spf13/viper)
* [Reflex](https://github.com/cespare/reflex)
* [Air](https://github.com/cosmtrek/air)




<!-- GETTING STARTED -->
## Getting Started

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running follow these simple example steps.

<!-- PREREQUISITES -->
#### Prerequisites

* First you need to install [Golang](https://golang.org/doc/install)

* After that check the installation and Golang version it must be above than 1.11 because we need the [Golang Modules](https://blog.golang.org/using-go-modules)
  ```sh
  > go version
  go version go1.16.3 darwin/amd64
   ```
* Check if your OS have [make](https://www.gnu.org/software/make/), to check it run this command (in my case is using OSX)
  ```
  ❯ make -v
  GNU Make 3.81
  Copyright (C) 2006  Free Software Foundation, Inc.
  This is free software; see the source for copying conditions.
  There is NO warranty; not even for MERCHANTABILITY or FITNESS FOR A
  PARTICULAR PURPOSE.

  This program built for i386-apple-darwin11.3.0
   ```
* Docker is required trought our development process, so highly need it
  ```
  ❯ docker version
  Client:
  Cloud integration: 1.0.14
  Version:           20.10.6
  API version:       1.41
  Go version:        go1.16.3
  Git commit:        370c289
  Built:             Fri Apr  9 22:46:57 2021
  OS/Arch:           darwin/amd64
  Context:           default
  Experimental:      true

  Server: Docker Engine - Community
  Engine:
    Version:          20.10.6
    API version:      1.41 (minimum version 1.12)
    Go version:       go1.13.15
    Git commit:       8728dd2
    Built:            Fri Apr  9 22:44:56 2021
    OS/Arch:          linux/amd64
    Experimental:     false
  containerd:
    Version:          1.4.4
    GitCommit:        05f951a3781f4f2c1911b05e61c160e9c30eaa8e
  runc:
    Version:          1.0.0-rc93
    GitCommit:        12644e614e25b05da6fd08a38ffa0cfe1903fdec
  docker-init:
    Version:          0.19.0
    GitCommit:        de40ad0
   ```
  As you can se the **_Engine_** > **_Version_** > **_20.10.6_** (on my OSX)


<!-- INSTALLATION -->
## Installation

#### Manual Installation
1. Clone the repo
   ```sh
   git clone https://github.com/jadahbakar/kawalrencanamu-backend.git
   ```
2. Install Modules packages
   ```sh
   go mod update
   ```
3. Before to next step, firstable you must manual install [reflex](https://github.com/cespare/reflex)
   ```sh
   go get github.com/cespare/reflex
   ```
4. Check reflex installation
   ```sh
   relex --help
   ```
5. This is Optional if point 4 shows error, you must manualy add to **_SYSTEM PATH_** \n
   in my case i'm using _**zsh**_, so add it to _**.zshrc**_  
   ```sh
   export PATH="$PATH:$(go env GOPATH)/bin"
   ```
   After that refresh the console with
   ```sh
   ❯ . ~/.zshrc
   ```
   Or 
   ```sh
   ❯ source ~/.zshrc
   ```
3. Run it
   ```sh
   make watch
   ```
#### Docker
1. Clone the repo
   ```sh
   git clone https://github.com/jadahbakar/kawalrencanamu-backend.git
   ```
2. Run it (it will create docker images and container with autoreload)
   ```sh
   make devel-up
   ```
3. Stop 
   ```sh
   make devel-down
   ```


<!-- Project structure -->
## Project Structure

```
kawalrencanamu-backend/
├── LICENSE.txt
├── Makefile
├── README.md
├── app.env
├── cmd
│   └── main.go
├── docker
│   ├── development
│   │   ├── dev.Dockerfile
│   │   └── docker-compose.yml
│   └── production
│       ├── Dockerfile
│       ├── exampledocker-compose.yml
│       ├── fixDockerfile\ copy
│       └── v1docker-compose.yml
├── go.mod
├── go.sum
├── images
│   ├── golang.svg
│   └── kawalrencanamu.png
├── internal
│   └── me
│       └── service.go
├── log
│   └── fiber.log
├── pkg
│   ├── config
│   │   ├── appConfig.go
│   │   ├── fiberConfig.go
│   │   └── logerConfig.go
│   ├── middleware
│   │   ├── fiberMiddleware.go
│   │   └── jwtMiddleware.go
│   ├── routes
│   │   ├── not_found_route.go
│   │   └── public_routes.go
│   └── utils
│       ├── jwt_generator.go
│       ├── jwt_parser.go
│       └── startServer.go
├── tmp
│   └── app
│       └── engine
└── version
    └── version.go
```


<!-- FEATURES -->
## Features
- [x] Create scaffolding for Monolith - Hexagonal Architechture
- [ ] Next Will be Microservice Architechture
- [ ] Create Database on PemKot Server
  - [ ] Schema hst (History)
  - [ ] Schema log (Logging)
  - [x] Schema mst (Master)
    * bod
    * connector
    * divisi
    * misi
    * misi_indikator
    * parameter
    * periode
    * unit_indikator
    * value
    * value_indikator
    * visi
    * visi_indikator
    * visi_phase
  - [x] Schema sec (Security)
    * group_user
    * login_method
    * menu_child
    * menu_master
    * role_menu
    * role_otoritas
    * user
  


<!-- LICENSE -->
## License

Distributed under the Proprietary License. See `LICENSE` for more information.


<!-- CONTACT -->
## Contact
Kawalrencanamu - [kawalrencanamu](http://kawalrencanamu.com) - admin@kawalrencanamu.com
<!-- Dedy Styawan - [dedy.styawan](https://twitter.com/dedystyawan) - dedy.styawan@gmail.com
Setiadi Akbar - [setiadi](https://twitter.com/setiadiakbar) - setiadiakbar0@gmail.com -->

<!-- ACKNOWLEDGEMENTS -->
## Acknowledgements
* [Build Docker Image](https://github.com/target/flottbot/blob/master/docker/Dockerfile)
* [Building APIs with Go — Part 2 --> the newest ](https://fernando-bandeira.medium.com/building-apis-with-go-part-2-4afc50fd0ff8)
* [hexagonal-architecture-in-go -> using it now](https://medium.com/@matiasvarela/hexagonal-architecture-in-go-cfd4e436faa3)
* [stygis-golang-hexagonal-architecture](https://idevoid.medium.com/stygis-golang-hexagonal-architecture-a2d89d01f84b)
* [how-i-structure-services-in-go](https://medium.com/@ott.kristian/how-i-structure-services-in-go-19147ad0e6bd)
* [Load config from file & environment variables in Golang with Viper](https://dev.to/techschoolguru/load-config-from-file-environment-variables-in-golang-with-viper-2j2d)
* [Go Fiber](https://gofiber.io/)
* [Go Fiber 2](https://dev.to/koddr/build-a-restful-api-on-go-fiber-postgresql-jwt-and-swagger-docs-in-isolated-docker-containers-475j)
* [semver for versioning](https://github.com/Masterminds/semver/)
* [string to ascii](https://patorjk.com/software/taag/#p=display&h=2&v=3&f=Graceful&t=BUILD)
* [how-to-build-a-smaller-docker-image](https://medium.com/@gdiener/how-to-build-a-smaller-docker-image-76779e18d48a)
* [go-docker-dev-environment-with-go-modules-and-live-code-reloading](https://threedots.tech/post/go-docker-dev-environment-with-go-modules-and-live-code-reloading/)t
* [self-compiling-go-docker-container](https://www.getwrecked.com/self-compiling-go-docker-container/)
* [monitoring-traefik-with-grafana <-- Next Implementation](https://medium.com/platform-engineering/monitoring-traefik-with-grafana-1d037af5b952)

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[company-website]:http://kawalrencanamu.com/
[contributors-shield]: https://img.shields.io/github/contributors/jadahbakar/kawalrencanamu-backend.svg?style=for-the-badge
[contributors-url]: https://github.com/jadahbakar/kawalrencanamu-backend/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/jadahbakar/kawalrencanamu-backend.svg?style=for-the-badge
[forks-url]: https://github.com/jadahbakar/kawalrencanamu-backend/network/members
[stars-shield]: https://img.shields.io/github/stars/jadahbakar/kawalrencanamu-backend.svg?style=for-the-badge
[stars-url]: https://github.com/jadahbakar/kawalrencanamu-backend/stargazers
[issues-shield]: https://img.shields.io/github/issues/jadahbakar/kawalrencanamu-backend.svg?style=for-the-badge
[issues-url]: https://github.com/jadahbakar/kawalrencanamu-backend/issues
[license-shield]: https://img.shields.io/github/license/jadahbakar/kawalrencanamu-backend.svg?style=for-the-badge
[license-url]: https://github.com/jadahbakar/kawalrencanamu-backend/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/dedystyawan
[product-screenshot]: images/screenshot.png
[logo-screenshot]: images/screenshot.png
[golang-image]: images/golang.png
[golang-install-url]: https://golang.org/doc/install

