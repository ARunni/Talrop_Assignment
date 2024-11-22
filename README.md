# Talrop_Assignment
 

# Search API in Go (Golang)

 ## *Overview*

**This project implements a Search API to efficiently query and fetch user data based on a given name, with support for phonetic similarity using algorithm Soundex. The API ranks results based on sound similarity and returns a JSON response containing the search results, including metadata such as the total matches and score.

## Deployment
 The application should runs on a specified port 8080

 ## Test the API

 We can test this api using postman or curl 

 End point: (http://localhost:8080/search?name="name"&page="pagenumber")



## Explanation of the algorithm used for phonetic matching.

 ## soundex Function

This function takes a string input and converts it into a Soundex code.
 Soundex is a phonetic algorithm that represents words by their sound rather than their exact spelling. 
 This can be useful for matching similar-sounding names or words, despite spelling variations.

## Function Logic
Initial Letter Retention:
The first letter of the input string is retained in its original case (upper or lower).

## Character Mapping:
The function uses a predefined set of mappings for letters to Soundex digits:
Consonants like 'b', 'f', 'p', 'v' are all mapped to "1", 'c', 'g', 'j', etc., to "2", and so on.
Vowels and non-mapped letters are replaced by the digit "0".

## Digit Replacement:
The rest of the string is converted to lowercase, and each character is replaced by its corresponding Soundex digit. If a character is not in the predefined mapping, it is replaced with "0".

## Duplicate Digit Removal:
Any consecutive duplicate digits are removed, as Soundex eliminates redundant digits.

## Formatting:
The Soundex code is limited to 4 characters. If the code is shorter than 4 characters, it is padded with "0"s. If it is longer, it is truncated to 4 characters.

## Final Output:
The final Soundex code is returned as a string, representing the phonetic sound of the original input.

## We can 
       # run the server using 'make run'
       # run the test using 'make test'
       # run the formating using 'make fmt'

## Env variables 

DB_USER = "your db user"
DB_PASSWORD = "password"
DB_NAME = "db name"

## DB Details

    PostgresSQL

   Create Table by : 

   CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    country TEXT NOT NULL
    );
