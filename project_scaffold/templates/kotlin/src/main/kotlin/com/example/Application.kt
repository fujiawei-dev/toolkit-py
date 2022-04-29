package com.example

import io.ktor.server.application.*
import com.example.plugins.*

fun main(args: Array<String>): Unit =
    io.ktor.server.netty.EngineMain.main(args)

@Suppress("unused")
fun Application.module() {
    configureRouting()
    configureAuthentication()
    configureHTTP()
    configureSerialization()
    configureSockets()
    configureAdministration()
}
