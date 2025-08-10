``User Authentication``:-
<br>
Register and Login endpoints using JWT for secure access.
<br>
``Schedule Emails``:-
<br>
Authenticated users can create email schedules (receiver, subject, body, send time, etc.).
<br>
``Background Scheduler``:-
<br>
A goroutine with time.Ticker checks the database every minute for pending emails and “sends” them (simulated with logs).
<br>
``Database Integration``:-
<br>
Emails and user data stored in a relational database via GORM ORM.
<br>
``JWT-Protected Endpoints``:-
<br>
Only logged-in users can schedule emails
<br>
<br>
Database : emails Tables: i)users 2)emails

[Public API]
User Register API
Endpoint : http://localhost:8080/register
```
{
    "username":"sree24432",
    "password":"123444"
}
```
Expected Output:-
```
{
    "message": "User registered successfully"
}
```
[Public API]
Login API
Endpoint : http://localhost:8080/login
```
{
    "username":"sree24432",
    "password":"123444"
}
```
Expected Output:-
```
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTQ5MzQxOTAsInVzZXJfaWQiOjZ9.5Bj3rELM3-fobgm6WpUqJmetQzvwZ1UVZHqL5zYpPY8"
}
```
[Protected API]
Schedule Email API
Endpoint :  http://localhost:8080/schedule-email/
<br>
Add JWT Token in the Authorization Bearer Token
```
{
  "subject": "Meeting Reminder",
  "body": "Don't forget about the meeting tomorrow at 10 AM.",
  "sender": "alice@example.com",
  "receiver": "bob@example.com",
  "attachments": [
    "agenda.pdf",
    "minutes.docx"
  ],
  "schedule_time": "2025-08-11T10:00:00Z",
  "is_draft": false
}
```
Expected Output:-
```
{
    "message": "Email scheduled successfully"
}
```
