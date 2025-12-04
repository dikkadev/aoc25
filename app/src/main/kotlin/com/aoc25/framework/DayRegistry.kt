package com.aoc25.framework

import org.slf4j.LoggerFactory
import kotlin.reflect.KClass
import kotlin.reflect.full.createInstance
import kotlin.reflect.full.findAnnotation
import kotlin.reflect.full.hasAnnotation

/**
 * Annotation to mark classes as AoC day implementations
 */
@Target(AnnotationTarget.CLASS)
@Retention(AnnotationRetention.RUNTIME)
annotation class AoCDay(val number: Int)

/**
 * Registry for discovering and managing AoC day implementations.
 * Uses reflection to automatically discover day classes.
 */
object DayRegistry {
    private val logger = LoggerFactory.getLogger(DayRegistry::class.java)
    private val days = mutableMapOf<Int, Day>()
    
    /**
     * Initialize the registry by scanning for day implementations
     */
    fun initialize() {
        logger.info("Discovering AoC day implementations...")
        
        // Get all classes in the com.aoc25.days package
        val dayClasses = discoverDayClasses()
        
        dayClasses.forEach { kClass ->
            try {
                val annotation = kClass.findAnnotation<AoCDay>()
                if (annotation != null) {
                    val dayInstance = kClass.createInstance() as Day
                    registerDay(dayInstance)
                    logger.info("Registered day ${annotation.number}: ${kClass.simpleName}")
                }
            } catch (e: Exception) {
                logger.error("Failed to instantiate day class ${kClass.simpleName}: ${e.message}")
            }
        }
        
        logger.info("Discovered ${days.size} day implementations")
    }
    
    /**
     * Register a day implementation
     */
    private fun registerDay(day: Day) {
        val dayNumber = day.dayNumber
        if (days.containsKey(dayNumber)) {
            throw IllegalArgumentException("Day $dayNumber is already registered")
        }
        days[dayNumber] = day
    }
    
    /**
     * Get a day implementation by number
     */
    fun getDay(dayNumber: Int): Day? = days[dayNumber]
    
    /**
     * Get all registered days
     */
    fun getAllDays(): Map<Int, Day> = days.toMap()
    
    /**
     * Discover day classes using reflection
     */
    private fun discoverDayClasses(): List<KClass<*>> {
        val dayClasses = mutableListOf<KClass<*>>()
        
        // For now, we'll use a simple approach
        // In a real implementation, you might want to use a classpath scanner
        try {
            // Get the current classloader
            val classLoader = Thread.currentThread().contextClassLoader
            
            // Look for classes in the days package
            val packageName = "com.aoc25.days"
            
            // This is a simplified discovery mechanism
            // In production, you might want to use a more robust scanning approach
            val knownDayClasses = listOf(
                "com.aoc25.days.Day1",
                "com.aoc25.days.Day1_Gold",
                "com.aoc25.days.Day2",
                "com.aoc25.days.Day3",
                "com.aoc25.days.Day4",
                "com.aoc25.days.Day5",
                "com.aoc25.days.Day6",
                "com.aoc25.days.Day7",
                "com.aoc25.days.Day8",
                "com.aoc25.days.Day9",
                "com.aoc25.days.Day10",
                "com.aoc25.days.Day11",
                "com.aoc25.days.Day12",
                "com.aoc25.days.Day13",
                "com.aoc25.days.Day14",
                "com.aoc25.days.Day15",
                "com.aoc25.days.Day16",
                "com.aoc25.days.Day17",
                "com.aoc25.days.Day18",
                "com.aoc25.days.Day19",
                "com.aoc25.days.Day20",
                "com.aoc25.days.Day21",
                "com.aoc25.days.Day22",
                "com.aoc25.days.Day23",
                "com.aoc25.days.Day24",
                "com.aoc25.days.Day25"
            )
            
            knownDayClasses.forEach { className ->
                try {
                    val clazz = Class.forName(className, false, classLoader).kotlin
                    if (clazz.hasAnnotation<AoCDay>()) {
                        dayClasses.add(clazz)
                    }
                } catch (e: ClassNotFoundException) {
                    // Class doesn't exist yet, which is fine
                }
            }
        } catch (e: Exception) {
            logger.warn("Failed to discover day classes: ${e.message}")
        }
        
        return dayClasses
    }
}