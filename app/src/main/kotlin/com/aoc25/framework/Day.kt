package com.aoc25.framework

import com.aoc25.framework.Input
import org.slf4j.Logger

/**
 * Interface for Advent of Code day implementations.
 * Each day should implement this interface to provide solutions.
 */
interface Day {
    /**
     * The day number (1-25)
     */
    val dayNumber: Int
    
    /**
     * Solve the puzzle for the given input.
     * 
     * @param input The parsed input data
     * @param logger Logger for debugging
     * @return The puzzle solution
     */
    suspend fun solve(input: Input, logger: Logger): Any
}