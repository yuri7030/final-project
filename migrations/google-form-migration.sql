CREATE DATABASE googleFormDb
USE googleFormDb

GO

CREATE TABLE users (
    userId INT IDENTITY(1,1) PRIMARY KEY,
    [name] NVARCHAR(50) NOT NULL,
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

CREATE TABLE options (
    optionId INT IDENTITY(1,1) PRIMARY KEY,
    questionId INT NOT NULL,
    optionText NVARCHAR(500) NOT NULL,
    CONSTRAINT fk_option_question FOREIGN KEY (questionId) REFERENCES surveys(questionId)
);

GO

CREATE TABLE answers (
    answerId INT IDENTITY(1,1) PRIMARY KEY,
    questionId INT NOT NULL,
    guestId VARCHAR(32) NULL,
	singleOptionId INT NULL,
	multipleOptionIds VARCHAR(50) NULL,
    createdDate DATETIME DEFAULT GETDATE(),
    CONSTRAINT fk_answers_questions FOREIGN KEY (questionId) REFERENCES questions(questionId)
);

GO