# Google Form Clone Database Design

![Golang Logo](https://golang.org/doc/gopher/frontpage.png)

The Google Form Clone Project is an interactive web application that replicates the functionality of Google Forms, allowing users to create surveys, design questions, and collect responses. This README provides an overview of the project's database structure, explaining the entities involved and the relationships among them.

The project uses a SQL Server database named `GoogleFormDb` to store user and survey data.

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

Please note that this README provides an overview of the database design and relationships. For detailed implementation and usage instructions, refer to the project documentation or source code.
