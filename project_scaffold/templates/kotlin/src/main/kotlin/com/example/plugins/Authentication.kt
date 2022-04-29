package com.example.plugins

import io.ktor.server.application.*
import io.ktor.server.auth.*


fun Application.configureAuthentication() {
    install(Authentication) {
        basic("auth-basic") {
            realm = "Authorization Required"
            validate { credentials ->
                if (credentials.name == "ebae31cf344828" && credentials.password == "ed4e250de6f26ca65f9") {
                    UserIdPrincipal(credentials.name)
                } else {
                    null
                }
            }
        }
    }
}
