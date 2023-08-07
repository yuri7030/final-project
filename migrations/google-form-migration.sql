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
    title VARCHAR(255) NOT NULL,
    [description] VARCHAR(500) NULL,
    createdDate DATETIME DEFAULT GETDATE(),
    createdBy INT NOT NULL,
    FOREIGN KEY (createdBy) REFERENCES users(userId)
);

GO

CREATE TABLE questions (
    questionId INT IDENTITY(1,1) PRIMARY KEY,
    surveyId INT NOT NULL,
    questionText NVARCHAR(500) NOT NULL,
	answerType INT NOT NULL,
    createdDate DATETIME DEFAULT GETDATE(),
    CONSTRAINT fk_questions_surveys FOREIGN KEY (surveyId) REFERENCES surveys(surveyId)
);

GO

CREATE TABLE options (
    optionId INT IDENTITY(1,1) PRIMARY KEY,
    questionId INT NOT NULL,
    optionText NVARCHAR(500) NOT NULL,
    createdDate DATETIME DEFAULT GETDATE(),
    CONSTRAINT fk_option_question FOREIGN KEY (questionId) REFERENCES questions(questionId)
);

GO

CREATE TABLE answers (
    Id INT IDENTITY(1,1) PRIMARY KEY,
    questionId INT NOT NULL,
    [guid] VARCHAR(50) NOT NULL,
    answerText NVARCHAR(500) NULL,
	singleOptionId INT NULL,
	multipleOptionIds VARCHAR(50) NULL,
    createdDate DATETIME DEFAULT GETDATE(),
    CONSTRAINT fk_answers_questions FOREIGN KEY (questionId) REFERENCES questions(questionId)
);

GO