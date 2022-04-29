package com.example.plugins

import io.ktor.http.HttpStatusCode.Companion.InternalServerError
import io.ktor.server.application.*
import io.ktor.server.auth.*
import io.ktor.server.request.*
import io.ktor.server.response.*
import io.ktor.server.routing.*
import kotlinx.serialization.Serializable
import kotlinx.serialization.decodeFromString
import kotlinx.serialization.json.Json
import nia.auth.gateway.sdk.NiaGatewaySdk
import nia.auth.gateway.sdk.dto.req.EquipmentIdAuthReq
import nia.auth.gateway.sdk.dto.req.EquipmentRegisterReq
import org.apache.commons.codec.digest.DigestUtils
import org.bouncycastle.jce.provider.BouncyCastleProvider
import java.security.InvalidAlgorithmParameterException
import java.security.InvalidKeyException
import java.security.NoSuchAlgorithmException
import java.security.Security
import java.text.SimpleDateFormat
import java.util.*
import javax.crypto.*
import javax.crypto.spec.SecretKeySpec

fun sha1LeftCut(text: String, length: Int = 16): ByteArray {
    return DigestUtils.sha1(text).copyOfRange(0, length)
}

@Throws(
    NoSuchPaddingException::class,
    NoSuchAlgorithmException::class,
    InvalidAlgorithmParameterException::class,
    InvalidKeyException::class,
    BadPaddingException::class,
    IllegalBlockSizeException::class
)
fun aesEcbEncrypt(plainText: String, password: String): String {
    val cipher: Cipher = Cipher.getInstance("AES/ECB/PKCS7Padding", "BC")
    val key: SecretKey = SecretKeySpec(sha1LeftCut(password), "AES")
    cipher.init(Cipher.ENCRYPT_MODE, key)
    val cipherText: ByteArray = cipher.doFinal(plainText.toByteArray())
    return Base64.getEncoder()
        .encodeToString(cipherText)
}

@Throws(
    NoSuchPaddingException::class,
    NoSuchAlgorithmException::class,
    InvalidAlgorithmParameterException::class,
    InvalidKeyException::class,
    BadPaddingException::class,
    IllegalBlockSizeException::class
)
fun aesEcbDecrypt(cipherText: String, password: String): String {
    val cipher: Cipher = Cipher.getInstance("AES/ECB/PKCS7Padding", "BC")
    val key: SecretKey = SecretKeySpec(sha1LeftCut(password), "AES")
    cipher.init(Cipher.DECRYPT_MODE, key)
    val plainText = cipher.doFinal(
        Base64.getDecoder()
            .decode(cipherText)
    )
    return String(plainText)
}

@Serializable
data class Response(val code: Int, val message: String, val error: String = "")

private val ResponseInternalServerError = Response(InternalServerError.value, InternalServerError.description)

private val json = Json { ignoreUnknownKeys = true }

private fun getTimestamp() = SimpleDateFormat("yyyyMMddHHmmssSSS").format(Date())

fun Application.configureRouting() {
    Security.addProvider(BouncyCastleProvider())

    routing {
        get("/debug/crypto") {
            var cipherText: String? = call.request.queryParameters["cipher_text"]
            var plainText: String? = call.request.queryParameters["plain_text"]
            val password: String = call.request.queryParameters["password"] ?: "password"

            call.application.environment.log.debug("cipher_text=${cipherText}")
            call.application.environment.log.debug("plainText=${plainText}")

            if (!cipherText.isNullOrBlank()) {
                plainText = aesEcbDecrypt(cipherText, password)
                cipherText = aesEcbEncrypt(plainText, password)
            } else if (!plainText.isNullOrBlank()) {
                cipherText = aesEcbEncrypt(plainText, password)
                plainText = aesEcbDecrypt(cipherText, password)
            }

            call.respond(
                mapOf(
                    "cipher_text" to cipherText,
                    "plain_text" to plainText,
                )
            )
        }
    }

    routing {
        authenticate("auth-basic") {

            }
        }
    }
}
