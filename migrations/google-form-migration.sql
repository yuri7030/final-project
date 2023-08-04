CREATE DATABASE GoogleFormDb
USE GoogleFormDb

GO

CREATE TABLE Users (
    UserId INT IDENTITY(1,1) PRIMARY KEY,
    FirstName NVARCHAR(50) NOT NULL,
    LastName NVARCHAR(50) NOT NULL,
    CreatedDate DATETIME DEFAULT GETDATE()
);

GO

CREATE TABLE Surveys (
    SurveyId INT IDENTITY(1,1) PRIMARY KEY,
    Title NVARCHAR(100) NOT NULL,
    CreatedDate DATETIME DEFAULT GETDATE(),
    CreatedById INT NOT NULL,
    FOREIGN KEY (CreatedById) REFERENCES Users(UserId)
);

GO

CREATE TABLE Questions (
    QuestionId INT IDENTITY(1,1) PRIMARY KEY,
    SurveyId INT NOT NULL,
    QuestionText NVARCHAR(500) NOT NULL,
	AnswerType INT NOT NULL,
    CONSTRAINT FK_Questions_Surveys FOREIGN KEY (SurveyId) REFERENCES Surveys(SurveyId)
);

GO

CREATE TABLE Answers (
    AnswerId INT IDENTITY(1,1) PRIMARY KEY,
    QuestionId INT NOT NULL,
    RespondentId INT NULL,
    AnswerText NVARCHAR(MAX) NULL,
	SingleChoiceAnswer INT NULL,
	MultipleChoiceAnswer INT NULL,
    CONSTRAINT FK_Answers_Questions FOREIGN KEY (QuestionId) REFERENCES Questions(QuestionId)
);

GO