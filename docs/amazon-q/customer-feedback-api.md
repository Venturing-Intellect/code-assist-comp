# Customer Feedback API

**Note** Before starting with prompts the go module was created via `go mod init`.

## Starting point

**Prompt 1**: I need to write the following code

For the starting point i've started the AI with the following prompt:

```
I need to write the following code:

Description:
Create a new backend API that enables customers to submit their feedback securely.

Features:

Required: User's email address.
Validate the email format using regular expressions to ensure data accuracy and prevent spam.
Required: The feedback as text.
The feedback is stored in an SQL database.
Business Benefit:
Ensures we collect valid contact information for follow-ups and maintain data integrity.
How can i write the following code using MVC Layered architecture based folder structure & Go CHI framework?

Use 12factor app principle.
```

Even though he didn't provide me with the actual code implementation for the whole features, he gave me a good starting point for the code structure:
```
project-root/
├── cmd/
│   └── server/
│       └── main.go
├── configs/
│   └── config.go
├── internal/
│   ├── controllers/
│   │   └── feedback.go
│   ├── models/
│   │   └── feedback.go
│   ├── repositories/
│   │   └── feedback.go
│   └── services/
│       └── feedback.go
├── pkg/
│   ├── database/
│   │   └── database.go
│   └── utils/
│       └── validator.go
├── Dockerfile
├── go.mod
└── go.sum
```

Additionally he provided me with the following controller, which i can use as Feedback controller:
```go
package controllers

import (
    "net/http"
    "your-project/internal/models"
    "your-project/internal/services"
    "your-project/pkg/utils"
)

type FeedbackController struct {
    feedbackService services.FeedbackService
}

}

func (c *FeedbackController) CreateFeedback(w http.ResponseWriter, r *http.Request) {
    var feedback models.Feedback
    if err := json.NewDecoder(r.Body).Decode(&feedback); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if !utils.IsValidEmail(feedback.Email) {
        http.Error(w, "Invalid email format", http.StatusBadRequest)
        return
    }

    if err := c.feedbackService.CreateFeedback(&feedback); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}
```

**Prompt 2**: Create the folders

Due to AI didn't create the folders automatically, i've asked him to provide me with the command to create the folders.

Here was the prompt:
```
Provide me the linux command which i could use to make all these directories in oneliner command
```

He provided me with the following command:

```
mkdir -p cmd/server internal/{controllers,models,repositories,services} pkg/{database,utils}
```

## Separating code tasks

Due to the complexity of the task for the AI, i had to separate the code tasks into multiple smaller ones.

### Feedback API

**Prompt 3**: Create a new backend API that enables customers to submit their feedback securely.

```
How can i use your previously mentioned code architecture to write the following API using Go CHI framework?

"Create a new backend API that enables customers to submit their feedback securely."
```

The response provided by AI had step-by-step instructions, but didn't show the boilerplate example code, which could be used for the API. He provided me only the controller code:

```go
package controllers

import (
    "encoding/json"
    "net/http"
    "your-project/internal/models"
    "your-project/internal/services"
    "your-project/pkg/utils"
)

 {
    var feedback models.Feedback
    if err := json.NewDecoder(r.Body).Decode(&feedback); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if !utils.IsValidEmail(feedback.Email) {
        http.Error(w, "Invalid email format", http.StatusBadRequest)
        return
    }

    if err := c.feedbackService.CreateFeedback(&feedback); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}
```

So i decided to ask him for the actual code implementation and also gave him the project name.

**Prompt 4**: Provide me with the actual code implementation

```
How can i use your previously mentioned code architecture to write the following API using Go CHI framework?

Provide me with the actual code implementation for all those eight steps you gave me. My project name is "ai-compare-app".
```

With this prompt the AI provided me with the actual code implementation for the API.

Even though the code was almost working, it has one little issue, which i needed to mention him:
```
func (s *FeedbackServiceImpl) CreateFeedback(feedback *models.Feedback) error {
	if !utils.IsValidEmail(feedback.Email) {
		return fmt.Errorf("invalid email format")
	}

	return s.repo.CreateFeedback(feedback)
}
```

The utils.IsValidEmail function was not implemented, so i had to ask him to provide me with the implementation.

**Prompt 5**: Provide me with the implementation of the IsValidEmail function

```
You didn't provide me with the following package and function called utils.IsValidEmail.
```

The AI provided me with the following implementation:
```go
package utils

import (
    "regexp"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func IsValidEmail(email string) bool {
    return emailRegex.MatchString(email)
}
```

**Prompt 6**: Continuous development

Due to Amazon Q Developer provided me all the information needed, i decided to ask him, how can i continue with the development.

Here's the prompt:

```
How can i compile the go code or run it via go run ?
```

The response he gave me was including the following commands:
```
## Solution one
go build -o server cmd/server/main.go
./server

## Alternative
go run cmd/server/main.go

## Set environmental variables
export DATABASE_URL="postgres://user:password@host:port/database?sslmode=disable"
export PORT=8080

And curl to check the POST request
curl -X POST -H "Content-Type: application/json" -d '{"email":"user@example.com","message":"This is a feedback message."}' http://localhost:8080/feedback
```


## Comments from Author

Time wasted:

17.09 - 1 hour

**1. Starting point prompts**

Even though the response of AI provided me with the structure of the code, it didn't provide me with the actual code implementation. I would have liked to see the actual code implementation as well. The prompts needs to be separated into multiple smalls.

Additionally the folder structure needs to be created manually.

**2. Code implementation**

Golang implementation was almost correct, but it had one little issue, which i needed to mention him. The AI didn't provide me with the actual implementation of the IsValidEmail function, so i had to ask him for the implementation. Even though the code is written idiomatically, the code is using one character variable names, which follows the golang development rules, but it's not very readable.