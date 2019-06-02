module Fizz_Buzz where

fizzBuzz :: Int -> String
fizzBuzz a | mod15 a = "FizzBuzz"
           | mod3 a = "Fizz"
           | mod5 a = "Buzz"
           | otherwise = show a

mod3 :: Int -> Bool
mod3 a = mod a 3 == 0

mod5 :: Int -> Bool
mod5 a = mod a 5 == 0

mod15 :: Int -> Bool
mod15 a = mod a 15 == 0
