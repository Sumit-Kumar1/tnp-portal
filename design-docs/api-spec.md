# TnP Portal API Specification

## Base URL

``` txt
https://api.tnp-portal.com/v1
```

## Authentication

All API endpoints require authentication using JWT tokens. Include the token in the Authorization header:

``` txt
Authorization: Bearer <jwt_token>
```

## API Endpoints

### User Management

#### Register User

```http
POST /users/register
```

Request Body:

```json
{
    "email": "string",
    "password": "string",
    "role": "STUDENT|COMPANY|ADMIN",
    "name": "string"
}
```

#### Login

```http
POST /users/login
```

Request Body:

```json
{
    "email": "string",
    "password": "string"
}
```

### Student Profile Management

#### Create/Update Student Profile

```http
PUT /students/profile
```

Request Body:

```json
{
    "personal_details": {
        "name": "string",
        "email": "string",
        "phone": "string",
        "address": "string"
    },
    "educational_background": [
        {
            "institution": "string",
            "degree": "string",
            "field_of_study": "string",
            "start_date": "date",
            "end_date": "date",
            "gpa": "float"
        }
    ],
    "skills": ["string"],
    "resume_url": "string"
}
```

#### Get Student Profile

```http
GET /students/profile
```

### Job Portal

#### Get Job Listings

```http
GET /jobs
```

Query Parameters:

- `page`: number
- `limit`: number
- `company`: string
- `role`: string
- `location`: string

#### Get Job Details

```http
GET /jobs/{job_id}
```

#### Apply for Job

```http
POST /jobs/{job_id}/apply
```

Request Body:

```json
{
    "resume_url": "string",
    "cover_letter": "string"
}
```

### Company Management

#### Create/Update Company Profile

```http
PUT /companies/profile
```

Request Body:

```json
{
    "name": "string",
    "description": "string",
    "website": "string",
    "industry": "string",
    "size": "string",
    "location": "string"
}
```

#### Post Job Opening

```http
POST /companies/jobs
```

Request Body:

```json
{
    "title": "string",
    "description": "string",
    "requirements": ["string"],
    "location": "string",
    "salary_range": {
        "min": "number",
        "max": "number"
    },
    "deadline": "date"
}
```

### Application Management

#### Get Applications (Student)

```http
GET /students/applications
```

#### Get Applications (Company)

```http
GET /companies/applications
```

#### Update Application Status

```http
PATCH /applications/{application_id}
```

Request Body:

```json
{
    "status": "PENDING|ACCEPTED|REJECTED|INTERVIEW_SCHEDULED|OFFERED"
}
```

### Interview Scheduling

#### Schedule Interview

```http
POST /interviews
```

Request Body:

```json
{
    "application_id": "string",
    "date_time": "datetime",
    "type": "ONLINE|OFFLINE",
    "location": "string",
    "notes": "string"
}
```

#### Get Interviews (Student)

```http
GET /students/interviews
```

#### Get Interviews (Company)

```http
GET /companies/interviews
```

### Document Management

#### Upload Document

```http
POST /documents
```

Request Body:

```json
{
    "type": "RESUME|TRANSCRIPT|CERTIFICATE",
    "file": "binary"
}
```

#### Get Document

```http
GET /documents/{document_id}
```

### Notifications

#### Get Notifications

```http
GET /notifications
```

#### Mark Notification as Read

```http
PATCH /notifications/{notification_id}
```

### Analytics

#### Get Placement Statistics

```http
GET /analytics/placements
```

#### Get Student Performance

```http
GET /analytics/students/performance
```

#### Get Company Recruitment Trends

```http
GET /analytics/companies/trends
```

## Response Format

All responses follow this format:

```json
{
    "status": "success|error",
    "message": "string",
    "data": {} | [],
    "error": {
        "code": "string",
        "message": "string"
    }
}
```

## Error Codes

- 400: Bad Request
- 401: Unauthorized
- 403: Forbidden
- 404: Not Found
- 500: Internal Server Error
