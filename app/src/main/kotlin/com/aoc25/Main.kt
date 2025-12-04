package com.aoc25

import com.github.ajalt.clikt.core.CliktCommand
import com.github.ajalt.clikt.core.main
import com.github.ajalt.clikt.parameters.options.flag
import com.github.ajalt.clikt.parameters.options.option
import com.github.ajalt.clikt.parameters.types.int
import com.aoc25.framework.DayRegistry
import com.aoc25.framework.Input
import kotlinx.coroutines.runBlocking
import org.slf4j.LoggerFactory
import java.io.File

/**
 * Main application class for AoC 2025
 */
class AoCApp : CliktCommand() {
    private val logger = LoggerFactory.getLogger(AoCApp::class.java)
    
    private val dayNumber: Int? by option("-d", "--day", help = "Day to run (1-25)").int()
    private val small: Boolean by option("-s", "--small", help = "Use small input").flag(default = false)
    private val debug: Boolean by option("--debug", help = "Enable debug logging").flag(default = false)
    
    override fun run() = runBlocking {
        // Configure logging level
        if (debug) {
            System.setProperty("org.slf4j.simpleLogger.defaultLogLevel", "debug")
        }
        
        logger.info("Starting AoC 2025 Solver")
        logger.info("Java version: ${System.getProperty("java.version")}")
        logger.info("Kotlin version: ${KotlinVersion.CURRENT}")
        
        // Initialize day registry
        DayRegistry.initialize()
        
        // Get the requested day
        val day = dayNumber?.let { DayRegistry.getDay(it) }
        if (day == null || dayNumber == null) {
            logger.error("Day $dayNumber not found or not implemented")
            echo("Day $dayNumber not found. Available days: ${DayRegistry.getAllDays().keys.sorted()}")
            return@runBlocking
        }
        
        // Prepare input
        val inputFileName = if (small) {
            Input.smallFileName(dayNumber!!)
        } else {
            Input.realFileName(dayNumber!!)
        }
        
        // Try to load from classpath resources first
        val resource = this::class.java.classLoader.getResource("input/$inputFileName")
        val inputFile = if (resource != null) {
            File(resource.toURI())
        } else {
            // Fallback to direct file path
            File("app/src/main/resources/input/$inputFileName")
        }
        if (!inputFile.exists()) {
            logger.error("Input file not found: $inputFileName")
            echo("Input file not found: $inputFileName")
            echo("Please create the file and add your puzzle input.")
            return@runBlocking
        }
        
        val input = Input.fromFile(inputFile.toPath())
        val dayLogger = LoggerFactory.getLogger("Day${dayNumber!!.toString().padStart(2, '0')}")
        
        logger.info("Solving day $dayNumber (small=$small)")
        
        try {
            val startTime = System.currentTimeMillis()
            val result = day.solve(input, dayLogger)
            val endTime = System.currentTimeMillis()
            
            echo("\n🎄 Day $dayNumber Result:")
            echo("📝 $result")
            echo("⏱️  Solved in ${endTime - startTime}ms")
            
            logger.info("Day $dayNumber solved successfully: $result")
            
        } catch (e: Exception) {
            logger.error("Failed to solve day $dayNumber", e)
            echo("❌ Error solving day $dayNumber: ${e.message}")
            if (debug) {
                e.printStackTrace()
            }
        }
    }
}

/**
 * Main entry point
 */
fun main(args: Array<String>) {
    AoCApp().main(args)
}