package com.aoc25.days

import com.aoc25.framework.AoCDay
import com.aoc25.framework.Day
import com.aoc25.framework.Input
import org.slf4j.Logger

/**
 * Template for AoC day implementations.
 * Copy this file and modify for each day.
 */
@AoCDay(number = 1) // Change this to the actual day number
class Day1_Gold : Day {
    
    override val dayNumber: Int = 1 // Change this to the actual day number
    
    override suspend fun solve(input: Input, logger: Logger): Any {
        var result = 0
        
        // Example: Process lines
        input.augmentedLineFlow().collect { line ->
            if (line.text.isNotEmpty()) {
                logger.debug("Line ${line.lineNumber}: ${line.text}")
                
                // Your solution logic here
                // Example: result += line.text.length
            }
        }
        
        // Alternative: Process words
        // input.wordFlow().collect { word ->
        //     logger.debug("Word: $word")
        // }
        
        // Alternative: Process characters
        // input.charFlow().collect { char ->
        //     logger.debug("Char: $char")
        // }
        
        // Alternative: Get all data at once
        // val lines = input.lines()
        // val words = input.words()
        // val text = input.asString()
        
        return result
    }
}