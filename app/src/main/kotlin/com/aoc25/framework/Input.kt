package com.aoc25.framework

import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.asFlow
import kotlinx.coroutines.flow.flow
import java.nio.file.Path
import kotlin.io.path.exists
import kotlin.io.path.readBytes
import kotlin.io.path.readText

/**
 * Represents puzzle input data with various processing methods.
 * Uses coroutines for efficient streaming of large inputs.
 */
class Input private constructor(
    private val data: ByteArray,
    val source: String
) {
    private val text = String(data)
    
    /**
     * Get all lines as a List
     */
    fun lines(): List<String> = text.split("\n")
    
    /**
     * Stream lines as a Flow for memory-efficient processing
     */
    fun lineFlow(): Flow<String> = lines().asFlow()
    
    /**
     * Stream lines with line numbers as AugmentedLine objects
     */
    fun augmentedLineFlow(): Flow<AugmentedLine> = flow {
        lines().forEachIndexed { index, line ->
            emit(AugmentedLine(line, index))
        }
    }
    
    /**
     * Get all words (whitespace-separated tokens)
     */
    fun words(): List<String> = text.split(Regex("\\s+")).filter { it.isNotEmpty() }
    
    /**
     * Stream words as a Flow
     */
    fun wordFlow(): Flow<String> = words().asFlow()
    
    /**
     * Get all characters
     */
    fun chars(): List<Char> = text.toList()
    
    /**
     * Stream characters as a Flow
     */
    fun charFlow(): Flow<Char> = chars().asFlow()
    
    /**
     * Get the entire input as a string
     */
    fun asString(): String = text
    
    /**
     * Get the raw byte array
     */
    fun asByteArray(): ByteArray = data
    
    companion object {
        /**
         * Create an Input from a file path
         */
        fun fromFile(path: Path): Input {
            if (!path.exists()) {
                throw IllegalArgumentException("Input file not found: $path")
            }
            return Input(path.readBytes(), path.toString())
        }
        
        /**
         * Create an Input from a string (useful for testing)
         */
        fun fromString(text: String, source: String = "string"): Input {
            return Input(text.toByteArray(), source)
        }
        
        /**
         * Generate the small input file name for a day
         */
        fun smallFileName(day: Int): String = "%d_small.input".format(day)
        
        /**
         * Generate the real input file name for a day
         */
        fun realFileName(day: Int): String = "%d.input".format(day)
    }
}

/**
 * Represents a line with its line number for enhanced processing
 */
data class AugmentedLine(
    val text: String,
    val lineNumber: Int
)