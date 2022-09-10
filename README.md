# Hexagonal Architecture template
## Description
 
This is a repository aimed to serve as a template for creating an hexagonal 
architecture project and by using dependency injection without the need of 
3rd party libraries.

## Usage

The injection system is heavily inspired by NestJS

Services, Repositories, Controllers, UseCases... must be registered in the InitAppModule() method.
There is a map for the controllers, for the repositories and for the services/use cases(providers)

The ping/pong controllers are examples on how to retrieve instances of each dependency

**This project is not meant for production use but as a side hobby project for learning purposes**

