# Google Form Clone Database Design

![Golang Logo](https://golang.org/doc/gopher/frontpage.png)

The Google Form Clone Project is an interactive web application that replicates the functionality of Google Forms, allowing users to create surveys, design questions, and collect responses. This README provides an overview of the project's database structure, explaining the entities involved and the relationships among them.

The project uses a SQL Server database named `GoogleFormDb` to store user and survey data.

## Features

- **User Management**: Users can sign up, providing their first name, last name, and creation date. This information is stored in the `Users` table.

- **Survey Creation**: Authenticated users can create surveys by providing a title, which is stored along with the creation date and the user ID in the `Surveys` table.

- **Question Creation**: Surveys can have multiple questions, each with a question text and an answer type. These questions are stored in the `Questions` table and linked to the corresponding survey.

- **Response Collection**: Users can respond to surveys by providing answers. Responses are stored in the `Answers` table, linked to the corresponding question. The response can include answer text, single choice selection, or multiple choice selections.

## Getting Started

To run the Google Form Clone project locally:

1. Set up a SQL Server instance and download this [script](google-form-migration.sql)
2. Execute the provided SQL script in your SQL Server management tool to create the necessary tables and relationships.
3. Clone this repository to your local machine.
4. Configure the database connection settings in the project to point to your SQL Server instance.
5. Run the application and start exploring the features.

Feel free to customize and enhance the project to meet your specific requirements or to further extend its functionality.

### Users Table

The `Users` table holds information about users who interact with the application. It stores the following details:

- `UserId`: A unique identifier for each user.
- `FirstName`: The first name of the user (up to 50 characters).
- `LastName`: The last name of the user (up to 50 characters).
- `CreatedDate`: The date and time when the user account was created.

### Surveys Table

The `Surveys` table represents the surveys created by users. It includes the following fields:

- `SurveyId`: A unique identifier for each survey.
- `Title`: The title of the survey (up to 100 characters).
- `CreatedDate`: The date and time when the survey was created.
- `CreatedById`: A reference to the `UserId` of the user who created the survey.

### Questions Table

The `Questions` table stores the questions associated with each survey. Key attributes are:

- `QuestionId`: A unique identifier for each question.
- `SurveyId`: A reference to the `SurveyId` of the survey to which the question belongs.
- `QuestionText`: The text of the question (up to 500 characters).
- `AnswerType`: An indicator of the type of answer expected for the question.

### Answers Table

The `Answers` table records responses to the survey questions. It includes:

- `AnswerId`: A unique identifier for each answer.
- `QuestionId`: A reference to the `QuestionId` of the question being answered.
- `RespondentId`: A reference to the `UserId` of the user who provided the answer.
- `AnswerText`: For open-ended questions, the text of the respondent's answer (up to maximum length).
- `SingleChoiceAnswer`: For single-choice questions, the selected answer option.
- `MultipleChoiceAnswer`: For multiple-choice questions, the selected answer options.

## Relationships

The database entities are connected through the following relationships:

- Each survey (in the `Surveys` table) is associated with a user who created it (via `CreatedById`).
- Survey questions (in the `Questions` table) belong to a specific survey (via `SurveyId`).
- Answers (in the `Answers` table) are linked to the question they respond to (via `QuestionId`) and the respondent (via `RespondentId`).

This project's database structure facilitates the creation of surveys, design of various types of questions, and collection of responses, enabling a simplified version of Google Forms functionality.

## Contributions

Contributions to the Google Form Clone project are welcome! Whether it's bug fixes, feature enhancements, or documentation improvements, your contributions can help make this project better.

We hope you enjoy working on and using the Google Form Clone project!
