package com.example.plugins

import io.ktor.client.*
import io.ktor.client.engine.cio.*
import io.ktor.client.plugins.auth.*
import io.ktor.client.plugins.auth.providers.*
import io.ktor.client.plugins.websocket.*
import io.ktor.http.*
import io.ktor.serialization.*
import io.ktor.serialization.kotlinx.*
import io.ktor.server.application.*
import io.ktor.websocket.*
import kotlinx.coroutines.delay
import kotlinx.coroutines.launch
import kotlinx.serialization.Required
import kotlinx.serialization.Serializable
import kotlinx.serialization.json.Json

@Serializable
data class WebsocketMessage(
    @Required val cmd: String,
    val code: Int?,
    val result: Map<String, String>?,
)

fun Application.configureSockets() {
    val host = environment.config.propertyOrNull("ktor.remote.websocket.host")?.getString()
    val port = environment.config.propertyOrNull("ktor.remote.websocket.port")?.getString()?.toInt()
    val path = environment.config.propertyOrNull("ktor.remote.websocket.path")?.getString() ?: ""
    val masterId = environment.config.propertyOrNull("ktor.remote.websocket.master_id")?.getString() ?: "kotlin"

    val client = HttpClient(CIO) {
        install(WebSockets) {
            contentConverter = KotlinxWebsocketSerializationConverter(Json { ignoreUnknownKeys = true })
        }
        install(Auth) {
            basic {
                credentials {
                    BasicAuthCredentials(
                        username = "ad58e54c8d4",
                        password = "4e7213c403618c4f646ef7"
                    )
                }
                realm = "Authorization Required"
            }
        }
    }

    launch {
        while (true) {
            try {
                client.webSocket(
                    method = HttpMethod.Get,
                    host = host,
                    port = port,
                    path = "$path/$masterId"
                ) {
                    while (true) {
                        try {
                            val message = receiveDeserialized<WebsocketMessage>()
                            if (message.cmd == "GetSessionKey") {
                                val sessionId = message.result?.get("session_id")
                                val sessionKey = message.result?.get("session_key")
                                if (sessionId != null && sessionKey != null) {
                                    SessionIdKeyMap[sessionId] = sessionKey
                                    log.debug("ws: session_id=${sessionId}, session_key=${sessionKey}")
                                } else {
                                    log.error("ws: message.result=${message.result}")
                                }
                            } else {
                                log.error("ws: message=${message}")
                            }
                        } catch (e: WebsocketDeserializeException) {
                            val message = incoming.receive() as? Frame.Text
                            log.error("ws: received '${message?.readText()}'")
                        }
                    }
                }
            } catch (e: Exception) {
                log.warn("ws: $e")
                delay(1000L)
            }

            log.debug("ws: waiting for a while and try again")
            delay(1000L)
        }
    }
}
