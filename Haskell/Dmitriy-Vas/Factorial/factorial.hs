module Factorial where

factorial :: Integer -> Integer
factorial x | x == 0 = 1
            | otherwise = product [1..x]
