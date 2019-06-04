module Fizz_Buzz where

fizzBuzz :: Int -> String
fizzBuzz a | mod15 = "FizzBuzz"
           | mod3 = "Fizz"
           | mod5 = "Buzz"
           | otherwise = show a
  where mod3 = mod a 3 == 0
        mod5 = mod a 5 == 0
        mod15 = mod a 15 == 0
