Feature: Translation Service
  Users should be able to submit a word to translate words within the application
 
  Scenario: Translation
    Given the word "hello"
    When I translate it to "german"
    Then the response should be "Hallo"