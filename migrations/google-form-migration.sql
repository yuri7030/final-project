CREATE DATABASE googleFormDb
USE googleFormDb

GO

CREATE TABLE users (
    userId INT IDENTITY(1,1) PRIMARY KEY,
    firstName NVARCHAR(50) NOT NULL,
    lastName NVARCHAR(50) NOT NULL,
	email NVARCHAR(100) NOT NULL,
	[password] NVARCHAR(MAX) NOT NULL,
    createdDate DATETIME DEFAULT GETDATE()
);

GO

CREATE TABLE surveys (
    surveyId INT IDENTITY(1,1) PRIMARY KEY,
    title NVARCHAR(100) NOT NULL,
    createdDate DATETIME DEFAULT GETDATE(),
    createdById INT NOT NULL,
    FOREIGN KEY (createdById) REFERENCES users(userId)
);

GO

CREATE TABLE questions (
    questionId INT IDENTITY(1,1) PRIMARY KEY,
    surveyId INT NOT NULL,
    questionText NVARCHAR(500) NOT NULL,
	answerType INT NOT NULL,
    CONSTRAINT fk_questions_surveys FOREIGN KEY (surveyId) REFERENCES surveys(surveyId)
);

GO

CREATE TABLE answers (
    answerId INT IDENTITY(1,1) PRIMARY KEY,
    questionId INT NOT NULL,
    [guid] VARCHAR(255) NULL,
    answerText NVARCHAR(MAX) NULL,
	numberChoiceAnswer DOUBLE NULL,
	singleChoiceAnswer INT NULL,
	multipleChoiceAnswer VARCHAR(50) NULL,
    CONSTRAINT fk_answers_questions FOREIGN KEY (questionId) REFERENCES questions(questionId)
);

GO