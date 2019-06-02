module Capitalize_String where

import Data.Char as Char

cap :: String -> String
cap s = Char.toUpper (head s) : tail s
